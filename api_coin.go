package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTrackedCoins(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	username := r.Header.Get("username")
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		return
	}

	rows, err := db.Query("SELECT id, name_id FROM coins WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Could not get coins", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coinResp CoinResp
	for rows.Next() {
		var coin Coin
		err := rows.Scan(&coin.ID, &coin.Name)
		if err != nil {
			http.Error(w, "Could not read coins", http.StatusInternalServerError)
			return
		}
		coin.Price, _ = GetCoinPrice(coin.Name)
		coinResp.Data = append(coinResp.Data, coin)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coinResp)
}

func AddCoin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	username := r.Header.Get("username")
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		return
	}

	var coinReq CoinReq
	json.NewDecoder(r.Body).Decode(&coinReq)

	// Verify coin existence in CoinCap API
	_, err = GetCoinPrice(coinReq.NameID)
	if err != nil {
		http.Error(w, "Coin not found in CoinCap API", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO coins (name_id, user_id) VALUES (?, ?)",
		coinReq.NameID, userID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Could not add coin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Add coin success"})
}

func RemoveCoin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	username := r.Header.Get("username")
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		return
	}

	vars := mux.Vars(r)
	coinID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Path variable id can't be found"})
		return
	}

	_, err = db.Exec("DELETE FROM coins WHERE id = ? AND user_id = ?", coinID, userID)
	if err != nil {
		http.Error(w, "Could not delete coin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete coin success"})
}
