package HttpProviders

import (
	Data "awesomeProject/src"
	Base "awesomeProject/src/DataProvider"
	"encoding/json"
	"fmt"

)

type infoStruct struct {
	AdjHigh float32 `json:"adj_high"`
	AdjLow float32 `json:"adj_low"`
}

type marketStackData struct {
	Data.MarketData
	infoStruct

}

type jsonResponse struct{
	Data  []marketStackData `json:"data"`
}


//=============================================================================


type MarketStack struct {
	Base.HTTPDataProvider
}


func(provider MarketStack) TransformData(dataChannel chan []byte){
	dataStream := make(chan []Data.MarketData)
	go provider.TransmitDataTo(dataStream)
	for {
		data := <- dataChannel
		var jsonData jsonResponse
		err := json.Unmarshal(data, &jsonData)

		if err != nil {
			fmt.Println("transformData[x] ERROR: unmarshalling: ",err)
		}

		var marketDataArr []Data.MarketData

		// we get the extra parameters in the type marketStackData and put then as info inside MarketData
		for _,data := range jsonData.Data{
			info := infoStruct{AdjHigh: data.AdjHigh,AdjLow: data.AdjLow}
			infoJSON,_ := json.Marshal(info)
			data.MarketData.Info = infoJSON
			marketDataArr = append(marketDataArr,data.MarketData)
		}

		dataStream <- marketDataArr

	}

}

func NewMarketStack(apiKey string,symbols  []string,time int) *MarketStack{
	url := "http://api.marketstack.com/v1/eod?access_key=" + apiKey + "&symbols="
	marketStack := &MarketStack{}
	marketStack.Url = url
	marketStack.Symbols = symbols
	marketStack.WaitingSeconds = time
	marketStack.IDataProvider = marketStack
	// marketStack.RetrieveData()
	return marketStack
}

