package main

import (
	"fmt"
	"reflect"
	"unsafe"
	"context"
	"time"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	var rpcEndpoint = "ws://127.0.0.1:9650/ext/bc/C/ws"
	wsClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		fmt.Println(err)
	}
	var clientValue reflect.Value
	clientValue = reflect.ValueOf(wsClient).Elem()
	fieldStruct := clientValue.FieldByName("c")
	clientPointer := reflect.NewAt(fieldStruct.Type(), unsafe.Pointer(fieldStruct.UnsafeAddr())).Elem()
	finalClient, _ := clientPointer.Interface().(*rpc.Client)
	newTxsChannel := make(chan common.Hash)
	finalClient.EthSubscribe(
		context.Background(), newTxsChannel, "newPendingTransactions",
	)
	fmt.Println("hearing mempool txs")
	for {
		select {
		case transaction := <-newTxsChannel:
			tx, is_pending, _ := wsClient.TransactionByHash(context.Background(), transaction)
			fmt.Println(time.Now().UnixNano(), tx.Hash().Hex(), is_pending)
		}
	}
}
