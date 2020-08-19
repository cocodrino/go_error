package DataProvider

import (
	Data "awesomeProject/src"
	"awesomeProject/src/DataReceiver"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type HTTPDataProvider struct {
	IDataProvider  IDataProvider
	Url            string
	Symbols        []string
	WaitingSeconds int
	Receiver       DataReceiver.DataReceiver
}

func (provider HTTPDataProvider) RetrieveData() {
	fmt.Println("RETRIEVING DATA")

	dataStream := make(chan []byte)
	go provider.IDataProvider.TransformData(dataStream)

	client := resty.New()

	for {
		for _, symbol := range provider.Symbols {
			url := fmt.Sprint(provider.Url, symbol)

			resp, err := client.R().EnableTrace().Get(url)
			if err != nil {
				fmt.Println("retrieveData[x] ERROR: get Url: ", err)
			}
			dataStream <- resp.Body()

		}
		time.Sleep(time.Duration(provider.WaitingSeconds))
	}
}

func (provider HTTPDataProvider) TransmitDataTo(data chan []Data.MarketData) {
		go provider.Receiver.ReceiveData(data)

}

func (provider *HTTPDataProvider) ConnectToReceiver(receiver DataReceiver.DataReceiver) {
	provider.Receiver = receiver
}

func NewHTTPDataProvider(url string, symbols []string, seconds int) HTTPDataProvider {
	provider := HTTPDataProvider{
		Url:            url,
		Symbols:        symbols,
		WaitingSeconds: seconds,
	}

	// Iprovider.RetrieveData()
	return provider
}
