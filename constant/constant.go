package constant

const baseURL = "https://www.binance.com/bapi/futures"

const (
	ENDPOINT_GET_POSITION              = baseURL + "/v1/public/future/leaderboard/getOtherPosition"
	ENDPOINT_GET_LEADERBOARD_BASE_INFO = baseURL + "/v2/public/future/leaderboard/getOtherLeaderboardBaseInfo"
	ENDPOINT_GET_PERFORMANCE           = baseURL + "/v1/public/future/leaderboard/getOtherPerformance"
)
