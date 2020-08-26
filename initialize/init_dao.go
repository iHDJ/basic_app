package initialize

import (
	"basic_app/conf"
	"basic_app/dao"
	"context"
	"errors"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDao() (err error) {
	var (
		config = conf.Conf
		db     *gorm.DB
		rds    *redis.Client
	)

	//load mysql
	db, err = gorm.Open("mysql", connUrl(config.MySQL))
	if err != nil {
		return fmt.Errorf("连接到MySQL失败 %e", err)
	}

	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	//load redis
	rds = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	if err = rds.Ping(context.Background()).Err(); err != nil {
		return errors.New("Redis 发送Ping包无响应")
	}

	dao.Current, err = dao.New(db, rds)
	return
}

func connUrl(conf conf.MySQL) string {
	return fmt.Sprintf(
		"%s:%s@(%s)/%s?%s",
		conf.Username,
		conf.Password,
		conf.Address,
		conf.Database,
		conf.Config,
	)
}
