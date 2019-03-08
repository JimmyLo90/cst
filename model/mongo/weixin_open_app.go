package mongo

import "gopkg.in/mgo.v2/bson"

//WeixinOpenApp 授权后的c端小程序数据
type WeixinOpenApp struct {
	Id                             bson.ObjectId `bson:"_id,omitempty"`
	AreaId                         int32         `bson:"a_areaId"`
	AuthorizerAppId                string        `bson:"authorizer_appid"`
	AuthorizerAccessToken          string        `bson:"authorizer_access_token"`
	AuthorizerAccessTokenExpiresIn int32         `bson:"authorizer_access_token_expires_in"`
	AuthorizerRefreshTokenstring   string        `bson:"authorizer_refresh_token"`
	FuncInfo                       []int32       `bson:"func_info"`
	NickName                       string        `bson:"nick_name"`
	HeadImg                        string        `bson:"head_img"`
	ServiceTypeInfo                int32         `bson:"service_type_info"`
	VerifyTypeInfo                 int32         `bson:"verify_type_info"`
	UserName                       string        `bson:"user_name"`
	AliasName                      string        `bson:"alias"`
	OauthStatus                    int32         `bson:"oauth_status"`
	JsapiTicketExpiresIn           int32         `bson:"jsapi_ticket_expires_in"`
	QrcodeUrl                      string        `bson:"qrcode_url"`
	OauthTime                      int32         `bson:"oauth_time"`
}

func (w WeixinOpenApp) DBName() string {
	return "w_center"
}

func (w WeixinOpenApp) TableName() string {
	return "weixin_open_app"
}

func (w WeixinOpenApp) GetAuthorizerAccessToken() string{
	return w.AuthorizerAccessToken
}

//WeixinOpenPlatform 第三方的数据：AreaId为0时的那一条数据
type WeixinOpenPlatform struct {
	Id                              bson.ObjectId `bson:"_id,omitempty"`
	AreaId                          int32         `bson:"a_areaId"`
	ComponentVerifyTicket           string        `bson:"componentVerifyTicket"`
	ComponentVerifyTicketCreateTime int32         `bson:"componentVerifyTicketCreateTime"`
	ComponentAccessTokenExpiresIn   int32         `bson:"component_access_token_expires_in"`
	ComponentAccessToken            string        `bson:"component_access_token"`
	PreAuthCodeExpiresIn            int32         `bson:"pre_auth_code_expires_in"`
	PreAuthCode                     string        `bson:"pre_auth_code"`
}

func (w WeixinOpenPlatform) DBName() string {
	return "w_center"
}

func (w WeixinOpenPlatform) TableName() string {
	return "weixin_open_app"
}

func (w *WeixinOpenPlatform) GetComponentVerifyTicket() string {
	return w.ComponentVerifyTicket
}
