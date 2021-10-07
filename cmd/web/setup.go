package main

import (
	"github.com/Ad3bay0c/BlogCMS/cmd/models"
	"github.com/Ad3bay0c/BlogCMS/pkg/postgresql"
	"log"
	"os"
)

func ApplicationSetUp() *Application {
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
		commentModel: &models.CommentModel{DB: db},
	}
	return app
}
