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
		userRouter.GET("/my-blog-posts", app.ViewMyBlogPost)
	}

	blogRouter := router.Group("/blog")
	blogRouter.Use(middlewares.CheckLogin())
	{
		blogRouter.POST("/create", app.CreateBlogPost)
		blogRouter.GET("/", app.ViewAllPosts)
		blogRouter.GET("/my-blogs", app.ViewAllPosts)

		subRouter := blogRouter.Group("/:id")
		subRouter.GET("/edit-page", app.editBlogPostPage)
		subRouter.POST("/update", app.updateBlogPost)
		subRouter.GET("/delete", app.DeleteBlogPost)
		subRouter.GET("/", app.ViewPostById)

		commentRouter := subRouter.Group("/comment")
		{
			commentRouter.POST("/create", app.CreateComment)
			commentRouter.DELETE("/:commentId")
			commentRouter.PUT("/:commentId")
		}
		likeRouter := subRouter.Group("/like")
		{
			likeRouter.GET("/", app.LikeAPost)
			likeRouter.GET("/unlike", app.UnLikeAPost)
		}
	}
	return router
}
