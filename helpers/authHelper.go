package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, role string)(err error){
	userType:= ctx.GetString("user_type")
	err= nil
	if userType!=role {
		err = errors.New("Unauthorized to access this resource")
		return
	}
	return
}
// i understand the end goal but i dont understand how this function accomplishes it
func MathcUserTypeToUid(ctx *gin.Context, userId string)(err error)  {
	userType:= ctx.GetString("user_type")
	uid:= ctx.GetString("uid")
	err=nil

	if userType=="USER" && uid!=userId{
		err = errors.New("Unauthorized to access this resource")
		return
	}
	err = CheckUserType(ctx, userType)
	return
}