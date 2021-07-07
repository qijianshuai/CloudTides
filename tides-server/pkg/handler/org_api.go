package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/org"
)


func AddOrgHandler(params org.AddOrgParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()

	newOrg := models.OrgNew{
		OrgName: body.Name,
	}

	err := db.Create(&newOrg).Error
	if err != nil {
		return org.NewAddOrgUnauthorized()
	}

	return org.NewAddOrgOK().WithPayload(&org.AddOrgOKBody{
		Message: "succeed",
	})
}

func ListOrgHandler(params org.ListOrgParams) middleware.Responder {
	var orgs []*models.OrgNew
	db := config.GetDB()
	db.Find(&orgs)
	var reponse []*org.ListOrgOKBodyItems0
	for _, tmpOrg := range orgs {
		newResult := org.ListOrgOKBodyItems0{
			ID: int64(tmpOrg.ID),
			Name: tmpOrg.OrgName,
			//TODO: CPU info, etc not got yet.
		}

		reponse = append(reponse, &newResult)
	}
	return org.NewListOrgOK().WithPayload(reponse)
}

func DeleteOrgHandler(params org.DeleteOrgParams) middleware.Responder {
	db := config.GetDB()
	var pol models.OrgNew
	if db.Unscoped().Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return org.NewDeleteOrgNotFound()
	}

	return org.NewDeleteOrgOK().WithPayload(&org.DeleteOrgOKBody{
		Message: "success",
	})
}

