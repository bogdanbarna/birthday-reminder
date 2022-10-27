package router

import (
	//"bogdanbarna/hello-rest-gin/repository"
	"bogdanbarna/hello-rest-gin/controller"

	"github.com/gin-gonic/gin"
)

/*
Notes:
- Person data model in model.go
- If `GET`, only `Form` binding engine (`query`) used.
- If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
- See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
*/

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()

	router.GET("/user", controller.GetPersons)
	router.GET("/user/:username", controller.GetPerson)
	router.GET("/user/:username/birthday", controller.GetPersonBirthday)
	router.PUT("/user/:username", controller.PutUser)
	router.POST("/user/", controller.PostUser)
	router.PATCH("/user/:username", controller.PatchUser)
	router.DELETE("/user/:username", controller.DeleteUser)

	return router
}
