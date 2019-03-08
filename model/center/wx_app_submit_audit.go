package center

type WxAppSubmitAudit struct {
	ID int32 `orm:"column(id)"`
	AreaID int32 `orm:"column(areaid)"`
	AuthAppID string `orm:"column(auth_appid)"`
	AuthUserName string `orm:"column(auth_user_name)"`
	VersionID int32 `orm:"column(version_id)"`
	Dateline int32 `orm:"column(dateline)"`
	AuditID string `orm:"column(auditid)"`
	Status int32 `orm:"column(status)"`
	Remark string `orm:"column(remark)"`
	AuditTime int32 `orm:"column(audit_time)"`
	IsVisible int32 `orm:"column(is_visible)"`
	VisibleTime int32 `orm:"column(visible_time)"`
}

func(w WxAppSubmitAudit) DBName() string{
	return "w_center"
}

func (w WxAppSubmitAudit) TableName() string{
	return "wx_app_submit_audit"
}


