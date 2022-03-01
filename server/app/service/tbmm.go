package service

import (
	"log"
	"math/rand"
	"tbmm/app/store"
	"tbmm/pkg/utils"
	"time"
)

const (
	UomgApi = iota
	XiaoApi
)

var randInUnix = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetOneMMPicInCycle() (picRes *utils.PicRes) {
	var (
		res *utils.PicRes
		err error
	)
	r := randInUnix.Int()
	curKey := r % (XiaoApi + 1)
	switch curKey {
	case UomgApi:
		res, err = utils.FindOneByUomg()
	case XiaoApi:
		res, err = utils.FindOneByXiaoApi()
	}
	if err != nil {
		return GetOneMMPicInCycle()
	}
	return res
}

func loadTbMMPicList() {
	num, _ := store.GetPicSurplusNum()
	if num < 100 {
		log.Printf("缓存数量不足，当前：%d", num)
		resChannel := make(chan *utils.PicRes)
		for i := 0; i < 5; i++ {
			go func() {
				resChannel <- GetOneMMPicInCycle()
			}()
		}
		for i := 0; i < 5; i++ {
			_, err := store.InsertPic(<-resChannel)
			if err != nil {
				log.Printf("插入出错%v", err)
			}
		}
		loadTbMMPicList()
	} else {
		time.Sleep(1 * time.Second)
		loadTbMMPicList()
	}
}

func init() {
	go loadTbMMPicList()
}
func GetTbMMPicList(len int) []*utils.PicRes {
	resList := make([]*utils.PicRes, len, len)
	resChannel := make(chan *utils.PicRes)
	num, _ := store.GetPicSurplusNum()
	if num > int64(len) {
		for i := 0; i < len; i++ {
			go func() {
				onePic, err := store.LoadOnePic()
				if err != nil {
					log.Printf("缓存查询失败%v", err)
					resChannel <- GetOneMMPicInCycle()
				} else {
					resChannel <- onePic
					log.Printf("消费缓存内容%v", onePic)
				}
			}()
		}
	} else {
		for i := 0; i < len; i++ {
			go func() {
				resChannel <- GetOneMMPicInCycle()
			}()
		}

	}
	for i := 0; i < len; i++ {
		resList[i] = <-resChannel
	}
	return resList
}
