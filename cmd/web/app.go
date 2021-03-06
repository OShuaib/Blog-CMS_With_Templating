package main

import (
	"flag"
	"github.com/Ad3bay0c/BlogCMS/cmd/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Application struct {
	infoLog			*log.Logger
	errorLog		*log.Logger
	userModel		interfaces.User
	postModel		interfaces.Blogger
	commentModel	interfaces.Commentable
}

type Message struct {
	Msg  string
}
func main() {
	addr := flag.String("addr", ":3500", "Enter the New Port")
	flag.Parse()

	r := gin.Default()

	r.LoadHTMLGlob("./ui/html/*")
	r.StaticFS("static", http.Dir("./ui/static/"))

	app := ApplicationSetUp()

	router := app.routes(r)

	server := &http.Server{
		ErrorLog: nil,
		Handler: router,
		Addr: *addr,
	}

	log.Printf("Database Connected\n")
	log.Printf("server started at port%v\n", *addr)
	
	err := server.ListenAndServe()
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

func (app *Application) Render(c *gin.Context, page string, userId string, data interface{}) {
	c.SetCookie("session", userId, 60*30, "/", "", true, true)
	c.HTML(200, page, data)
}