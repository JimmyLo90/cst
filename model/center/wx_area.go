package center

type WxArea struct {
	AreaID     int32 `orm:"column(id);pk"`
	WeixinType int32 `orm:"column(weixinType)"`
}

func (w WxArea) DBName() string {
	return "w_center"
}

func (w WxArea) TableName() string {
	return "wx_area"
}
