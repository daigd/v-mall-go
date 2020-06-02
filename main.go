package main

import (
	"github.com/daigd/v-mall-go/app"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app.InitIris()
}
