package service

import (
	"math/rand"
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

func GetTbMMPicList(len int) []*utils.PicRes {
	resList := make([]*utils.PicRes, len, len)

	resChannel := make(chan int8)
	for i := 0; i < len; i++ {
		i := i
		go func() {
			resList[i] = GetOneMMPicInCycle()
			resChannel <- 0
		}()
	}
	for i := 0; i < len; i++ {
		<-resChannel
	}
	return resList
}
