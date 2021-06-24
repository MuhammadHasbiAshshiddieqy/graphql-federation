// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Channel struct {
	ID       string `json:"id"`
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Initials string `json:"initials"`
	IsPos    bool   `json:"isPos"`
	Name     string `json:"name"`
}

type Order struct {
	ID        string                     `json:"id"`
	AccountID int                        `json:"accountId"`
	ChannelID int                        `json:"channelId"`
	Account   *ProfileChannelAssociation `json:"account"`
	Channel   *Channel                   `json:"channel"`
}

func (Order) IsEntity() {}

type ProfileChannelAssociation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
