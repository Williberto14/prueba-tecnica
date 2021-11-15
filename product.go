package main

import (
	"time"

	"gorm.io/datatypes"
)

type Product struct {
	Product_id			string			`gorm:"type:varchar(36);primaryKey" json:"product_id"`
	Name  				string 			`gorm:"type:varchar(45)" json:"name"`
	Description 		string 			`gorm:"type:text(300)" json:"description"`
	Status       		string 			`gorm:"type:varchar(45)" json:"status"`
	Creation_date 		time.Time 		`json:"creation_date"`
	Update_date    		time.Time 		`json:"update_date"`
	Account_id     		string 			`gorm:"type:varchar(25)" json:"account_id"`
	Format_product 		datatypes.JSON 	`json:"format_product"`
	Value_unit      	float64 		`gorm:"type:decimal(15,2)" json:"value_unit"`
	Unit_name       	string 			`gorm:"type:varchar(45)" json:"unit_name"`
	Unit_description	string	 		`gorm:"type:text(300)" json:"unit_description"`
	Stock     			int 			`json:"stock"`
}