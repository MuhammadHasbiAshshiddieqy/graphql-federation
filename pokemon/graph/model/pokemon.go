package model

type Pokemon struct {
	ID        string   `json:"id"`
	TrainerID *int     `json:"trainerId"`
	Number    string   `json:"number"`
	Name      string   `json:"name"`
	Types     []*Type  `json:"types"`
	Trainer   *Trainer `json:"trainer"`
}

type Type struct {
	ID        int    `json:"id"`
	PokemonID string `json:"pokemonId"`
	TypeName  string `json:"typeName"`
}
