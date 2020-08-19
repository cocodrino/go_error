package DataReceivers

import (
	Data "awesomeProject/src"
	"awesomeProject/src/DataReceiver"
	"fmt"
)

type TestStore struct {
	DataReceiver.DataReceiver
}


func (store *TestStore) StoreData(dataChan chan []Data.MarketData) {
	for{
		dataArray :=<- dataChan
		for _,data := range dataArray{
			fmt.Println(data)
		}
	}
}

func NewTestStore() TestStore{
	return TestStore{}
}

