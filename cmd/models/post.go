package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)
const (
	POST_TABLE = "posts"
)
type Post struct {
	ID   			string 	`json:"id,omitempty"`
	UserId			string	`json:"user_id,omitempty"`
	UserFirstName	string	`json:"user_first_name"`
	UserLastName	string	`json:"user_last_name"`
	Title			string	`json:"title,omitempty"`
	Details			string	`json:"details,omitempty"`
	Access			int		`json:"access,omitempty"`
	CreatedAt 		int64	`json:"created_at,omitempty"`
	UpdatedAt		int64	`json:"updated_at,omitempty"`
	Comments		int		`json:"comments,omitempty"`
}

type PostModel struct {
	DB *sql.DB
}
func (model *PostModel) SavePost(post Post) error {
	stmt, err := model.DB.Prepare(fmt.Sprintf("INSERT INTO %s (id, title, details, access, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);",POST_TABLE))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.ID, post.Title, post.Details, post.Access, post.UserId, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (model *PostModel) GetPostsByUserId(userId string) ([]Post, error) {
	rows, err := model.DB.Query(fmt.Sprintf("SELECT id, title, details, access, created_at FROM %s WHERE user_id = $1 ORDER BY created_at DESC", POST_TABLE), userId)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Details, &post.Access, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (model *PostModel) GetAllPost(userId string) ([]Post, error) {
	rows, err := model.DB.Query(fmt.Sprintf("SELECT p.id, p.title, p.details, p.access, p.created_at, p.user_id, u.firstname, u.lastname FROM %s p INNER JOIN %s u ON p.user_id = u.id WHERE p.access = 0 ORDER BY p.created_at DESC;", POST_TABLE, TABLE))
	if err != nil {
		return nil, err
	}
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Details, &post.Access, &post.CreatedAt, &post.UserId, &post.UserFirstName, &post.UserLastName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (model *PostModel) UpdatePost(post Post) error {
	stmt, err := model.DB.Prepare(fmt.Sprintf("UPDATE %s SET title = $1, details = $2, access = $3, updated_at = $4 WHERE id = $5 AND user_id = $6", POST_TABLE))
	if err != nil {
		return err
	}
	res, err := stmt.Exec(post.Title, post.Details, post.Access, post.UpdatedAt, post.ID, post.UserId)
	if err != nil {
		return err
	}
	total, _ := res.RowsAffected()
	if total < 1 {
		return errors.New("error Updating Row")
	}
	return nil
}

func (model *PostModel) ViewBlogPostById(postId string) (Post, error) {
	var post Post
	row := model.DB.QueryRow(fmt.Sprintf("SELECT id, title, details, user_id, created_at, access FROM %s WHERE id = $1", POST_TABLE), postId)
	err := row.Scan(&post.ID, &post.Title, &post.Details, &post.UserId, &post.CreatedAt, &post.Access)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (model *PostModel) DeletePostById(postId string, userId string) error {
	stmt, err  := model.DB.Prepare(fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2;", POST_TABLE))
	if err != nil {
		return err
	}
	res, err := stmt.Exec(userId, postId)
	if err != nil {
		return err
	}
	total, _ := res.RowsAffected()
	if total < 1 {
		return errors.New("error Updating Row")
	}
	return nil
}

func (p Post) FormatTime(t int64) string {
	return time.Unix(t, 0).Format("Jan 02, 2006")
}

//func (p Post) GetUser() string {
//	return fmt.Sprintf("%s...", detail[:3])
//}