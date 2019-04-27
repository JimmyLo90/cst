package aimo

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"zhyq132/cst/config"
	"zhyq132/cst/model"
	"zhyq132/cst/model/shop"
)

type WssMessage struct {
	TimeStamp   uint32 `json:"timestamp"` //[int64]秒级时间戳
	Capture     string `json:"capture"`   //[string]图⽚片base64编码
	CaptureRect struct {
		X      float32 `json:"x"`      //人脸检测到的框x坐标
		Y      float32 `json:"y"`      //人脸检测到的框y坐标
		Witdh  float32 `json:"width"`  //人脸检测到的框宽度
		Height float32 `json:"height"` //人脸检测到的框高度
	} `json:"capture_rect"`
	DeviceType  string `json:"device_type"`  //[string]抓拍源类型,此字段⽤用于区分不不同的抓拍设备类型，"box":抓拍盒的视频源
	DeviceSN    string `json:"device_sn"`    //[string]设备序列号
	SourceToken string `json:"source_token"` //[string]视频源句柄
	Recognize   struct {
		Match struct {
			FaceSetToken string  `json:"face_set_token"` //[string]人脸集合token
			FaceSetName  string  `json:"face_set_name"`  //[string]人脸集合名
			FaceToken    string  `json:"face_token"`     //[string]人脸token
			FaceUserMark string  `json:"face_user_mark"` //[string]人脸用户自定义信息
			Score        float32 `json:"score"`          //[float32]比对分数(0-1)
		} `json:"match"`
	} `json:"recognize"`
	Attributes struct {
		Age    uint8 `json:"age"`    //[int32]年龄
		Gender uint8 `json:"gender"` //[int32]1男，2女
	} `json:"attributes"`
}

//把[]byte 数据转换为struct数据,存储数到mysql
func (m *WssMessage) Stored(d []byte) (uint32, error) {
	err := json.Unmarshal(d, m)
	if err != nil {
		return 0, err
	}

	db, err := model.GormOpenDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	aimoFacePic := m.structToModel()
	if err := db.Table(aimoFacePic.TableName()).Create(&aimoFacePic).Error; err != nil {
		return 0, err
	}

	//保存base64到文件
	if savePath, err := m.saveCaptureToJpg(); err != nil {
		log.Printf("stored message err:%s,%v \n", err, savePath)
	} else {
		if err := db.Table(aimoFacePic.TableName()).Model(&aimoFacePic).UpdateColumn("save_path", savePath).Error; err != nil {
			log.Printf("stored message err:%v \n", err)
		}
	}

	return aimoFacePic.AimoFacePicID, nil
}

func (m WssMessage) structToModel() shop.AimoFacePic {
	aimoFacePic := shop.AimoFacePic{}

	aimoFacePic.IsAddToFacesets = 0
	aimoFacePic.TimeStamp = m.TimeStamp
	aimoFacePic.CaptureRectX = strconv.FormatFloat(float64(m.CaptureRect.X), 'f', -1, 32)
	aimoFacePic.CaptureRectY = strconv.FormatFloat(float64(m.CaptureRect.Y), 'f', -1, 32)
	aimoFacePic.CaptureRectWidth = strconv.FormatFloat(float64(m.CaptureRect.Witdh), 'f', -1, 32)
	aimoFacePic.CaptureRectHeight = strconv.FormatFloat(float64(m.CaptureRect.Height), 'f', -1, 32)
	aimoFacePic.DeviceType = m.DeviceType
	aimoFacePic.DeviceSn = m.DeviceSN
	aimoFacePic.RtspSourceToken = m.SourceToken
	aimoFacePic.FaceSetsToken = m.Recognize.Match.FaceSetToken
	aimoFacePic.FaceSetsName = m.Recognize.Match.FaceSetName
	aimoFacePic.FaceToken = m.Recognize.Match.FaceToken
	aimoFacePic.FaceUserMark = m.Recognize.Match.FaceUserMark
	aimoFacePic.Score = uint32(m.Recognize.Match.Score * 100)
	aimoFacePic.Age = m.Attributes.Age
	aimoFacePic.Gender = m.Attributes.Gender
	aimoFacePic.IsDel = 0

	return aimoFacePic
}

func (m WssMessage) saveCaptureToJpg() (string, error) {
	b, err := base64.StdEncoding.DecodeString(m.Capture)
	if err != nil {
		return "", err
	}

	path := config.Config.Aimo.PicSavePath
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModeDir)
		return "", err
	}

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.FormatInt(int64(rand.Uint64()), 10)
	extType := ".jpg"

	uri := path + fileName + extType

	f, err := os.Create(uri)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return "", err
	}

	return uri, nil
}
