package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"

	"github.com/0xj1mmy/minimal-mempool-logger/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load("config.env")
	logger := logging.InitLogger(os.Getenv("LOG_OUTPUT"))
	zap.ReplaceGlobals(logger)
	var rpcEndpoint = "ws://127.0.0.1:9650/ext/bc/C/ws"
	wsClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", err.Error()))
	}
	clientValue := reflect.ValueOf(wsClient).Elem()
	fieldStruct := clientValue.FieldByName("c")
	clientPointer := reflect.NewAt(fieldStruct.Type(), unsafe.Pointer(fieldStruct.UnsafeAddr())).Elem()
	finalClient, _ := clientPointer.Interface().(*rpc.Client)
	newTxsChannel := make(chan common.Hash)
	finalClient.EthSubscribe(
		context.Background(), newTxsChannel, "newPendingTransactions",
	)
	zap.L().Info("\t", zap.Bool("Success", true), zap.String("hearing mempool txs", ""))
	for {
		select {
		case transaction := <-newTxsChannel:
			tx, is_pending, _ := wsClient.TransactionByHash(context.Background(), transaction)
			fmt.Println(time.Now().UnixNano(), tx.Hash().Hex(), is_pending)
		}
	}
}
