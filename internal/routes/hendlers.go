package routes

import "github.com/qulDev/jwt-gin-gorm/internal/handler"

type Handlers struct {
	User    *handler.UserHandler
	Article *handler.ArticleHandler
}
