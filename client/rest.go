package client

import (
	"encoding/json"
	"example/hoangdz/gin/utils"
	"fmt"
)

type SquareResponse struct {
	Result int64 `json:"result"`
}

var (
	rest_endpoint = "http://127.0.0.1:3000"
)

func RestGetSquare(number int32) (SquareResponse, error) {
	url := fmt.Sprintf("%s/api/square/%d", "http://127.0.0.1:3000", number)

	resBytes, err := utils.Get(url)
	if err != nil {
		panic(err)
	}
	var result SquareResponse
	if err := json.Unmarshal(resBytes, &result); err != nil {
		panic(err)
	}
	return result, nil
}
