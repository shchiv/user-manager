package permissions

import (
	"github.com/mikespook/gorbac"
	"github.com/users-manager/src/constants"
)

var (
	CreateAll  = gorbac.NewStdPermission(constants.CreateAll)
	CreateUser = gorbac.NewStdPermission(constants.CreateUser)
	ReadAll    = gorbac.NewStdPermission(constants.ReadAll)
	ReadUser   = gorbac.NewStdPermission(constants.ReadUser)
	Delete     = gorbac.NewStdPermission(constants.Delete)
)
