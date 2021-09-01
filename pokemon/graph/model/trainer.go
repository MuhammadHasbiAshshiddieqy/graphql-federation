package model

type Trainer struct {
	ID int `json:"id"`
}

func (Trainer) IsEntity() {}
