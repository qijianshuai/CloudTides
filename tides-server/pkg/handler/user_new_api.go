package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/user"
)


func AddUserHandler(params user.AddUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()
	//SITE_ADMIN/ ORG_ADMIN/ USER
	if body.Role != "SITE_ADMIN" && body.Role != "ORG_ADMIN" && body.Role != "USER" {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "User Role Invalid. Could only be SITE_ADMIN/ORG_ADMIN/USER",
		})
	}

	var orgNew models.OrgNew;
	if db.Where("org_name = ?", body.OrgName).First(&orgNew).RowsAffected == 0 {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "Org Name Invalid.",
		})
	}
	newUser := models.UserNew{
		Username: body.Name,
		Role:     body.Role,
		Email:    body.Email,
		PwReset:  false,
		Phone:    body.Phone,
		OrgName:    body.OrgName,
	}
	var userOld models.UserNew;
	if db.Unscoped().Where("username = ?", body.Name).First(&userOld).RowsAffected == 1 {
		//delete user permently when new user created
		db.Unscoped().Delete(&userOld)
	}
	// SELECT * FROM users WHERE age = 20;
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
	//db := config.GetDB()
	//var pol models.UserNew
	body := params.ReqBody
	var pol models.UserNew
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
	if db.Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return user.NewDeleteUserNotFound()
	}

	return user.NewDeleteUserOK().WithPayload(&user.DeleteUserOKBody{
		Message: "success",
	})
}

