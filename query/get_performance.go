package query

import (
	"bispy-agent/constant"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetPerformanceRequest struct {
	EncryptedUid string `json:"encryptedUid"`
	TradeType    string `json:"tradeType"`
}

type GetPerformanceResponse struct {
	Code          string            `json:"code"`
	Message       interface{}       `json:"message"`
	MessageDetail interface{}       `json:"messageDetail"`
	Data          []PerformanceData `json:"data"`
	Success       bool              `json:"success"`
}

type PerformanceData struct {
	PeriodType     string  `json:"periodType"`
	StatisticsType string  `json:"statisticsType"`
	Value          float64 `json:"value"`
	Rank           int     `json:"rank"`
}

func GetPerformance(request *GetPerformanceRequest) GetPerformanceResponse {

	requestToJson, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		log.Fatal("[GetPerformance]: Error while marshalling request to json", marshalErr)
	}
	body := bytes.NewBuffer(requestToJson)

	response, postErr := http.Post(constant.ENDPOINT_GET_PERFORMANCE, "application/json", body)
	if postErr != nil {
		log.Fatal("[GetPerformance]: Error while posting request", postErr)
	}

	defer response.Body.Close()

	bytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("[GetPerformance]: Error while reading response body", readErr)
	}

	var responseObject GetPerformanceResponse

	unmarshalError := json.Unmarshal(bytes, &responseObject)
	if unmarshalError != nil {
		log.Fatal("[GetPerformance]: Error while unmarshalling response body", unmarshalError)
	}

	return responseObject
}
