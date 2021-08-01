package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"

	"tides-server/pkg/restapi/operations/user"

	"tides-server/pkg/config"
	"tides-server/pkg/models"
)

// RegisterUserHandler is API handler for /users/register POST
func RegisterUserHandler(params user.RegisterUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()
	var queryUser models.User
	db.Where("username = ?", body.Username).First(&queryUser)
	if queryUser.Username != "" {
		return user.NewRegisterUserBadRequest().WithPayload(&user.RegisterUserBadRequestBody{Message: "Username already used!"})
	}

	newUser := models.User{
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

// UserLoginHandler is API handler for /users/login POST
func UserLoginHandler(params user.UserLoginParams) middleware.Responder {
	body := params.ReqBody

	db := config.GetDB()
	var queryUser models.User
	db.Where("Username = ?", body.Username).First(&queryUser)
	if queryUser.Username == "" {
		return user.NewUserLoginUnauthorized()
	} else if queryUser.Password != body.Password {
		return user.NewUserLoginUnauthorized()
	}

	expirationTime := time.Now().Add(expireTime)
	claims := Claims{
		ID: queryUser.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := config.GetConfig().SecretKey
	signedToken, _ := token.SignedString([]byte(secretKey))

	res := user.UserLoginOKBodyUserInfo{Priority: queryUser.Priority, Username: queryUser.Username,
		OrgName: queryUser.OrgName, PwReset: fmt.Sprintf("%t", queryUser.PwReset), Role: queryUser.Role}

	return user.NewUserLoginOK().WithPayload(&user.UserLoginOKBody{Token: signedToken, UserInfo: &res})
}

// GetUserProfileHandler is API handler for /users/profile GET
func GetUserProfileHandler(params user.GetUserProfileParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewGetUserProfileUnauthorized()
	}

	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var u models.User
	if db.Where("id = ?", uid).First(&u).RowsAffected == 0 {
		return user.NewGetUserProfileNotFound()
	}

	res := user.GetUserProfileOKBodyResults{
		City:        u.City,
		CompanyName: u.CompanyName,
		Country:     u.Country,
		Email:       u.Email,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Phone:       u.Phone,
		Position:    u.Position,
		Priority:    u.Priority,
		Username:    u.Username,
	}

	return user.NewGetUserProfileOK().WithPayload(&user.GetUserProfileOKBody{
		Message: "success",
		Results: &res,
	})
}

// UpdateUserProfileHandler is API handler for /users/profile PUT
func UpdateUserProfileHandler(params user.UpdateUserProfileParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewUpdateUserProfileUnauthorized()
	}

	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	db := config.GetDB()
	var u models.User
	if db.Where("id = ?", uid).First(&u).RowsAffected == 0 {
		return user.NewUpdateUserProfileNotFound()
	}

	u.City = body.City
	u.CompanyName = body.CompanyName
	u.Country = body.Country
	u.Email = body.Email
	u.FirstName = body.FirstName
	u.LastName = body.LastName
	u.Phone = body.Phone
	u.Position = body.Position

	err := db.Save(&u).Error
	if err != nil {
		return user.NewUpdateUserProfileNotFound().WithPayload(&user.UpdateUserProfileNotFoundBody{
			Message: err.Error(),
		})
	}

	return user.NewUpdateUserProfileOK().WithPayload(&user.UpdateUserProfileOKBody{
		Message: "success",
	})
}

func AddUserHandler(params user.AddUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewAddUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	db := config.GetDB()
	//SITE_ADMIN/ ORG_ADMIN/ USER
	if body.Role != "SITE_ADMIN" && body.Role != "ORG_ADMIN" && body.Role != "USER" {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "User Role Invalid. Could only be SITE_ADMIN/ORG_ADMIN/USER",
		})
	}

	var orgNew models.Org;
	if db.Where("org_name = ?", body.OrgName).First(&orgNew).RowsAffected == 0 {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "Org Name Invalid.",
		})
	}
	newUser := models.User{
		Username: body.Name,
		Role:     body.Role,
		Email:    body.Email,
		PwReset:  false,
		Phone:    body.Phone,
		OrgName:    body.OrgName,
	}
	var userOld models.User;
	if db.Unscoped().Where("username = ?", body.Name).First(&userOld).RowsAffected == 1 {
		//delete user permently when new user created
		db.Unscoped().Delete(&userOld)
	}
	// SELECT * FROM users WHERE age = 20;
	err := db.Create(&newUser).Error
	if err != nil {
		return user.NewAddUserUnauthorized()
	}

	newLog := models.Log{
		UserID: uid,
		Operation: "Add User:" + body.Name,
		Time: time.Now(),
		Status: "Succeed",
	}
	if db.Create(&newLog).Error != nil {
		return user.NewModifyUserForbidden()
	}

	return user.NewAddUserOK().WithPayload(&user.AddUserOKBody{
		Message: "succeed",
	})
}

func ListUserHandler(params user.ListUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewListUserUnauthorized()
	}
	var users []*models.User
	db := config.GetDB()
	db.Find(&users)
	var response []*user.ListUserOKBodyItems0
	for _, tmpUser := range users {
		newResult := user.ListUserOKBodyItems0{
			Email: tmpUser.Email,
			ID: int64(tmpUser.ID),
			Name: tmpUser.Username,
			Phone: tmpUser.Phone,
			Role: tmpUser.Role,
			OrgName: tmpUser.OrgName,
		}

		response = append(response, &newResult)
	}
	return user.NewListUserOK().WithPayload(response)
}

func ListUserOfOrgHandler(params user.ListUserOfOrgParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewListUserUnauthorized()
	}
	var users []*models.User
	db := config.GetDB()
	db.Where("org_name = ?", params.OrgName).Find(&users)
	var response []*user.ListUserOKBodyItems0
	for _, tmpUser := range users {
		newResult := user.ListUserOKBodyItems0{
			Email: tmpUser.Email,
			ID: int64(tmpUser.ID),
			Name: tmpUser.Username,
			Phone: tmpUser.Phone,
			Role: tmpUser.Role,
			OrgName: tmpUser.OrgName,
		}

		response = append(response, &newResult)
	}
	return user.NewListUserOK().WithPayload(response)
}


func ModifyUserHandler(params user.ModifyUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewModifyUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	var pol models.User
	db := config.GetDB()
	if db.Where("id = ?", params.ID).First(&pol).RowsAffected == 0 {
		return user.NewModifyUserNotFound()
	}

	//pol.Org.OrgName = body.Org
	if body.Role != "SITE_ADMIN" && body.Role != "ORG_ADMIN" && body.Role != "USER" {
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "User Role Invalid. Could only be SITE_ADMIN/ORG_ADMIN/USER",
		})
	}
	pol.Username = body.Name
	pol.Phone = body.Phone
	pol.Email = body.Email
	pol.Role = body.Role
	err := db.Save(&pol).Error
	if err != nil {
		return user.NewModifyUserForbidden()
	}

	newLog := models.Log{
		UserID: uid,
		Operation: "Modify User:" + body.Name,
		Time: time.Now(),
		Status: "Succeed",
	}
	if db.Create(&newLog).Error != nil {
		return user.NewModifyUserForbidden()
	}

	return user.NewModifyUserOK().WithPayload(&user.ModifyUserOKBody{
		Message: "success",
	})

}

func DeleteUserHandler(params user.DeleteUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewDeleteUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var pol models.User
	if db.Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return user.NewDeleteUserNotFound()
	}

	newLog := models.Log{
		UserID: uid,
		Operation: "Delete User with Id: " + strconv.FormatInt(params.ID, 10),
		Time: time.Now(),
		Status: "Succeed",
	}
	if db.Create(&newLog).Error != nil {
		return user.NewDeleteUserForbidden()
	}

	return user.NewDeleteUserOK().WithPayload(&user.DeleteUserOKBody{
		Message: "success",
	})
}