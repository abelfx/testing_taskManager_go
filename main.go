package main

import (
	"restfulapi/router"
)



func main() {
	r := router.SetupRouter()
	r.Run()
}