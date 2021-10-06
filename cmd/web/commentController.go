package main

import (
	"fmt"
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func (app *Application) CreateComment(c *gin.Context){
	param, _ := c.Get("userId")
	userId := param.(string)

	id := c.Param("id")

	if strings.TrimSpace(c.PostForm("comment")) == "" {
		m["Message"] = "Comment is Required"
		m["Color"] = "danger"
		c.Redirect(302, fmt.Sprintf("/blog/%s/", id))
		return
	}
	comment := models.Comment{
		ID:      uuid.New().String(),
		Comment: strings.TrimSpace(c.PostForm("comment")),
		PostId:  id,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		UserId: userId,
	}

	err := app.commentModel.CreateComment(comment)

	if err != nil {
		app.errorLog.Printf(err.Error())
		m["Message"] = http.StatusText(http.StatusInternalServerError)
		m["Color"] = "danger"
		c.Redirect(302, fmt.Sprintf("/blog/%s/", id))

		return
	}

	m["Message"] = "Comment Created Successfully"
	m["Color"] = "success"

	c.Redirect(http.StatusFound, fmt.Sprintf("/blog/%s/", id))
}
