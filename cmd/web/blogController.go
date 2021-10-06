package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Create a new Blog Post
func (app *Application) CreateBlogPost(c *gin.Context) {
	var blogPost models.Post
	userId, _ := c.Get("userId")
	id := userId.(string)
	//err := c.BindJSON(&blogPost)

	blogPost.Title = strings.TrimSpace(c.PostForm("title"))
	blogPost.Details = strings.TrimSpace(c.PostForm("details"))
	access := c.PostForm("access")
	blogPost.Access, _ = strconv.Atoi(access)

	if blogPost.Title == "" || blogPost.Details == "" {
		m["Message"] = "Title/Details cannot be empty"
		m["Color"] = "danger"
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Title/Details cannot be empty"})
		c.Redirect(302, "/user/my-blog-posts")
		return
	}

	blogPost.ID = uuid.New().String()
	blogPost.UserId = id
	blogPost.CreatedAt = time.Now().Unix()
	blogPost.UpdatedAt = time.Now().Unix()

	err := app.postModel.SavePost(blogPost)
	if err != nil {
		app.errorLog.Printf("%v", err.Error())
		m["Message"] = "OOPS!!!, Something Went Wrong"
		m["Color"] = "danger"
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Title/Details cannot be empty"})
		c.Redirect(302, "/user/my-blog-posts")
		return
	}
	m["Message"] = "Post Created Successfully"
	m["Color"] = "success"
	c.Redirect(302, "/user/my-blog-posts")
}
// View All Public Posts in the database
func (app *Application) ViewAllPosts(c *gin.Context) {
	param, _ := c.Get("userId")
	user_id := param.(string)

	posts, err := app.postModel.GetAllPost(user_id)
	if err != nil {
		app.errorLog.Printf("%v", err)

	}
	var check bool
	if len(posts) > 0 {
		check = true
	}
	app.Render(c, "blog.page.html", user_id, gin.H{"Check": check, "Post": posts})
	//c.HTML(http.StatusOK, "blog.page.html", gin.H{"Check": check, "Post": posts})
}

// View all the Post for a particular user
func (app *Application) ViewMyBlogPost(c *gin.Context) {
	param, _ := c.Get("userId")
	user_id := param.(string)

	posts, err := app.postModel.GetPostsByUserId(user_id)
	if err != nil {
		app.ServerError(c, err)
		return
	}
	m["Data"] = posts
	//c.JSON(http.StatusOK, gin.H{"data": posts})
	app.Render(c, "create.page.html", user_id, m)
	m["Message"] = ""
	m["Color"] = ""
	delete(m, "Post")
}
func (app *Application) editBlogPostPage(c *gin.Context) {

}

func (app *Application) updateBlogPost(c *gin.Context) {
	param, _ := c.Get("userId")
	user_id := param.(string)

	post_id := c.Param("id")

	var blogPost models.Post

	err := c.BindJSON(&blogPost)
	if err != nil {
		app.ServerError(c, err)
		return
	}
	if blogPost.Title == "" || blogPost.Details == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Title/Details cannot be empty"})
		return
	}

	blogPost.ID = post_id
	blogPost.UserId = user_id
	blogPost.UpdatedAt = time.Now().Unix()

	err = app.postModel.UpdatePost(blogPost)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error Updating Value"})
		app.errorLog.Printf("%v", err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "Updated Successfully"})
}

func (app *Application) ViewPostById(c *gin.Context) {
	//param, _ := c.Get("userId")
	//userId := param.(string)

	postId := c.Param("id")

	post, err := app.postModel.ViewBlogPostById(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		app.errorLog.Printf("%v\n",err.Error())
		return
	}
	c.JSON(200, gin.H{"data": post})
}

func (app *Application) DeleteBlogPost(c *gin.Context) {
	param, _ := c.Get("userId")
	userId := param.(string)

	postId := c.Param("id")

	err := app.postModel.DeletePostById(postId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		app.errorLog.Printf("%v", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}