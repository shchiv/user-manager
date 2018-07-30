package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mikespook/gorbac"
	"log"
)

const (
	getAllURL = "/getAll"
	userURL   = "/user"

	authHeader = "Authorization"
)

var (
	dbManager   *gorm.DB
	authManager *gorbac.RBAC
)

type server struct{}

func NewServer() *server {
	return new(server)
}

func (s *server) Start(db *gorm.DB, rbac *gorbac.RBAC) {
	dbManager = db
	authManager = rbac

	r := gin.Default()
	r.GET(getAllURL, getAllUsersHandler)
	r.POST(userURL, addUserHandler)
	r.DELETE(userURL, removeUserHandler)

	log.Fatal(r.Run())
}
