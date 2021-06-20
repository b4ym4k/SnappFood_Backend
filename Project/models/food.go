package models

type Food struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
	Ready bool   `json:"isReady,omitempty"`
	Rate  int    `json:"rate,omitempty"`
}
