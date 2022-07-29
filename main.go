package main

import (
	_ "toubiaogo/boot"
	_ "toubiaogo/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
