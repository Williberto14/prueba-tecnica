package entity

type Product struct{
	Id_product	string	`json:"id_product"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Status	string	`json:"status"`
	Creation_date	string	`json:"creation_date"`
	Update_date	string	`json:"update_date"`
	Account_id	string	`json:"acount_id"`
	Format_product	string	`json:"format_product"`
	Value_unit	float64	`json:"value_unit"`
	Unit_name	string	`json:"unit_name"`
	Unit_description	string	`json:"unit_description"`
	Stock	int	`json:"stock"`
}