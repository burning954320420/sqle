package v2

import (
	"net/http"

	"github.com/actiontech/sqle/sqle/api/controller"
	v1 "github.com/actiontech/sqle/sqle/api/controller/v1"
	"github.com/actiontech/sqle/sqle/model"
	"github.com/labstack/echo/v4"
)

type GetAuditPlanSQLsReqV2 struct {
	PageIndex uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize  uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlanSQLsResV2 struct {
	controller.BaseRes
	Data      AuditPlanSQLResV2 `json:"data"`
	TotalNums uint64            `json:"total_nums"`
}

type AuditPlanSQLResV2 struct {
	Head []AuditPlanSQLHeadV2                 `json:"head"`
	Rows []map[string] /* head name */ string `json:"rows"`
}

type AuditPlanSQLHeadV2 struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type,omitempty" enums:"sql"`
}

// @Summary 获取指定审核计划的SQLs信息(不包括审核结果)
// @Description get audit plan SQLs
// @Id getAuditPlanSQLsV2
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v2.GetAuditPlanSQLsResV2
// @router /v2/audit_plans/{audit_plan_name}/sqls [get]
func GetAuditPlanSQLs(c echo.Context) error {
	s := model.GetStorage()

	req := new(GetAuditPlanSQLsReqV2)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}

	apName := c.Param("audit_plan_name")
	err := v1.CheckCurrentUserCanAccessAuditPlan(c, apName)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	var offset uint32
	if req.PageIndex >= 1 {
		offset = req.PageSize * (req.PageIndex - 1)
	}

	data := map[string]interface{}{
		"audit_plan_name": apName,
		"limit":           req.PageSize,
		"offset":          offset,
	}
	auditPlanSQLs, count, err := s.GetAuditPlanSQLsByReq(data)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	auditPlanSQLResV2 := AuditPlanSQLResV2{
		Head: []AuditPlanSQLHeadV2{
			{
				Name: "fingerprint",
				Desc: "SQL指纹",
				Type: "sql",
			},
			{
				Name: "last_receive_sql_text",
				Desc: "最后一次匹配到该指纹的语句",
				Type: "sql",
			},
			{
				Name: "counter",
				Desc: "匹配到该指纹的语句数量",
			},
			{
				Name: "last_receive_timestamp",
				Desc: "最后一次匹配到该指纹的时间",
			},
		},
	}
	auditPlanSQLResV2.Rows = make([]map[string]string, 0, len(auditPlanSQLs))

	for _, auditPlanSQL := range auditPlanSQLs {
		auditPlanSQLResV2.Rows = append(auditPlanSQLResV2.Rows, map[string]string{
			"fingerprint":            auditPlanSQL.Fingerprint,
			"last_receive_sql_text":  auditPlanSQL.LastReceiveText,
			"counter":                auditPlanSQL.Counter,
			"last_receive_timestamp": auditPlanSQL.LastReceiveTimestamp,
		})
	}
	return c.JSON(http.StatusOK, &GetAuditPlanSQLsResV2{
		BaseRes:   controller.NewBaseReq(nil),
		Data:      auditPlanSQLResV2,
		TotalNums: count,
	})
}

type GetAuditPlanReportSQLsReqV2 struct {
	PageIndex uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize  uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlanReportSQLsResV2 struct {
	controller.BaseRes
	Data      []AuditPlanReportSQLResV2 `json:"data"`
	TotalNums uint64                    `json:"total_nums"`
}

type AuditPlanReportSQLResV2 struct {
	SQL         string `json:"audit_plan_report_sql" example:"select * from t1 where id = 1"`
	AuditResult string `json:"audit_plan_report_sql_audit_result" example:"same format as task audit result"`
}

// @Summary 获取指定审核计划的SQL审核详情
// @Description get audit plan report SQLs
// @Id getAuditPlanReportSQLsV2
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param audit_plan_report_id path string true "audit plan report id"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v2.GetAuditPlanReportSQLsResV2
// @router /v2/audit_plans/{audit_plan_name}/report/{audit_plan_report_id}/ [get]
func GetAuditPlanReportSQLs(c echo.Context) error {
	s := model.GetStorage()

	req := new(GetAuditPlanReportSQLsReqV2)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}

	apName := c.Param("audit_plan_name")
	err := v1.CheckCurrentUserCanAccessAuditPlan(c, apName)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	var offset uint32
	if req.PageIndex >= 1 {
		offset = req.PageSize * (req.PageIndex - 1)
	}

	data := map[string]interface{}{
		"audit_plan_name":      apName,
		"audit_plan_report_id": c.Param("audit_plan_report_id"),
		"limit":                req.PageSize,
		"offset":               offset,
	}
	auditPlanReportSQLs, count, err := s.GetAuditPlanReportSQLsByReq(data)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	auditPlanReportSQLsResV2 := make([]AuditPlanReportSQLResV2, len(auditPlanReportSQLs))
	for i, auditPlanReportSQL := range auditPlanReportSQLs {
		auditPlanReportSQLsResV2[i] = AuditPlanReportSQLResV2{
			SQL:         auditPlanReportSQL.LastReceiveText,
			AuditResult: auditPlanReportSQL.AuditResult,
		}
	}
	return c.JSON(http.StatusOK, &GetAuditPlanReportSQLsResV2{
		BaseRes:   controller.NewBaseReq(nil),
		Data:      auditPlanReportSQLsResV2,
		TotalNums: count,
	})
}
