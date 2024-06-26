<template>
  <div class="space-y-3 divide-y pb-4 px-2">
    <div
      v-for="(row, i) in tableRows"
      :key="i"
      class="pt-3 first:pt-2 space-y-2"
    >
      <div class="flex items-center space-x-3">
        <div
          class="relative w-5 h-5 flex flex-shrink-0 items-center justify-center rounded-full select-none"
          :class="statusIconClass(row.checkResult.status)"
        >
          <template
            v-if="row.checkResult.status === PlanCheckRun_Result_Status.SUCCESS"
          >
            <heroicons-solid:check class="w-4 h-4" />
          </template>
          <template
            v-if="row.checkResult.status === PlanCheckRun_Result_Status.WARNING"
          >
            <heroicons-outline:exclamation class="h-4 w-4" />
          </template>
          <template
            v-else-if="
              row.checkResult.status === PlanCheckRun_Result_Status.ERROR
            "
          >
            <span class="text-white font-medium text-base" aria-hidden="true">
              !
            </span>
          </template>
        </div>
        <div v-if="showCategoryColumn" class="shrink-0">
          {{ row.category }}
        </div>
        <div class="font-semibold">{{ row.title }}</div>
      </div>
      <div class="textinfolabel">
        <span>{{ row.checkResult.content }}</span>
        <template v-if="row.checkResult.sqlReviewReport?.detail">
          <span
            class="ml-1 normal-link"
            @click="
              state.activeResultDefinition =
                row.checkResult.sqlReviewReport!.detail
            "
            >{{ $t("sql-review.view-definition") }}</span
          >
          <span class="border-r border-control-border ml-1"></span>
        </template>
        <template
          v-if="
            row.checkResult.sqlReviewReport &&
            getActiveRule(row.checkResult.title)
          "
        >
          <span
            class="ml-1 normal-link"
            @click="setActiveRule(row.checkResult.title)"
            >{{ $t("sql-review.rule-detail") }}</span
          >
          <span class="border-r border-control-border ml-1"></span>
        </template>
        <template v-if="row.checkResult.sqlSummaryReport">
          {{ row.checkResult.sqlSummaryReport.affectedRows }}
        </template>

        <HideInStandaloneMode>
          <a
            v-if="row.link"
            class="ml-1 normal-link"
            :href="row.link.url"
            :target="row.link.target"
          >
            {{ row.link.title }}
          </a>
        </HideInStandaloneMode>

        <!-- Only show the error line for latest plan check run -->
        <template v-if="isLatest && row.checkResult.sqlReviewReport?.line">
          <span class="border-r border-control-border ml-1"></span>
          <span
            class="ml-1 normal-link"
            @click="
              handleClickPlanCheckDetailLine(
                row.checkResult.sqlReviewReport!.line
              )
            "
          >
            L{{ row.checkResult.sqlReviewReport.line }}
          </span>
        </template>

        <slot name="row-extra" :row="row" />
      </div>
    </div>
  </div>

  <SQLRuleEditDialog
    v-if="state.activeRule"
    :editable="false"
    :rule="state.activeRule"
    :disabled="true"
    @cancel="state.activeRule = undefined"
  />

  <ResultDefinitionModal
    v-if="state.activeResultDefinition"
    :definition="state.activeResultDefinition"
    @cancel="state.activeResultDefinition = undefined"
  />
</template>

<script setup lang="ts">
import { computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { SQLRuleEditDialog } from "@/components/SQLReview/components";
import { WORKSPACE_ROUTE_SQL_REVIEW } from "@/router/dashboard/workspaceRoutes";
import { useReviewPolicyByEnvironmentName } from "@/store";
import type { RuleTemplate } from "@/types";
import {
  GeneralErrorCode,
  SQLReviewPolicyErrorCode,
  UNKNOWN_ENVIRONMENT_NAME,
  findRuleTemplate,
  getRuleLocalization,
  ruleTemplateMap,
} from "@/types";
import { convertPolicyRuleToRuleTemplate } from "@/types";
import type { PlanCheckRun } from "@/types/proto/v1/plan_service";
import {
  PlanCheckRun_Result,
  PlanCheckRun_Result_Status,
  PlanCheckRun_Status,
} from "@/types/proto/v1/plan_service";
import ResultDefinitionModal from "./ResultDefinitionModal.vue";

interface ErrorCodeLink {
  title: string;
  target: string;
  url: string;
}

export type PlanCheckDetailTableRow = {
  checkResult: PlanCheckRun_Result;
  category: string;
  title: string;
  link: ErrorCodeLink | undefined;
};

type LocalState = {
  activeRule?: RuleTemplate;
  activeResultDefinition?: string;
};

const props = defineProps<{
  planCheckRun: PlanCheckRun;
  environment?: string;
  isLatest?: boolean;
}>();

const { t } = useI18n();
const router = useRouter();
const state = reactive<LocalState>({
  activeRule: undefined,
  activeResultDefinition: undefined,
});

const emit = defineEmits<{
  (event: "close"): void;
}>();

const statusIconClass = (status: PlanCheckRun_Result_Status) => {
  switch (status) {
    case PlanCheckRun_Result_Status.SUCCESS:
      return "bg-success text-white";
    case PlanCheckRun_Result_Status.WARNING:
      return "bg-warning text-white";
    case PlanCheckRun_Result_Status.ERROR:
      return "bg-error text-white";
  }
};

const checkResultList = computed((): PlanCheckRun_Result[] => {
  if (props.planCheckRun.status === PlanCheckRun_Status.DONE) {
    return props.planCheckRun.results;
  } else if (props.planCheckRun.status === PlanCheckRun_Status.FAILED) {
    return [
      PlanCheckRun_Result.fromPartial({
        status: PlanCheckRun_Result_Status.ERROR,
        title: t("common.error"),
        content: props.planCheckRun.error,
      }),
    ];
  } else if (props.planCheckRun.status === PlanCheckRun_Status.CANCELED) {
    return [
      PlanCheckRun_Result.fromPartial({
        status: PlanCheckRun_Result_Status.WARNING,
        title: t("common.canceled"),
        content: props.planCheckRun.error,
      }),
    ];
  }

  return [];
});

const categoryAndTitle = (
  checkResult: PlanCheckRun_Result
): [string, string] => {
  if (checkResult.sqlReviewReport) {
    const code = checkResult.sqlReviewReport?.code ?? checkResult.code;
    if (!code) {
      return ["", checkResult.title];
    }
    if (code === SQLReviewPolicyErrorCode.EMPTY_POLICY) {
      const title = messageWithCode(checkResult.title, code);
      return ["", title];
    }
    const rule = ruleTemplateMap.get(checkResult.title);
    if (rule) {
      const ruleLocalization = getRuleLocalization(rule.type);
      const key = `sql-review.category.${rule.category.toLowerCase()}`;
      const category = t(key);
      const title = messageWithCode(ruleLocalization.title, code);
      return [category, title];
    }
    return ["", messageWithCode(checkResult.title, code)];
  }
  if (checkResult.sqlSummaryReport) {
    if (typeof checkResult.sqlSummaryReport.affectedRows === "number") {
      return ["", t("task.check-type.affected-rows")];
    }
    return ["", checkResult.title];
  }

  return ["", checkResult.title];
};

const messageWithCode = (message: string, code: number) => {
  return `${message} #${code}`;
};

const errorCodeLink = (
  checkResult: PlanCheckRun_Result
): ErrorCodeLink | undefined => {
  const code = checkResult.sqlReviewReport?.code ?? checkResult.code;
  switch (code) {
    case undefined:
      return;
    case GeneralErrorCode.OK:
      return;
    case SQLReviewPolicyErrorCode.EMPTY_POLICY:
      return {
        title: t("sql-review.configure-policy"),
        target: "__blank",
        url: router.resolve({
          name: WORKSPACE_ROUTE_SQL_REVIEW,
        }).fullPath,
      };
    default: {
      const errorCodeNamespace =
        checkResult.sqlReviewReport !== undefined ? "advisor" : "core";
      const domain = "https://www.bytebase.com";
      const path = `/docs/reference/error-code/${errorCodeNamespace}/`;
      const query = `source=console#${code}`;
      const url = `${domain}${path}?${query}`;
      return {
        title: t("common.view-doc"),
        target: "__blank",
        url: url,
      };
    }
  }
};

const tableRows = computed(() => {
  return checkResultList.value.map<PlanCheckDetailTableRow>((checkResult) => {
    const [category, title] = categoryAndTitle(checkResult);
    const link = errorCodeLink(checkResult);
    return {
      checkResult,
      category,
      title,
      link,
    };
  });
});

const showCategoryColumn = computed((): boolean =>
  tableRows.value.some((row) => row.category !== "")
);

const reviewPolicy = useReviewPolicyByEnvironmentName(
  computed(() => {
    return props.environment || UNKNOWN_ENVIRONMENT_NAME;
  })
);
const getActiveRule = (type: string): RuleTemplate | undefined => {
  const rule = reviewPolicy.value?.ruleList.find((rule) => rule.type === type);
  if (!rule) {
    return undefined;
  }

  const ruleTemplate = findRuleTemplate(type);
  if (!ruleTemplate) {
    return undefined;
  }

  return convertPolicyRuleToRuleTemplate([rule], ruleTemplate);
};
const setActiveRule = (type: string) => {
  state.activeRule = getActiveRule(type);
};

const handleClickPlanCheckDetailLine = (line: number) => {
  window.location.hash = `L${line}`;
  emit("close");
};
</script>
