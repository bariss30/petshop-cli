package domain

type Animal struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	OwnerID  int     `json:"ownerid"`
	Nickname string  `json:"nickname"`
}
