package utils

import (
	"encoding/base64"

	"github.com/labstack/gommon/log"
)

type (
	ErrorResponse struct {
		ErrorMessage string `json:"error_msg"`
	}
)

func MappingErrorResponse(errObj error) (errResp ErrorResponse) {
	if errObj != nil {
		errResp.ErrorMessage = errObj.Error()
		return
	}
	return
}

func Encodebase64ToString(data []byte) string {
	encodedText := base64.StdEncoding.EncodeToString(data)
	return encodedText
}

func DecodeBase64(text string) (decodedTextByte []byte, err error) {
	decodedTextByte, err = base64.StdEncoding.DecodeString(text)
	if err != nil {
		return
	}
	return
}

func LogQueryDB(query string) {
	log.Infof("[DB] QUERY: %+v", query)
}

// func HTTPRequestToMovies(url string) (res ResMovies, err error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("error http request: ", err.Error())
// 		return
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("error read body request: ", err.Error())
// 		return
// 	}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		fmt.Println("error unmarshal body request: ", err.Error())
// 		return
// 	}
// 	return
// }
