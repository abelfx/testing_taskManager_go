package main

import (
	"restfulapi/data"
	"restfulapi/router"
)



func main() {
	data.InitDB()
	r := router.SetupRouter()
	r.Run()
}