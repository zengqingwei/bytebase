package databricks

import (
	"context"
	"database/sql"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/db"

	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
	v1pb "github.com/bytebase/bytebase/proto/generated-go/v1"
)

func init() {
	db.Register(storepb.Engine_DATABRICKS, NewDatabricksDriver)
}

var _ db.Driver = (*Driver)(nil)

type Driver struct {
	curCatalog string
	client     *databricks.WorkspaceClient
}

func NewDatabricksDriver(db.DriverConfig) db.Driver {
	return &Driver{}
}

// Each Databricks driver is associated with a single Databricks Workspace (Workspace -> catalog -> schema -> table).
func (d *Driver) Open(_ context.Context, _ storepb.Engine, config db.ConnectionConfig) (db.Driver, error) {
	databricksConfig := &databricks.Config{
		Host: config.Host,
	}
	// Support Databricks native authentication.
	// ref: https://github.com/databricks/databricks-sdk-go?tab=readme-ov-file#databricks-native-authentication
	if config.AuthenticationPrivateKey != "" {
		// Token.
		databricksConfig.Token = config.AuthenticationPrivateKey
	} else {
		// Basic username and password.
		databricksConfig.Username = config.Username
		databricksConfig.Password = config.Password
		databricksConfig.AccountID = config.AccountID
	}
	client, err := databricks.NewWorkspaceClient(databricksConfig)
	if err != nil {
		return nil, err
	}
	d.client = client
	d.curCatalog = config.Database
	return d, nil
}

func (*Driver) Close(_ context.Context) error {
	return nil
}

func (d *Driver) Ping(ctx context.Context) error {
	_, err := d.client.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{})
	if err != nil {
		return errors.Wrapf(err, "failed to ping instance")
	}
	return nil
}

func (*Driver) GetDB() *sql.DB {
	return nil
}

func (*Driver) GetType() storepb.Engine {
	return storepb.Engine_DATABRICKS
}

func (*Driver) QueryConn(_ context.Context, _ *sql.Conn, _ string, _ *db.QueryContext) ([]*v1pb.QueryResult, error) {
	return nil, nil
}

func (*Driver) RunStatement(_ context.Context, _ *sql.Conn, _ string) ([]*v1pb.QueryResult, error) {
	return nil, nil
}

func (*Driver) Execute(_ context.Context, _ string, _ db.ExecuteOptions) (int64, error) {
	return 0, nil
}

func (*Driver) CheckSlowQueryLogEnabled(_ context.Context) error {
	return nil
}
