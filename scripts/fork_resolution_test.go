package scripts

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

var urls = []string{
	"https://localhost:8080/chain",
	"https://localhost:8081/chain",
	"https://localhost:8082/chain",
	"https://localhost:8083/chain",
	"https://localhost:8084/chain",
}

type ChainResponse struct {
	Success bool    `json:"success"`
	Result  []Block `json:"result"`
}

type Block struct {
	Index        int           `json:"Index"`
	Data         []Transaction `json:"Data"`
	MerkleRoot   string        `json:"MerkleRoot"`
	Hash         string        `json:"Hash"`
	PreviousHash string        `json:"PreviousHash"`
	Timestamp    string        `json:"Timestamp"`
	Nonce        int           `json:"Nonce"`
}

type Transaction struct {
	Amount       float64     `json:"amount"`
	Data         interface{} `json:"data"`
	Fee          float64     `json:"fee"`
	ID           string      `json:"id"`
	MiningStatus int         `json:"miningStatus"`
	Receiver     string      `json:"receiver"`
	Sender       string      `json:"sender"`
	Size         int         `json:"size"`
}

func TestForkResolution_IT(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // DO NOT use in production
	}
	client := http.Client{
		Transport: tr,
	}
	//post transaction
	postTransaction(client, t)
	//verify nonce of last block
	nonce := 0
	for _, url := range urls {
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		var chainResp ChainResponse
		err = json.NewDecoder(resp.Body).Decode(&chainResp)
		if err != nil {
			t.Error(err)
		}
		if nonce == 0 {
			nonce = chainResp.Result[0].Nonce
		} else if nonce != chainResp.Result[0].Nonce {
			t.Error("Nonce of last block is not same across all nodes")
		}
		fmt.Println(chainResp.Result[0].Nonce)
		resp.Body.Close()
	}
	fmt.Println("All nodes have same nonce, means identical blocks")
}

func postTransaction(client http.Client, t *testing.T) {
	reqBody := `{"amount":20,"receiver":"x_123","sender":"x_000","fee":1,"address":"b68f18db-5b58-4f61-8b55-3a67535cdc0d"}`

	for _, v := range urls[:2] {
		req, err := http.NewRequest("POST", v, bytes.NewBuffer([]byte(reqBody)))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response:", string(body))
	}

}
