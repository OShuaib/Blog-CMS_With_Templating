package interfaces

import "github.com/Ad3bay0c/BlogCMS/cmd/models"

type Commentable interface {
	ViewCommentByPostId(postId string) ([]models.Comment, error)
	CreateComment(comment models.Comment) error
}

type User interface {
	GetUserByEmail(email string) (bool, models.User)
	SaveUser(user *models.User) (string, error)
	GetUserById(id string) (models.User, error)
}

type Blogger interface {
	SavePost(post models.Post) error
	GetPostsByUserId(userId string) ([]models.Post, error)
	GetAllPost(userId string) ([]models.Post, error)
	UpdatePost(post models.Post) error
	ViewBlogPostById(postId string) (models.Post, error)
	DeletePostById(postId string, userId string) error
	IncrementViews(postId string, views int) error
}