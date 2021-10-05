package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *Application) routes(router *gin.Engine) *gin.Engine {
	//router.GET("/", app.Index)
	router.GET("/", app.Index)
	router.GET("/signup", app.SignupPageHandler)
	router.POST("/signUp", app.SignUpUser)
	router.GET("/login", app.LoginPageHandler)
	router.POST("/login", app.LoginUser)

	router.GET("/logout", app.LogoutUser)

	userRouter := router.Group("/user")

	{
		userRouter.GET("/", app.GetUser)
		userRouter.GET("/viewBlogPost", app.ViewMyBlogPost)
	}

	blogRouter := router.Group("/blog")
	blogRouter.Use(middlewares.CheckLogin())
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
