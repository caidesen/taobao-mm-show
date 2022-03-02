package service

import (
	"log"
	"math"
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

func loadTbMMPicList(t *time.Ticker) {
	num, _ := store.GetPicSurplusNum()
	if num < 100 {
		log.Printf("缓存数量不足，当前：%d", num)
		t.Stop()
		cycleNum := int(math.Ceil(float64(100-int(num)) / 5))
		for r := 0; r < cycleNum; r++ {
			resChannel := make(chan *utils.PicRes, 5)
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
		}
		num, _ := store.GetPicSurplusNum()
		log.Printf("补充完成，当前缓存数量：%d", num)
		t.Reset(time.Second * 2)
	}
}

func init() {
	t := time.NewTicker(time.Second * 2)
	go func() {
		for {
			select {
			case <-t.C:
				loadTbMMPicList(t)
			}
		}
	}()
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
