package friend

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/model"
	"go-im/im-server/db"
)

func GetFriendsByAccount(account string) (friend []model.FriendInfo, err error) {
	if db.RedisClient == nil {
		fmt.Println("db redis client is nil")
		return
	}
	value, err := db.RedisClient.HGet("Friend", account).Result()
	if err != nil {
		fmt.Println("get Friend err, err =>", err)
		return
	}
	// friend=friend
	fmt.Println("get friend from redis , value =", value)

	err = json.Unmarshal([]byte(value), &friend)
	if err != nil {
		fmt.Println("json unmarshal friend err, err=", err)
		return
	}
	fmt.Printf("json unmarshal friend success, user=%v\n", friend)
	return
}
