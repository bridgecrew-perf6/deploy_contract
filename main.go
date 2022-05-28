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

func _checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
