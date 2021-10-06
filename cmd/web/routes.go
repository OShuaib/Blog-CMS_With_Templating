package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *Application) routes(router *gin.Engine) *gin.Engine {
	//router.GET("/", app.Index)
	//router.Use(middlewares.CheckNotLogin())
	router.GET("/", app.Index)
	router.GET("/signup", middlewares.CheckNotLogin(), app.SignupPageHandler)
	router.POST("/signUp", middlewares.CheckNotLogin(), app.SignUpUser)
	router.GET("/login", middlewares.CheckNotLogin(), app.LoginPageHandler)
	router.POST("/login", middlewares.CheckNotLogin(), app.LoginUser)



	userRouter := router.Group("/user")
	userRouter.Use(middlewares.CheckLogin())
	{
		userRouter.GET("/logout", app.LogoutUser)
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
