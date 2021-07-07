package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/org"
	"tides-server/pkg/restapi/operations/user"
)


func AddUserHandler(params user.AddUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()

	//TODO retrive Org by name, user id generate
	newUser := models.UserNew{
		Username: body.Name,
		Role: body.Role,
		Email: body.Email,
		PwReset: false,
		Phone: body.Phone,
	}

	err := db.Create(&newUser).Error
	if err != nil {
		return user.NewAddUserUnauthorized()
	}

	return user.NewAddUserOK().WithPayload(&user.AddUserOKBody{
		Message: "succeed",
	})
}

func ListUserHandler(params user.ListUserParams) middleware.Responder {
	var users []*models.UserNew
	db := config.GetDB()
	db.Find(&users)
	var reponse []*user.ListUserOKBodyItems0
	for _, tmpUser := range users {
		newResult := user.ListUserOKBodyItems0{
			Email: tmpUser.Email,
			ID: int64(tmpUser.ID),
			Name: tmpUser.Username,
			Org: tmpUser.Org.OrgName,
			Phone: tmpUser.Phone,
			Role: tmpUser.Role,
		}

		reponse = append(reponse, &newResult)
	}
	return user.NewListUserOK().WithPayload(reponse)
}


func ModifyUserHandler(params user.ModifyUserParams) middleware.Responder {
	//db := config.GetDB()
	//var pol models.UserNew
	userId, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	var pol models.UserNew
	db := config.GetDB()
	if db.Where("user_id = ?", userId).First(&pol).RowsAffected == 0 {
		return user.NewModifyUserNotFound()
	}

	//pol.Org.OrgName = body.Org //TODO modify org name, need to change through
	pol.Phone = body.Phone
	pol.Email = body.Email
	pol.Role = body.Role

	err := db.Save(&pol).Error
	if err != nil {
		return user.NewModifyUserForbidden()
	}

	return user.NewModifyUserOK().WithPayload(&user.ModifyUserOKBody{
		Message: "success",
	})

}

func DeleteUserHandler(params user.DeleteUserParams) middleware.Responder {
	db := config.GetDB()
	var pol models.UserNew
	if db.Unscoped().Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return org.NewDeleteOrgNotFound()
	}

	return user.NewDeleteUserOK().WithPayload(&user.DeleteUserOKBody{
		Message: "success",
	})
}

