package main

import (
	"context"
	TestFilter "deploy_contract/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strconv"
)

var URL_ = ""
var private_ = ""

func main() {
	r := gin.Default()
	r.POST("/setParams", setParams)
	r.GET("/createContract", createContract)
	r.GET("/invoke", invokeFunc)
	r.GET("/getTopics", getTopics)
	r.GET("/ongTransfer", ongTransfer)
	err := r.Run()
	if err != nil {
		return
	}
}

func setParams(c *gin.Context) {
	URL_ = c.PostForm("URL")
	private_ = c.PostForm("private")
	fmt.Println(URL_)
	fmt.Println(private_)
	c.JSON(200, gin.H{
		"result": "success",
	})
}

func createContract(c *gin.Context) {
	fmt.Println(URL_)
	fmt.Println(private_)
	contractAddr := TestFilter.DeployFilterContract(URL_, private_)
	c.JSON(200, gin.H{
		"contractAddr": contractAddr,
	})
}

func invokeFunc(c *gin.Context) {
	contractAddr := c.Query("addr")
	funcName := c.Query("funcName")
	result := "success"
	txHash := ""
	testFilter := getTestFilter(contractAddr)
	opts := getTransaction()
	var tx *types.Transaction
	var err error
	if funcName == "getAge" {
		tx, err = testFilter.GetAge(opts)
		_checkErr(err)
	}
	if funcName == "getName" {
		tx, err = testFilter.GetName(opts)
		_checkErr(err)
	}
	if funcName == "setAge" {
		_age := c.Query("age")
		age, err1 := strconv.ParseInt(_age, 10, 64)
		_checkErr(err1)
		tx, err = testFilter.SetAge(opts, big.NewInt(age))
		_checkErr(err)
	}
	if funcName == "setName" {
		name := c.Query("name")
		tx, err = testFilter.SetName(opts, name)
		_checkErr(err)
	}

	rec, err := TestFilter.GetTransferInfoByHash(URL_, tx.Hash())
	_checkErr(err)
	if rec.Status != 1 {
		result = "fail,tx status = 0"
	}
	txHash = tx.Hash().String()
	c.JSON(200, gin.H{
		"result": result,
		"txHash": txHash,
	})

}

func getTopics(c *gin.Context) {
	funcName := c.Query("funcName")
	fmt.Println("funcName :" + funcName)
	if funcName == "" {
		TestGetAge := crypto.Keccak256Hash([]byte("TestGetAge(string,uint256)"))
		TestGetName := crypto.Keccak256Hash([]byte("TestGetName(string,uint256)"))
		TestSetAge := crypto.Keccak256Hash([]byte("TestSetAge(uint256,uint256,uint256)"))
		TestSetName := crypto.Keccak256Hash([]byte("TestSetName(string,string,uint256)"))
		c.JSON(200, gin.H{
			"TestGetAge":  TestGetAge.String(),
			"TestGetName": TestGetName.String(),
			"TestSetAge":  TestSetAge.String(),
			"TestSetName": TestSetName.String(),
			"oldName":     crypto.Keccak256Hash([]byte("")),
		})
	} else {
		_funcName := crypto.Keccak256Hash([]byte(funcName))
		c.JSON(200, gin.H{
			funcName: _funcName.String(),
		})
	}
}

func getTransaction() *bind.TransactOpts {
	ethClient := TestFilter.GetEthClient(URL_)
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	Private, err := crypto.HexToECDSA(private_)
	_checkErr(err)
	addr := crypto.PubkeyToAddress(Private.PublicKey)
	_checkErr(err)
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	_checkErr(err)
	opts, err := bind.NewKeyedTransactorWithChainID(Private, chainId)
	_checkErr(err)
	opts.From = addr
	opts.GasPrice = big.NewInt(500000000000)
	opts.GasLimit = 3000000
	opts.Nonce = big.NewInt(int64(nonce))
	return opts
}

func getTestFilter(contractAddr string) *TestFilter.TestFilter {
	ethClient := TestFilter.GetEthClient(URL_)
	addr := common.HexToAddress(contractAddr)
	testFilter, err := TestFilter.NewTestFilter(addr, ethClient)
	_checkErr(err)
	return testFilter
}

func ongTransfer(c *gin.Context) {
	toAddr := c.Query("addr")
	// 构建一个私链对象
	ethClient := TestFilter.GetEthClient(URL_)
	//查chainIdz
	chainId, err := ethClient.ChainID(context.Background())
	_checkErr(err)
	testPrivateKey, err := crypto.HexToECDSA(private_)
	_checkErr(err)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	_checkErr(err)
	log.Printf("nonce: %d", nonce)
	_to := common.HexToAddress(toAddr)
	gasPrice := big.NewInt(500000000000)
	rawTx := types.NewTransaction(nonce, _to, big.NewInt(1), 250000, gasPrice, nil)
	signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	key, err := crypto.HexToECDSA(private_)
	_checkErr(err)
	//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	_checkErr(err)
	//fmt.Println("52ln rawTx: ", sigTransaction)
	err = ethClient.SendTransaction(context.Background(), sigTransaction)
	_checkErr(err)
	log.Printf("tx hash(evm): %s", sigTransaction.Hash())
	rec, err := TestFilter.GetTransferInfoByHash(URL_, sigTransaction.Hash())
	_checkErr(err)
	result := ""
	if rec.Status != 1 {
		result = "fail,tx status = 0"
	}
	result = sigTransaction.Hash().String()
	c.JSON(200, gin.H{
		"result": result,
	})
}

func _checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
