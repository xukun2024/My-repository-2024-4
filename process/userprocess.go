package process

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

type UserProcess struct{}

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	GasPrice string `json:"gasPrice"`
	Nonce    string `json:"nonce"`
	Input    string `json:"input"`
}

func (this *UserProcess) Watch(url string) (err error) {
	// 替换为你的Infura项目ID和网络（例如：mainnet、ropsten等）
	infuraURL := url

	// 连接到Infura提供的以太坊节点的RPC接口
	client, err := rpc.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	// 调用 eth_getBlockByNumber RPC方法获取最新的区块信息
	var block map[string]interface{}
	err = client.CallContext(context.Background(), &block, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		log.Fatal(err)
	}

	// 提取区块中的交易信息
	var transactions []Transaction
	transactionsRaw := block["transactions"].([]interface{})
	for _, txRaw := range transactionsRaw {
		tx := txRaw.(map[string]interface{})

		var from, to, nonce, input string

		if txFrom, ok := tx["from"].(string); ok {
			from = txFrom
		}

		if txTo, ok := tx["to"].(string); ok {
			to = txTo
		}

		if txNonce, ok := tx["nonce"].(string); ok {
			nonce = txNonce
		}

		if txInput, ok := tx["input"].(string); ok {
			input = txInput
		}

		// 转换 value 和 gas price 为十进制
		valueHex := tx["value"].(string)
		valueDec := hexToDecimal(valueHex)

		gasPriceHex := tx["gasPrice"].(string)
		gasPriceDec := hexToDecimal(gasPriceHex)

		transactions = append(transactions, Transaction{
			From:     from,
			To:       to,
			Value:    valueDec,
			GasPrice: gasPriceDec,
			Nonce:    nonce,
			Input:    input,
		})
	}

	// 打印交易信息
	for _, tx := range transactions {
		fmt.Println("From:", tx.From)
		fmt.Println("To:", tx.To)
		fmt.Println("Value:", tx.Value)
		fmt.Println("Gas Price:", tx.GasPrice)
		fmt.Println("Nonce:", tx.Nonce)
		//fmt.Println("Data:", tx.Input)
		fmt.Println("---------------")
	}
	return
}

// 将十六进制转换为十进制
func hexToDecimal(hexValue string) string {
	decimalValue, _ := new(big.Int).SetString(hexValue[2:], 16)
	return decimalValue.String()
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum/rpc"
// )

// type Transaction struct {
// 	From     string `json:"from"`
// 	To       string `json:"to"`
// 	Value    string `json:"value"`
// 	GasPrice string `json:"gasPrice"`
// 	Nonce    string `json:"nonce"`
// 	Input    string `json:"input"`
// }

// func main() {
// 	// 替换为你的Infura项目ID和网络（例如：mainnet、ropsten等）
// 	infuraURL := "https://sepolia.infura.io/v3/e395230f1edd4c44a600972cd2096230"

// 	// 连接到Infura提供的以太坊节点的RPC接口
// 	client, err := rpc.Dial(infuraURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 调用 eth_getBlockByNumber RPC方法获取最新的区块信息
// 	var block map[string]interface{}
// 	err = client.CallContext(context.Background(), &block, "eth_getBlockByNumber", "latest", true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 提取区块中的交易信息
// 	var transactions []Transaction
// 	transactionsRaw := block["transactions"].([]interface{})
// 	for _, txRaw := range transactionsRaw {
// 		tx := txRaw.(map[string]interface{})

// 		var from, to, value, gasPrice, nonce, input string

// 		if txFrom, ok := tx["from"].(string); ok {
// 			from = txFrom
// 		}

// 		if txTo, ok := tx["to"].(string); ok {
// 			to = txTo
// 		}

// 		if txValue, ok := tx["value"].(string); ok {
// 			value = txValue
// 		}

// 		if txGasPrice, ok := tx["gasPrice"].(string); ok {
// 			gasPrice = txGasPrice
// 		}

// 		if txNonce, ok := tx["nonce"].(string); ok {
// 			nonce = txNonce
// 		}

// 		if txInput, ok := tx["input"].(string); ok {
// 			input = txInput
// 		}

// 		transactions = append(transactions, Transaction{
// 			From:     from,
// 			To:       to,
// 			Value:    value,
// 			GasPrice: gasPrice,
// 			Nonce:    nonce,
// 			Input:    input,
// 		})
// 	}

// 	// 打印交易信息
// 	for _, tx := range transactions {
// 		fmt.Println("From:", tx.From)
// 		fmt.Println("To:", tx.To)
// 		fmt.Println("Value:", tx.Value)
// 		fmt.Println("Gas Price:", tx.GasPrice)
// 		fmt.Println("Nonce:", tx.Nonce)
// 		//fmt.Println("Data:", tx.Input)
// 		fmt.Println("---------------")
// 	}
// }
