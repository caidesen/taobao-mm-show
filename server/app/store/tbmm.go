package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"tbmm/pkg/utils"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}

func GetPicSurplusNum() (len int64, err error) {
	return rdb.LLen(ctx, "TBMM_PIC").Result()
}

func LoadOnePic() (*utils.PicRes, error) {
	result, err := rdb.BLPop(ctx, 10*time.Second, "TBMM_PIC").Result()
	if err != nil {
		return nil, err
	}
	var data utils.PicRes
	err = json.Unmarshal([]byte(result[1]), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func InsertPic(picRes *utils.PicRes) (int64, error) {
	marshal, err := json.Marshal(picRes)
	if err != nil {
		return 0, err
	}
	return rdb.LPush(ctx, "TBMM_PIC", string(marshal)).Result()
}
