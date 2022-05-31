package TestFilter

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

func DeployFilterContract(URL string, private string) (string,string) {

	testContractAbi, err := ReadAll("contract/TestFilter_sol_TestFilter.abi")
	_checkErr(err)
	testContractBin, err := ReadAll("contract/TestFilter_sol_TestFilter.bin")
	_checkErr(err)

	ethClient := GetEthClient(URL)
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	testPrivateKey, err := crypto.HexToECDSA(private)
	_checkErr(err)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	fmt.Println(nonce)
	_checkErr(err)
	opts, err := bind.NewKeyedTransactorWithChainID(testPrivateKey, chainId)
	_checkErr(err)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(500000000000)
	opts.GasLimit = 3000000
	binDecode, _ := hex.DecodeString(string(testContractBin))
	rawTx, err := _newDeployEvmContract(opts, binDecode, string(testContractAbi), "Test_filter", big.NewInt(18))
	_checkErr(err)
	err = ethClient.SendTransaction(context.Background(), rawTx)
	_checkErr(err)
	txReceipt, err := GetTransferInfoByHash(URL, rawTx.Hash())
	_checkErr(err)
	if txReceipt.Status != 1 {
		panic("tx error!")
	}
	fmt.Println(txReceipt.ContractAddress.String())
	return txReceipt.ContractAddress.String(),rawTx.Hash().String()

}
func _newDeployEvmContract(opts *bind.TransactOpts, code []byte, jsonABI string, params ...interface{}) (*types.Transaction, error) {
	fmt.Println("params :", params)
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	_checkErr(err)
	input, err := parsed.Pack("", params...)
	fmt.Println("input :", input)
	_checkErr(err)
	input = append(code, input...)
	log.Print("input: ")
	log.Println(input)
	deployTx := types.NewContractCreation(opts.Nonce.Uint64(), opts.Value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, deployTx)
	_checkErr(err)
	return signedTx, err
}

func _checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetTransferInfoByHash(URL string, hash common.Hash) (*types.Receipt, error) {
	var resErr error
	var res *types.Receipt
	for i := 0; i < 10; i++ {
		time.Sleep(3000000000) //设置10次请求，每次间隔3s 30s不落块超时 抛异常
		client_ := GetEthClient(URL)
		receipt, err := client_.TransactionReceipt(context.Background(), hash)
		if err == nil {
			resErr = nil
			res = receipt
			break
		} else {
			resErr = err
		}
	}
	return res, resErr
}

func GetEthClient(URL string) *ethclient.Client {
	client, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}
	return client
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
