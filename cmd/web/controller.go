package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/helpers"
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

var m = map[string]interface{}{}

func (app *Application) Index(c *gin.Context) {
	log.Print(os.Getenv("SESSION_ID"))
	c.Redirect(302, "/login")
}

func (app *Application) SignupPageHandler(c *gin.Context) {

	c.HTML(200, "signup.page.html", m)
	m["Message"] = ""
	m["Color"] = ""
}

func (app *Application) LoginPageHandler(c *gin.Context) {

	c.HTML(200, "login.page.html", m)
	m["Message"] = ""
	m["Color"] = ""
}

func (app *Application) SignUpUser(c *gin.Context) {
	var user = &models.User{}

	user.Email = c.PostForm("email")
	user.FirstName = c.PostForm("firstname")
	user.LastName = c.PostForm("lastname")
	user.Password = c.PostForm("password")
	user.Confirm = c.PostForm("confirm")

	if user.FirstName == "" || user.LastName == "" {
		m["Message"] = "First Name or Last Name is required"
		m["Color"] = "danger"
		//c.JSON(http.StatusBadRequest, gin.H{"message": "First Name or Last Name is required"})
		c.Redirect(http.StatusFound, "/signup")
		return
	}
	if user.Password == "" || len(user.Password) < 4{
		m["Message"] = "Password cannot be less than 4"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be less than 4"})
		return
	}
	if !helpers.ValidateEmail(user.Email) {
		m["Message"] = "Invalid Email"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		return
	}
	if user.Password != user.Confirm {
		m["Message"] = "Password and Confirm Password must match"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Password and Confirm Password must match"})
		return
	}
	if ok, _ := app.userModel.GetUserByEmail(user.Email); ok {
		m["Message"] = "Email Already Exist"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusForbidden, gin.H{"message": "Email Already Exist"})
		return
	}
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	newHash, err := helpers.GeneratePassword(user.Password)
	if err != nil {
		app.errorLog.Printf(err.Error())
		m["Message"] = "Oops!!!, Something went wrong"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	user.Password = newHash

	_, err = app.userModel.SaveUser(user)

	if err != nil {
		app.errorLog.Printf(err.Error())
		m["Message"] = "Oops!!!, Something went wrong"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	app.infoLog.Printf("Registration Successful")
	m["Message"] = "Registered Successfully, Please Login"
	m["Color"] = "success"
	c.Redirect(http.StatusFound, "/signup")
}

func (app *Application) LoginUser(c *gin.Context) {
	app.errorLog.Printf("%v", os.Getenv("SESSION_ID"))
	var user = &models.User{}

	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	if !helpers.ValidateEmail(user.Email) {
		m["Message"] = "Invalid Email"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/login")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		return
	}
	ok, userM := app.userModel.GetUserByEmail(user.Email)
	if !ok {
		m["Message"] = "Email Does Not Exist"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/login")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		return
	}
	if !helpers.ComparePassword(user.Password, userM.Password) {
		m["Message"] = "Incorrect Email/Password"
		m["Color"] = "danger"
		c.Redirect(http.StatusFound, "/login")
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		return
	}
	c.SetCookie("session", userM.ID, 60*60, "/", "", true, true)
	m["Message"] = "Logged In Successfully"
	m["Color"] = "success"
	c.Redirect(http.StatusFound, "/")
	//c.JSON(200, gin.H{"message" : "Logged In Successfully", "token": token})
}

func (app *Application) LogoutUser(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", true, true)
	c.Redirect(http.StatusFound, "/login")
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
