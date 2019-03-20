package mongo

import "gopkg.in/mgo.v2/bson"

//WeixinOpen 授权后的c端小程序数据
type WeixinOpen struct {
	Id                             bson.ObjectId    `bson:"_id,omitempty"`
	AreaId                         int32            `bson:"a_areaId"`
	AuthorizerAppId                string           `bson:"authorizer_appid"`
	AuthorizerAccessToken          string           `bson:"authorizer_access_token"`
	AuthorizerAccessTokenExpiresIn int32            `bson:"authorizer_access_token_expires_in"`
	AuthorizerRefreshTokenstring   string           `bson:"authorizer_refresh_token"`
	FuncInfo                       []int32          `bson:"func_info"`
	NickName                       string           `bson:"nick_name"`
	HeadImg                        string           `bson:"head_img"`
	ServiceTypeInfo                int32            `bson:"service_type_info"`
	VerifyTypeInfo                 int32            `bson:"verify_type_info"`
	UserName                       string           `bson:"user_name"`
	AliasName                      string           `bson:"alias"`
	OauthStatus                    int32            `bson:"oauth_status"`
	OauthTime                      int32            `bson:"oauth_time"`
	JsapiTicket                    string           `bson:"jsapi_ticket"`
	JsapiTicketExpiresIn           int32            `bson:"jsapi_ticket_expires_in"`
	LastSetIndustryTime            int32            `bson:"last_set_industry_time"`
	QrcodeUrl                      string           `bson:"qrcode_url"`
	UnauthorizedCreateTime         int32            `bson:"unauthorizedCreateTime"`
	Signature                      string           `bson:"signature"`
	PrincipalName                  string           `bson:"principal_name"`
	BusinessInfo                   map[string]int32 `bson:"business_info"`
}

func (w WeixinOpen) DBName() string {
	return "w_center"
}

func (w WeixinOpen) TableName() string {
	return "weixin_open"
}

func (w *WeixinOpen) GetAuthorizerAccessToken() string {
	return w.AuthorizerAccessToken
}
