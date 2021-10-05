package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/helpers"
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (app *Application) Index(c *gin.Context) {
	c.JSON(200, gin.H{"Message" : "This is the Index Page"})
}

func (app *Application) LoginPageHandler(c *gin.Context) {
	c.HTML(200, "login.page.html", nil)
}
func (app *Application) SignUpUser(c *gin.Context) {
	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "An error Occurred"})
		return
	}
	if user.FirstName == "" || user.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "First Name or Last Name is required"})
		return
	}
	if user.Password == "" || len(user.Password) < 4{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be less than 4"})
		return
	}
	if !helpers.ValidateEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		return
	}
	if user.Password != user.Confirm {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password and Confirm Password must match"})
		return
	}
	if ok, _ := app.userModel.GetUserByEmail(user.Email); ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Email Already Exist"})
		return
	}
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	newHash, err := helpers.GeneratePassword(user.Password)
	if err != nil {
		app.errorLog.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	user.Password = newHash

	id, err := app.userModel.SaveUser(user)
	if err != nil {
		app.errorLog.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	token, err := helpers.GenerateToken(id)

	if err != nil {
		app.errorLog.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	app.infoLog.Printf("Registered In Successfully")
	c.JSON(201, gin.H{"message": "Registered Successfully", "token": token})
}

func (app *Application) LoginUser(c *gin.Context) {

	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		app.ServerError(c, err)
		return
	}
	if !helpers.ValidateEmail(user.Email) {
		app.InfoError(c, "Invalid Email", http.StatusBadRequest)
		return
	}
	ok, userM := app.userModel.GetUserByEmail(user.Email)
	if !ok {
		app.InfoError(c, "Email Does Not Exist", http.StatusBadRequest)
		return
	}
	if !helpers.ComparePassword(user.Password, userM.Password) {
		app.InfoError(c, "Incorrect Email/Password", http.StatusForbidden)
		return
	}
	token, err := helpers.GenerateToken(userM.ID)
	if err != nil {
		app.ServerError(c, err)
		return
	}
	app.infoLog.Printf("Logged In Successfully")
	c.JSON(200, gin.H{"message" : "Logged In Successfully", "token": token})
}

func (app *Application) GetUser(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		app.errorLog.Printf("Server Error")
		c.JSON(500, gin.H{"message": "Server Error"})
		return
	}
	newId := id.(string)
	user, err := app.userModel.GetUserById(newId)
	if err != nil {
		app.ServerError(c, err)
	}

	c.JSON(200, gin.H{"data" : user})
}
