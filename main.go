//go:generate go-bindata -o=./assets/assets.go -pkg=assets conf/config.yaml locales/...
package main

import (
	"basic_app/initialize"
	"basic_app/router"
)

func main() {
	if err := initialize.Initialize(); err != nil {
		panic(err.Error())
	}

	router.LoadRouter(":3000")
}
