package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURL         = "https://api.binance.com"
	spotEndpoint    = "/api/v3/account"
	futuresEndpoint = "https://fapi.binance.com/fapi/v2/balance"
)

type SpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type FutureBalance struct {
	Asset             string `json:"asset"`
	Balance           string `json:"balance"`
	AvailableBalance  string `json:"availableBalance"`
	MaxWithdrawAmount string `json:"maxWithdrawAmount"`
}

func main() {
	spotBalances := getSpotBalances()
	fmt.Println("Spot Balances:")
	for _, asset := range spotBalances {
		if asset.Free != "0" || asset.Locked != "0" {
			fmt.Printf("Asset: %s, Free: %s, Locked: %s\n", asset.Asset, asset.Free, asset.Locked)
		}
	}

	futureBalances := getFuturesBalances()
	fmt.Println("\nFutures Account Balance V2:")
	for _, balance := range futureBalances {
		fmt.Printf("Asset: %s, Balance: %s, Available: %s, Max Withdraw: %s\n",
			balance.Asset, balance.Balance, balance.AvailableBalance, balance.MaxWithdrawAmount)
	}
}

func getSpotBalances() []SpotBalance {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	queryString := "timestamp=" + timestamp
	signature := createHmac(queryString, secretKey)
	url := baseURL + spotEndpoint + "?" + queryString + "&signature=" + signature

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}
	req.Header.Add("X-MBX-APIKEY", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %s", err)
	}

	var result struct {
		Balances []SpotBalance `json:"balances"`
	}
	json.Unmarshal(body, &result)

	var nonZeroBalances []SpotBalance
	for _, balance := range result.Balances {
		free, _ := strconv.ParseFloat(balance.Free, 64)
		locked, _ := strconv.ParseFloat(balance.Locked, 64)
		if free != 0 || locked != 0 {
			nonZeroBalances = append(nonZeroBalances, balance)
		}
	}
	return nonZeroBalances
}

func getFuturesBalances() []FutureBalance {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	queryString := "timestamp=" + timestamp
	signature := createHmac(queryString, secretKey)
	url := futuresEndpoint + "?" + queryString + "&signature=" + signature

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}
	req.Header.Add("X-MBX-APIKEY", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %s", err)
	}

	var balances []FutureBalance
	json.Unmarshal(body, &balances)

	var nonZeroBalances []FutureBalance
	for _, balance := range balances {
		if balance.Balance != "0.00000000" {
			nonZeroBalances = append(nonZeroBalances, balance)
		}
	}
	return nonZeroBalances
}

func createHmac(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
