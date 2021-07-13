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
	var orgOld models.OrgNew;
	if db.Unscoped().Where("org_name = ?", body.Name).First(&orgOld).RowsAffected == 1 {
		//delete user permently when new user created
		db.Unscoped().Delete(&orgOld)
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
		// get resource id  from resource new table
		var resources []*models.ResourceNew
		db.Where("org_id = ? ", tmpOrg.ID).Find(&resources)
		var totalCpu, totalRAM, totalDisk, curCpu, curRAM, curDisk float64
		for _, tmpRes :=  range resources{
			var resUsages models.ResourceUsage
			db.Where("id = ? ", tmpRes.ID).Find(&resUsages)
			curCpu += resUsages.CurrentCPU
			curRAM += resUsages.CurrentRAM
			curDisk += resUsages.CurrentDisk
			totalCpu += resUsages.TotalCPU
			totalRAM += resUsages.TotalRAM
			totalDisk += resUsages.TotalDisk
		}
		newResult := org.ListOrgOKBodyItems0{
			ID: int64(tmpOrg.ID),
			Name: tmpOrg.OrgName,
			CurrentCPU: curCpu,
			CurrentRAM: curRAM,
			CurrentDisk: curDisk,
			TotalCPU:  totalCpu,
			TotalRAM:  totalRAM,
			TotalDisk:  totalDisk,
		}

		reponse = append(reponse, &newResult)
	}
	return org.NewListOrgOK().WithPayload(reponse)
}

func DeleteOrgHandler(params org.DeleteOrgParams) middleware.Responder {
	db := config.GetDB()
	var pol models.OrgNew
	db.Where("id = ? ", params.ID).Find(&pol)
	var orgName = pol.OrgName
	if db.Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return org.NewDeleteOrgNotFound()
	}

	//Delete User in that org
	var users []*models.UserNew
	db.Where("org_name = ? ", orgName).Delete(&users)


	return org.NewDeleteOrgOK().WithPayload(&org.DeleteOrgOKBody{
		Message: "success",
	})
}

