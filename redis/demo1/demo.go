package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDb *redis.Client

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "35.201.213.247:6379",
		Password: "",
		DB:       0,
	})

	str, err := redisDb.Ping().Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(str)
	return
}
func redisDemo() {
	err := redisDb.Set("name", "张三", 0).Err()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	val, err := redisDb.Get("name").Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(val)
	//取一个不存在的键值错误处理
	val2, err := redisDb.Get("age").Result()
	if err == redis.Nil {
		fmt.Println("this Key is not ")
	} else if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(val2)
}
func redisZsetDemo() (err error) {
	zsetkey := "rank"
	languages := []redis.Z{
		redis.Z{Score: 99, Member: "PHP"},
		redis.Z{Score: 98, Member: "Golang"},
		redis.Z{Score: 96, Member: "C"},
		redis.Z{Score: 92, Member: "Python"},
		redis.Z{Score: 82, Member: "JS"},
	}
	//添加元素组到redis某个key中
	n, err := redisDb.ZAdd(zsetkey, languages...).Result() //返回值是一共多少元素添加到了key中
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(n)
	//对某个数据增加
	editNum, err := redisDb.ZIncrBy(zsetkey, 20, "Golang").Result() //返回值是修改后的结果
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(editNum)
	//取分数最高的3个
	res, err := redisDb.ZRevRangeWithScores(zsetkey, 0, 2).Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range res {
		fmt.Println(v.Member, v.Score)
	}

	//取区间值
	op := redis.ZRangeBy{
		Min: "90",
		Max: "150",
	}
	rest, err := redisDb.ZRangeByScoreWithScores(zsetkey, op).Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range rest {
		fmt.Println(v.Member, v.Score)
	}
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// redisDemo()

	redisZsetDemo()

}
