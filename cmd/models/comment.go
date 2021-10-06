package models

import (
	"database/sql"
	"fmt"
	"github.com/Ad3bay0c/BlogCMS/pkg/postgresql"
	"log"
)

const (
	COMMENT_TABLE = "comments"
)
type Comment struct {
	ID   			string 	`json:"id,omitempty"`
	UserId			string	`json:"user_id,omitempty"`
	PostId			string	`json:"post_id,omitempty"`
	Comment			string	`json:"comment,omitempty"`
	CreatedAt 		int64	`json:"created_at,omitempty"`
	UpdatedAt		int64	`json:"updated_at,omitempty"`
}

type CommentModel struct {
	DB *sql.DB
}

func (model *CommentModel) ViewCommentByPostId(postId string) ([]Comment, error){
	var comments []Comment

	rows, err := model.DB.Query(fmt.Sprintf("SELECT id, comment, user_id, created_at FROM %s WHERE post_id = $1", COMMENT_TABLE), postId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.ID, &comment.Comment, &comment.UserId, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (model *CommentModel) CreateComment(comment Comment) error {
	stmt, err := model.DB.Prepare(fmt.Sprintf("INSERT INTO %s (id, comment, post_id, created_at, updated_at, user_id) VALUES($1, $2, $3, $4, $5, $6);", COMMENT_TABLE))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.ID, comment.Comment, comment.PostId, comment.CreatedAt, comment.UpdatedAt, comment.UserId)
	if err != nil {
		return err
	}
	return nil
}
func CountComment(postId string) int {
	db, _ := postgresql.ConnectDb()
	row := db.QueryRow(fmt.Sprintf("SELECT count(*) as count FROM %s WHERE post_id = $1", COMMENT_TABLE), postId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Printf("%v", err.Error())
		return 0
	}
	return count
}