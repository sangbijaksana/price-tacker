package main

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_conf"`
}

type CoinReq struct {
	NameID string `json:"name_id"`
}

type Coin struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CoinResp struct {
	Data []Coin `json:"data"`
}

type CoinInfo struct {
	NameID   string `json:"id"`
	PriceUsd string `json:"priceUsd"`
}
