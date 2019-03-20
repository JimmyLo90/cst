package tongji

type TjMpNewsSend struct {
	ID               int32  `orm:"column(id)"`
	AreaID           int32  `orm:"column(area_id)"`
	RefDate          string `orm:"column(ref_date)"`
	MsgID            string `orm:"column(msgid)"`
	Title            string `orm:"column(title)"`
	IntPageReadUser  int32  `orm:"column(int_page_read_user)"`
	IntPageReadCount int32  `orm:"column(int_page_read_count)"`
	OriPageReadUser  int32  `orm:"column(ori_page_read_user)"`
	OriPageReadCount int32  `orm:"column(ori_page_read_count)"`
	ShareUser        int32  `orm:"column(share_user)"`
	ShareCount       int32  `orm:"column(share_count)"`
	AddToFavUser     int32  `orm:"column(add_to_fav_user)"`
	AddToFavCount    int32  `orm:"column(add_to_fav_count)"`
	UserSource       string `orm:"column(user_source)"`
	Dateline         int32  `orm:"column(dateline)"`
}

func (t TjMpNewsSend) DBName() string {
	return "tongji"
}
func (t TjMpNewsSend) TableName() string {
	return "tj_mpnews_send"
}