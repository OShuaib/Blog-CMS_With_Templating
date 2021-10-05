package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/helpers"
	"github.com/gin-gonic/gin"
)

func (app *Application) routes(router *gin.Engine) *gin.Engine {
	//router.GET("/", app.Index)
	router.GET("/signup", app.SignupPageHandler)
	router.POST("/signup", app.SignUpUser)
	router.GET("/login", app.LoginPageHandler)
	router.POST("/login", app.LoginUser)

	userRouter := router.Group("/user")
	userRouter.Use(helpers.VerifyTokenMiddleware()) // A middleware to check for token validity
	{
		userRouter.GET("/", app.GetUser)
		userRouter.GET("/viewBlogPost", app.ViewMyBlogPost)
	}

	blogRouter := router.Group("/blog")
	blogRouter.Use(helpers.VerifyTokenMiddleware())
	{
		blogRouter.POST("/create", app.CreateBlogPost)
		blogRouter.GET("/", app.ViewAllPosts)

		subRouter := blogRouter.Group("/:id")
		subRouter.PUT("/update", app.updateBlogPost)
		subRouter.DELETE("/", app.DeleteBlogPost)
		subRouter.GET("/", app.ViewPostById)

		commentRouter := subRouter.Group("/comment")
		{
			commentRouter.POST("/")
			commentRouter.DELETE("/:commentId")
			commentRouter.PUT("/:commentId")
		}
	}
	return router
}
