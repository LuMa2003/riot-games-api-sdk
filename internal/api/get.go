package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type GetResponse interface {
	IsGetResponse()
}

func Get[T GetResponse](url string) (T, error) {
	var response T
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	req.Header.Set("X-Riot-Token", os.Getenv("RIOT_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}