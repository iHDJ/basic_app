package main

import (
	"basic_app/conf"
	"basic_app/initialize"
	"basic_app/seed"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func seedMain() {
	if err := initialize.InitConfig(); err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open("mysql", connUrl(conf.Conf.MySQL))
	if err != nil {
		panic(err.Error())
	}

	for _, seedHandle := range []func(db *gorm.DB) (err error){
		seed.SeedGroup,
		seed.SeedManagerData,
	} {

		if err = seedHandle(db); err != nil {
			panic(err.Error())
		}

	}
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
