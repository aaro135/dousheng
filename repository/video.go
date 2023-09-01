package repository

import (
	"dousheng/util"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	Id         int64  `gorm:"column:video_id"`
	Playurl    string `gorm:"column:play_url"`
	Coverurl   string `gorm:"column:cover_url"`
	Title      string `gorm:"column:title"`
	CreateTime string `gorm:"column:create_time"`
	UserId     int64  `gorm:"column:user_id"`
}

func (v *Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

func (*VideoDao) VideoPrepare(latest int64) ([]*Video, error) {
	var videoList []*Video = make([]*Video, 0, 20)
	tm := time.Unix(latest/1000, 0)
	err := db.Model(&Video{}).Where("create_time < ?", tm).Order("create_time desc").Find(&videoList).Error
	fmt.Println("length of video list", len(videoList))
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("prepare video error" + err.Error())
	}
	return videoList, nil
}

// create new video in db and video_id + 1
func (*VideoDao) VideoCreate(newVideo Video) (*Video, error) {
	var oldVideo Video
	db.Last(&oldVideo)
	newVideo.Id = oldVideo.Id + 1
	err := db.Create(&newVideo).Error
	if err != nil {
		util.Logger.Error("create video error:" + err.Error())
		return nil, err
	}
	return &newVideo, nil
}

func (*VideoDao) SearchVideoById(userId int64) ([]*Video, error) {
	var videoList []*Video
	err := db.Model(&Video{}).Where("user_id =?", userId).Find(&videoList).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("search user videoList error" + err.Error())
	}
	return videoList, nil
}

func (*VideoDao) SearchUserByVideo(videoId int64) (*Video, error) {
	var video Video
	err := db.Model(&Video{}).Where("video_id =?", videoId).First(&video).Error
	if err != nil {
		util.Logger.Error("search video author error:" + err.Error())
		return nil, err
	}
	return &video, nil
}

func CountLike(videoId int64) (int64, error) {
	var count int64
	err := db.Model(&Video{}).Where("video_id =?", videoId).Count(&count).Error
	if err != nil {
		util.Logger.Error("search video count error:" + err.Error())
		return -1, err
	}
	return count, nil
}
