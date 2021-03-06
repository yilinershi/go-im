package user

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/model"
	"go-im/im-server/db"
)



func GetUserByAccount(account string) (user *model.User, err error) {
	if db.RedisClient == nil {
		fmt.Println("db redis client is nil")
		return
	}
	value, err := db.RedisClient.HGet("User", account).Result()
	if err != nil {
		fmt.Println("get user err, err=", err)
		return
	}
	user = new(model.User)
	fmt.Println("get user from redis , value =", value)
	err = json.Unmarshal([]byte(value), user)
	if err != nil {
		fmt.Println("json unmarshal user err, err=", err)
		return
	}
	fmt.Printf("json unmarshal user success, user=%v\n", user)
	return
}

func AddUser(account string, user *model.User) (err error) {
	if db.RedisClient == nil {
		fmt.Println("db redis client is nil")
		return
	}

	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		fmt.Println("db redis client is nil")
		return 
	}

	_, dbErr := db.RedisClient.HSet("User", account, string(jsonData)).Result()
	if dbErr != nil {
		fmt.Println("get user err, err=", dbErr)
		return
	}
	return nil
}


