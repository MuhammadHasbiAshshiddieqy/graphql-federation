package model

type BasicResponse struct {
	Message string `json:"message"`
}

type UpdatePokemonInput struct {
	ID    string    `json:"id"`
	Name  *string   `json:"name"`
	Types []*string `json:"types"`
}

type DeletePokemonInput struct {
	ID string `json:"id"`
}

type CreatePokemonInput struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type CreateResponse struct {
	Message string   `json:"message"`
	Pokemon *Pokemon `json:"pokemon"`
}
