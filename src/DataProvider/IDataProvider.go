package DataProvider

import (
	Data "awesomeProject/src"
)

type IDataProvider interface {
	RetrieveData()
	TransformData(data chan []byte)
	TransmitDataTo(data chan []Data.MarketData)
}



