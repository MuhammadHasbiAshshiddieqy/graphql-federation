package model

type Trainer struct {
	ID     int    `json:"id"`
	TeamID *int   `json:"teamId"`
	Name   string `json:"name"`
	Team   *Team  `json:"team"`
}

func (Trainer) IsEntity() {}
