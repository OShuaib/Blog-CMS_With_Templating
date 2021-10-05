package main

import (
	"flag"
	"github.com/Ad3bay0c/BlogCMS/cmd/interfaces"
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/Ad3bay0c/BlogCMS/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Application struct {
	infoLog			*log.Logger
	errorLog		*log.Logger
	userModel		interfaces.User
	postModel		interfaces.Blogger
}

func main() {
	addr := flag.String("addr", ":3500", "Enter the New Port")
	flag.Parse()

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.SetOutput(file)
	errorLog := log.New(os.Stderr, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.SetOutput(file)

	db, err := postgresql.ConnectDb()
	if err != nil {
		errorLog.Println(err.Error())
		panic(err)
	}

	app := &Application{
		infoLog:  infoLog,
		errorLog: errorLog,
		userModel: &models.UserModel{DB: db},
		postModel: &models.PostModel{DB: db},
	}

	r := gin.Default()

	r.LoadHTMLGlob("./ui/html/*")
	r.StaticFS("static", http.Dir("./ui/static/"))

	router := app.routes(r)

	server := &http.Server{
		ErrorLog: nil,
		Handler: router,
		Addr: *addr,
	}

	log.Printf("Database Connected\n")
	log.Printf("server started at port%v\n", *addr)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (app *Application) ServerError(c *gin.Context, err error) {
	app.errorLog.Printf(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
}

func (app *Application) InfoError(c *gin.Context, message string, code int) {
	app.infoLog.Println(message)
	c.JSON(code, gin.H{"message": message})
}