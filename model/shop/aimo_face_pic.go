package shop

type AimoFacePic struct {
	AimoFacePicID     uint32 `gorm:"column:aimo_face_pic_id;type:int(11);PRIMARY_KEY;AUTO_INCREMENT;"`
	IsAddToFacesets   uint8  `gorm:"column:is_add_to_facesets;type:tinyint(4);DEFAULT:0;NOT NULL;"` //该照片是否已经添加进人脸底库
	TimeStamp         uint32 `gorm:"column:timestamp;type:int(11);NOT NULL;"`                       //秒级时间戳
	CaptureRectX      string `gorm:"column:capture_rect_x;type:varchar(32);NOT NULL;"`              //人脸x轴
	CaptureRectY      string `gorm:"column:capture_rect_y;type:varchar(32);NOT NULL;"`              //人脸y轴
	CaptureRectWidth  string `gorm:"column:capture_rect_width;type:varchar(32);NOT NULL;"`          //人脸高度
	CaptureRectHeight string `gorm:"column:capture_rect_height;type:varchar(32);NOT NULL;"`         //人脸宽度
	DeviceType        string `gorm:"column:device_type;type:varchar(127);NOT NULL;"`                //图像来源类型
	DeviceSn          string `gorm:"column:device_sn;type:varchar(255);NOT NULL;"`                  //拍摄该图像的摄像头所属box的设备码
	RtspSourceToken   string `gorm:"column:rtsp_source_token;type:varchar(255);NOT NULL;"`          //摄像头的唯一标示
	FaceSetsToken     string `gorm:"column:face_sets_token;type:varchar(255);NOT NULL;"`            //人脸底库token
	FaceSetsName      string `gorm:"column:face_sets_name;type:varchar(127);NOT NULL;"`             //人脸底库备注名
	FaceToken         string `gorm:"column:face_token;type:varchar(255);NOT NULL;"`                 //每张图片的唯一标示
	FaceUserMark      string `gorm:"column:face_user_mark;type:varchar(127);NOT NULL;"`             //识别出人脸后的关联标示
	Score             uint32 `gorm:"column:score;type:int(11);NOT NULL;"`                           //人脸对比分熟，乘以100以后的整形
	Age               uint8  `gorm:"column:age;type:tinyint(4);NOT NULL;"`                          //年龄
	Gender            uint8  `gorm:"column:gender;type:tinyint(4);NOT NULL;"`                       //性别1男；2女
	SavePath          string `gorm:"column:save_path;type:varchar(255);NOT NULL;"`                  //图片的保存路径
	IsDel             uint8  `gorm:"column:is_del;type:tinyint(4);NOT NULL;"`                       //是否删除
}

func (a AimoFacePic) DBName() string {
	return "4s_wx_db"
}
func (a AimoFacePic) TableName() string {
	return "aimo_face_pic"
}
