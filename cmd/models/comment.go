package models

import (
	"database/sql"
	"fmt"
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
