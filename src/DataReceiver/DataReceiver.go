package DataReceiver

import Data "awesomeProject/src"

type DataReceiver interface {
	ReceiveData(data chan []Data.MarketData)
}

