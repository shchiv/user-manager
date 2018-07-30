package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mikespook/gorbac"
	"github.com/users-manager/src/constants"
	"github.com/users-manager/src/models"
	"github.com/users-manager/src/router"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "gotest:gotest@tcp(db)/test")
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}
	defer db.Close()

	db.CreateTable(models.User{})
	db.Create(&models.User{Name: "admin", Role: constants.AdminRole})
	db.Create(&models.User{Name: "user", Role: constants.UserRole})

	authManager := gorbac.New()
	if err := createAccessControlRules(authManager); err != nil {
		log.Fatalf("Can't give access to the users")
	}

	server := router.NewServer()
	server.Start(db, authManager)
}

func createAccessControlRules(rbac *gorbac.RBAC) error {
	adminRole := gorbac.NewStdRole(constants.AdminRole)
	userRole := gorbac.NewStdRole(constants.UserRole)

	if err := adminRole.Assign(createAllPermissions); err != nil {
		return err
	}

	if err := adminRole.Assign(createUserPermissions); err != nil {
		return err
	}

	if err := adminRole.Assign(readAllPermissions); err != nil {
		return err
	}

	if err := adminRole.Assign(readUserPermissions); err != nil {
		return err
	}

	if err := adminRole.Assign(deletePermissions); err != nil {
		return err
	}

	if err := userRole.Assign(createUserPermissions); err != nil {
		return err
	}

	if err := userRole.Assign(readUserPermissions); err != nil {
		return err
	}

	if err := rbac.Add(adminRole); err != nil {
		return err
	}

	if err := rbac.Add(userRole); err != nil {
		return err
	}
	return nil
}
