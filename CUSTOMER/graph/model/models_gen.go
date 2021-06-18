// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Customer struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	ProfileID   int    `json:"profileId"`
}

type Order struct {
	ID       string    `json:"id"`
	Customer *Customer `json:"customer"`
}

func (Order) IsEntity() {}
