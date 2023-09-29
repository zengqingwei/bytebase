package mysql

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	"github.com/bytebase/bytebase/backend/plugin/parser/tokenizer"
)

// ParseResult is the result of parsing a MySQL statement.
type ParseResult struct {
	Tree     antlr.Tree
	Tokens   *antlr.CommonTokenStream
	BaseLine int
}

// ParseMySQL parses the given SQL statement and returns the AST.
func ParseMySQL(statement string) ([]*ParseResult, error) {
	var err error
	statement, err = DealWithDelimiter(statement)
	if err != nil {
		return nil, err
	}

	return parseInputStream(antlr.NewInputStream(statement))
}

// DealWithDelimiter deals with delimiter in the given SQL statement.
func DealWithDelimiter(statement string) (string, error) {
	has, list, err := hasDelimiter(statement)
	if err != nil {
		return "", err
	}
	if has {
		var result []string
		delimiter := `;`
		for _, sql := range list {
			if IsDelimiter(sql.Text) {
				delimiter, err = ExtractDelimiter(sql.Text)
				if err != nil {
					return "", err
				}
				result = append(result, "-- "+sql.Text)
				continue
			}
			// TODO(rebelice): after deal with delimiter, we may cannot get the right line number, fix it.
			if delimiter != ";" {
				result = append(result, fmt.Sprintf("%s;", strings.TrimSuffix(sql.Text, delimiter)))
			} else {
				result = append(result, sql.Text)
			}
		}

		statement = strings.Join(result, "\n")
	}
	return statement, nil
}

// SplitMySQL splits the given SQL statement into multiple SQL statements.
func SplitMySQL(statement string) ([]base.SingleSQL, error) {
	statement = strings.TrimRight(statement, " \r\n\t\f;") + "\n;"
	var err error
	statement, err = DealWithDelimiter(statement)
	if err != nil {
		return nil, err
	}

	lexer := parser.NewMySQLLexer(antlr.NewInputStream(statement))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return splitMySQLStatement(stream)
}

// SplitMySQLStream splits the given SQL stream into multiple SQL statements.
// Note that the reader is read completely into memory and so it must actually
// have a stopping point - you cannot pass in a reader on an open-ended source such
// as a socket for instance.
func SplitMySQLStream(src io.Reader) ([]base.SingleSQL, error) {
	text := antlr.NewIoStream(src).String()
	return SplitMySQL(text)
}

// ParseMySQLStream parses the given SQL stream and returns the AST.
// Note that the reader is read completely into memory and so it must actually
// have a stopping point - you cannot pass in a reader on an open-ended source such
// as a socket for instance.
func ParseMySQLStream(src io.Reader) ([]*ParseResult, error) {
	text := antlr.NewIoStream(src).String()
	return ParseMySQL(text)
}

func getDefaultChannelTokenType(tokens []antlr.Token, base int, offset int) int {
	current := base
	step := 1
	remaining := offset
	if offset < 0 {
		step = -1
		remaining = -offset
	}
	for remaining != 0 {
		current += step
		if current < 0 || current >= len(tokens) {
			return parser.MySQLParserEOF
		}

		if tokens[current].GetChannel() == antlr.TokenDefaultChannel {
			remaining--
		}
	}

	return tokens[current].GetTokenType()
}

func splitMySQLStatement(stream *antlr.CommonTokenStream) ([]base.SingleSQL, error) {
	var result []base.SingleSQL
	stream.Fill()
	tokens := stream.GetAllTokens()
	start := 0
	// Splitting multiple statements by semicolon symbol should consider the special case.
	// For CASE/REPLACE/IF/LOOP/WHILE/REPEAT statement, the semicolon symbol is not the end of the statement.
	// So we should skip the semicolon symbol in these statements.
	// These statements are begin with BEGIN/REPLACE/IF/LOOP/WHILE/REPEAT symbol and end with END/END REPLACE/END IF/END LOOP/END WHILE/END REPEAT symbol.
	// So this is a parenthesis matching problem.
	type openParenthesis struct {
		tokenType int
		pos       int
	}
	var stack []openParenthesis
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].GetTokenType() {
		case parser.MySQLParserBEGIN_SYMBOL:
			isBeginWork := getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserWORK_SYMBOL
			isBeginWork = isBeginWork || (getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserSEMICOLON_SYMBOL)
			isBeginWork = isBeginWork || (getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserEOF)
			if isBeginWork {
				continue
			}

			isXa := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserXA_SYMBOL
			if isXa {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserCASE_SYMBOL:
			isEndCase := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserEND_SYMBOL
			if isEndCase {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserIF_SYMBOL:
			isEndIf := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserEND_SYMBOL
			if isEndIf {
				continue
			}

			isIfExists := getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserEXISTS_SYMBOL
			if isIfExists {
				continue
			}

			isIfNotExists := (getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserNOT_SYMBOL ||
				getDefaultChannelTokenType(tokens, i, 1) == parser.MySQLParserNOT2_SYMBOL) &&
				getDefaultChannelTokenType(tokens, i, 2) == parser.MySQLParserEXISTS_SYMBOL
			if isIfNotExists {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserLOOP_SYMBOL:
			isEndLoop := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserEND_SYMBOL
			if isEndLoop {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserWHILE_SYMBOL:
			isEndWhile := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserEND_SYMBOL
			if isEndWhile {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserREPEAT_SYMBOL:
			isEndRepeat := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserUNTIL_SYMBOL
			if isEndRepeat {
				continue
			}

			stack = append(stack, openParenthesis{tokenType: tokens[i].GetTokenType(), pos: i})
		case parser.MySQLParserEND_SYMBOL:
			isXa := getDefaultChannelTokenType(tokens, i, -1) == parser.MySQLParserXA_SYMBOL
			if isXa {
				continue
			}

			// There are some special case for IF and REPEAT statement.
			// MySQL has two functions: IF(expr1,expr2,expr3) and REPEAT(str,count).
			// So we may meet single IF/REPEAT symbol without END IF/REPEAT symbol.
			// For these cases, we will see the END XXX symbol is not matched with the top of the stack.
			// We should skip these single IF/REPEAT symbol and backtracking these processes after IF/REPEAT symbol.
			if len(stack) == 0 {
				return nil, errors.New("invalid statement: failed to split multiple statements")
			}

			nextDefaultChannelTokenType := getDefaultChannelTokenType(tokens, i, 1)

			isEndIf := nextDefaultChannelTokenType == parser.MySQLParserIF_SYMBOL
			if isEndIf {
				if stack[len(stack)-1].tokenType != parser.MySQLParserIF_SYMBOL {
					// Backtracking the process.
					i = stack[len(stack)-1].pos
					stack = stack[:len(stack)-1]
					continue
				}
				stack = stack[:len(stack)-1]
				continue
			}

			isEndCase := nextDefaultChannelTokenType == parser.MySQLParserCASE_SYMBOL
			if isEndCase {
				if stack[len(stack)-1].tokenType != parser.MySQLParserCASE_SYMBOL {
					// Backtracking the process.
					i = stack[len(stack)-1].pos
					stack = stack[:len(stack)-1]
					continue
				}
				stack = stack[:len(stack)-1]
				continue
			}

			isEndLoop := nextDefaultChannelTokenType == parser.MySQLParserLOOP_SYMBOL
			if isEndLoop {
				if stack[len(stack)-1].tokenType != parser.MySQLParserLOOP_SYMBOL {
					// Backtracking the process.
					i = stack[len(stack)-1].pos
					stack = stack[:len(stack)-1]
					continue
				}
				stack = stack[:len(stack)-1]
				continue
			}

			isEndWhile := nextDefaultChannelTokenType == parser.MySQLParserWHILE_SYMBOL
			if isEndWhile {
				if stack[len(stack)-1].tokenType != parser.MySQLParserWHILE_SYMBOL {
					// Backtracking the process.
					i = stack[len(stack)-1].pos
					stack = stack[:len(stack)-1]
					continue
				}
				stack = stack[:len(stack)-1]
				continue
			}

			isEndRepeat := nextDefaultChannelTokenType == parser.MySQLParserREPEAT_SYMBOL
			if isEndRepeat {
				if stack[len(stack)-1].tokenType != parser.MySQLParserREPEAT_SYMBOL {
					// Backtracking the process.
					i = stack[len(stack)-1].pos
					stack = stack[:len(stack)-1]
					continue
				}
				stack = stack[:len(stack)-1]
				continue
			}

			// is BEGIN ... END or CASE .. END case
			leftTokenType := stack[len(stack)-1].tokenType
			if leftTokenType != parser.MySQLParserBEGIN_SYMBOL && leftTokenType != parser.MySQLParserCASE_SYMBOL {
				// Backtracking the process.
				i = stack[len(stack)-1].pos
				stack = stack[:len(stack)-1]
				continue
			}
			stack = stack[:len(stack)-1]
		case parser.MySQLParserSEMICOLON_SYMBOL:
			if len(stack) != 0 {
				continue
			}

			result = append(result, base.SingleSQL{
				Text:     stream.GetTextFromTokens(tokens[start], tokens[i]),
				BaseLine: tokens[start].GetLine() - 1,
				LastLine: tokens[i].GetLine()},
			)
			start = i + 1
		case parser.MySQLParserEOF:
			if len(stack) != 0 {
				// Backtracking the process.
				i = stack[len(stack)-1].pos
				stack = stack[:len(stack)-1]
				continue
			}

			if start <= i-1 {
				result = append(result, base.SingleSQL{
					Text:     stream.GetTextFromTokens(tokens[start], tokens[i-1]),
					BaseLine: tokens[start].GetLine() - 1,
					LastLine: tokens[i-1].GetLine()},
				)
			}
		}
	}
	return result, nil
}

func parseSingleStatement(statement string) (antlr.Tree, *antlr.CommonTokenStream, error) {
	input := antlr.NewInputStream(statement)
	lexer := parser.NewMySQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMySQLParser(stream)

	lexerErrorListener := &base.ParseErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)

	parserErrorListener := &base.ParseErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrorListener)

	p.BuildParseTrees = true

	tree := p.Script()

	if lexerErrorListener.Err != nil {
		return nil, nil, lexerErrorListener.Err
	}

	if parserErrorListener.Err != nil {
		return nil, nil, parserErrorListener.Err
	}

	return tree, stream, nil
}

func mysqlAddSemicolonIfNeeded(sql string) string {
	lexer := parser.NewMySQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()
	tokens := stream.GetAllTokens()
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i].GetChannel() != antlr.TokenDefaultChannel || tokens[i].GetTokenType() == parser.MySQLParserEOF {
			continue
		}

		// The last default channel token is a semicolon.
		if tokens[i].GetTokenType() == parser.MySQLParserSEMICOLON_SYMBOL {
			return sql
		}

		var result []string
		result = append(result, stream.GetTextFromInterval(antlr.NewInterval(0, tokens[i].GetTokenIndex())))
		result = append(result, ";")
		result = append(result, stream.GetTextFromInterval(antlr.NewInterval(tokens[i].GetTokenIndex()+1, tokens[len(tokens)-1].GetTokenIndex())))
		return strings.Join(result, "")
	}
	return sql
}

func parseInputStream(input *antlr.InputStream) ([]*ParseResult, error) {
	var result []*ParseResult
	lexer := parser.NewMySQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	list, err := splitMySQLStatement(stream)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		list[len(list)-1].Text = mysqlAddSemicolonIfNeeded(list[len(list)-1].Text)
	}

	for _, s := range list {
		tree, tokens, err := parseSingleStatement(s.Text)
		if err != nil {
			return nil, err
		}

		if isEmptyStatement(tokens) {
			continue
		}

		result = append(result, &ParseResult{
			Tree:     tree,
			Tokens:   tokens,
			BaseLine: s.BaseLine,
		})
	}

	return result, nil
}

func isEmptyStatement(tokens *antlr.CommonTokenStream) bool {
	for _, token := range tokens.GetAllTokens() {
		if token.GetChannel() == antlr.TokenDefaultChannel && token.GetTokenType() != parser.MySQLParserSEMICOLON_SYMBOL && token.GetTokenType() != parser.MySQLParserEOF {
			return false
		}
	}
	return true
}

// IsDelimiter returns true if the statement is a delimiter statement.
func IsDelimiter(stmt string) bool {
	delimiterRegex := `(?i)^\s*DELIMITER\s+`
	re := regexp.MustCompile(delimiterRegex)
	return re.MatchString(stmt)
}

// ExtractDelimiter extracts the delimiter from the delimiter statement.
func ExtractDelimiter(stmt string) (string, error) {
	delimiterRegex := `(?i)^\s*DELIMITER\s+(?P<DELIMITER>[^\s\\]+)\s*`
	re := regexp.MustCompile(delimiterRegex)
	matchList := re.FindStringSubmatch(stmt)
	index := re.SubexpIndex("DELIMITER")
	if index >= 0 && index < len(matchList) {
		return matchList[index], nil
	}
	return "", errors.Errorf("cannot extract delimiter from %q", stmt)
}

func hasDelimiter(statement string) (bool, []base.SingleSQL, error) {
	// use splitTiDBMultiSQL to check if the statement has delimiter
	t := tokenizer.NewTokenizer(statement)
	list, err := t.SplitTiDBMultiSQL()
	if err != nil {
		return false, nil, errors.Errorf("failed to split multi sql: %v", err)
	}

	for _, sql := range list {
		if IsDelimiter(sql.Text) {
			return true, list, nil
		}
	}

	return false, list, nil
}
