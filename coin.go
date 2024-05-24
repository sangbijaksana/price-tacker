package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetCoinPrice(nameID string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.coincap.io/v2/assets/%s", nameID))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("coin not found")
	}

	var result struct {
		Data CoinInfo `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	price, err := strconv.ParseFloat(result.Data.PriceUsd, 64)
	if err != nil {
		return -1, err
	}

	return price * 14250, nil // Converting USD to Rupiah (approximate rate)
}
