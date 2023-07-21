package query

import (
	"bispy-agent/constant"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetPositionRequest struct {
	EncryptedUid string `json:"encryptedUid"`
	TradeType    string `json:"tradeType"`
}

type GetPositionResponse struct {
	Code    string       `json:"code"`
	Data    PositionData `json:"data"`
	Success bool         `json:"success"`
}

type PositionData struct {
	OtherPositionRetList []struct {
		Symbol          string  `json:"symbol"`
		EntryPrice      float64 `json:"entryPrice"`
		MarkPrice       float64 `json:"markPrice"`
		Pnl             float64 `json:"pnl"`
		Roe             float64 `json:"roe"`
		UpdateTime      []int   `json:"updateTime"`
		Amount          float64 `json:"amount"`
		UpdateTimeStamp int64   `json:"updateTimeStamp"`
		Yellow          bool    `json:"yellow"`
		TradeBefore     bool    `json:"tradeBefore"`
	} `json:"otherPositionRetList"`
	UpdateTime      []int `json:"updateTime"`
	UpdateTimeStamp int64 `json:"updateTimeStamp"`
}

func GetPosition(request *GetPositionRequest) GetPositionResponse {

	requestToJson, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		log.Fatal("[GetPosition]: Error while marshalling request to json", marshalErr)
	}
	body := bytes.NewBuffer(requestToJson)

	response, postErr := http.Post(constant.ENDPOINT_GET_POSITION, "application/json", body)
	if postErr != nil {
		log.Fatal("[GetPosition]: Error while posting request", postErr)
	}

	defer response.Body.Close()

	bytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("[GetPosition]: Error while reading response body", readErr)
	}

	var responseObject GetPositionResponse

	unmarshalError := json.Unmarshal(bytes, &responseObject)
	if unmarshalError != nil {
		log.Fatal("[GetPosition]: Error while unmarshalling response body", unmarshalError)
	}

	return responseObject
}
