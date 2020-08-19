package Data

import "fmt"

type MarketData struct {
	Open     float32 `json:"open"`
	Close    float32 `json:"close"`
	High     float32 `json:"high"`
	Low      float32 `json:"low"`
	Symbol   string  `json:"symbol"`
	Exchange string  `json:"exchange"`
	Date     string  `json:"date"`
	Info     []byte
}

func (d MarketData) String() string {
	return fmt.Sprintf("%v:\n_______\nopen:%v\nclose:%v\nhigh:%v\nlow:%v\ndate:%v",
		d.Exchange, d.Open, d.Close, d.High, d.Low, d.Date)
}
