package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v9"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitMysql() {
	// 自定义日志打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "/r/n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.hostIp"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("mysql initialization succeeded.")
	}
}

const PublishKey = "websocket"

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	//ctx := context.Context()
	//Rdb.Ping(ctx)
	//_, err := Rdb.Ping().Result()
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("redis initialization succeeded.")
	//}
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := RDB.Publish(ctx, channel, msg).Err()
	fmt.Println("Publish: ", msg)
	return err
}

// Subscribe 订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RDB.Subscribe(ctx, channel)
	fmt.Println("sub......", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println("sub err.....", err)
	} else {
		fmt.Println("subscribe: ", msg.Payload)
	}
	return msg.Payload, err
}
