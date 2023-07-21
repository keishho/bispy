package query

import (
	"bispy-agent/constant"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetLeaderboardInfoRequest struct {
	EncryptedUid string `json:"encryptedUid"`
	TradeType    string `json:"tradeType"`
}

type GetLeaderboardInfoResponse struct {
	Code          string                 `json:"code"`
	Message       interface{}            `json:"message"`
	MessageDetail interface{}            `json:"messageDetail"`
	Data          GetLeaderboardInfoData `json:"data"`
	Success       bool                   `json:"success"`
}

type GetLeaderboardInfoData struct {
	NickName               string      `json:"nickName"`
	UserPhotoURL           string      `json:"userPhotoUrl"`
	PositionShared         bool        `json:"positionShared"`
	DeliveryPositionShared bool        `json:"deliveryPositionShared"`
	FollowingCount         int         `json:"followingCount"`
	FollowerCount          int         `json:"followerCount"`
	TwitterURL             interface{} `json:"twitterUrl"`
	Introduction           string      `json:"introduction"`
}

func GetLeaderboardInfo(request *GetLeaderboardInfoRequest) GetLeaderboardInfoResponse {

	requestToJson, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		log.Fatal("[GetLeaderboardInfo]: Error while marshalling request to json", marshalErr)
	}
	body := bytes.NewBuffer(requestToJson)

	response, postErr := http.Post(constant.ENDPOINT_GET_LEADERBOARD_BASE_INFO, "application/json", body)
	if postErr != nil {
		log.Fatal("[GetLeaderboardInfo]: Error while posting request", postErr)
	}

	defer response.Body.Close()

	bytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("[GetLeaderboardInfo]: Error while reading response body", readErr)
	}

	var responseObject GetLeaderboardInfoResponse

	unmarshalError := json.Unmarshal(bytes, &responseObject)
	if unmarshalError != nil {
		log.Fatal("[GetLeaderboardInfo]: Error while unmarshalling response body", unmarshalError)
	}

	return responseObject
}
