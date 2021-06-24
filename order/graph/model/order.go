package model

type OrderAddress struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Address1    *string `gorm:"column:address_1" json:"address1"`
	Address2    *string `gorm:"column:address_2" json:"address2"`
	SubDistrict *string `json:"subDistrict"`
	District    *string `json:"district"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
	PostalCode  *string `json:"postalCode"`
	Coordinate  *string `json:"coordinate"`
}
