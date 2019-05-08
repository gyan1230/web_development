package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://gyan1230.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"Dor2CjFs31IkASmdACiNq3mZdu8IiI5D\",\"client_secret\":\"zxKzmiUTa_B9zx7BqKutaSlXGPvdt-9lWT9_wDKpy8jsn137opLuSWM2SLoIzOdH\",\"audience\":\"https://gyan1230.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
