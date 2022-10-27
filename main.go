package main

import (
	"bogdanbarna/hello-rest-gin/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
