package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"tides-server/pkg/restapi/operations/user"

	"tides-server/pkg/config"
	"tides-server/pkg/models"
)


func RegisterOrgHandler(params user.RegisterUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()
	var queryUser models.User
	db.Where("username = ?", body.Username).First(&queryUser)
	if queryUser.Username != "" {
		return user.NewRegisterUserBadRequest().WithPayload(&user.RegisterUserBadRequestBody{Message: "Username already used!"})
	}

	newUser := models.Org{
		City:        body.City,
		CompanyName: body.CompanyName,
		Country:     body.Country,
		Email:       body.Email,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Password:    body.Password,
		Phone:       body.Phone,
		Position:    body.Position,
		Priority:    models.UserPriorityLow,
		Username:    body.Username,
	}

	err := db.Create(&newUser).Error
	if err != nil {
		return user.NewRegisterUserBadRequest()
	}

	res := &user.RegisterUserOKBodyUserInfo{
		City:        body.City,
		CompanyName: body.CompanyName,
		Country:     body.Country,
		Email:       body.Email,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Password:    body.Password,
		Phone:       body.Phone,
		Position:    body.Position,
		Priority:    models.UserPriorityLow,
		Username:    body.Username,
	}

	return user.NewRegisterUserOK().WithPayload(&user.RegisterUserOKBody{UserInfo: res})
}

