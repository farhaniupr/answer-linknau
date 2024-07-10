package main

import (
	"fmt"
	"sync"
	"time"

	"gitee.com/go-package/carbon"
	"github.com/go-resty/resty/v2"
)

// async await
func GetDataFromApi(url []string) string {
	var waitAsync sync.WaitGroup

	start := time.Now()

	errorChannel := make(chan error, len(url))
	defer close(errorChannel)
	for _, valueUrl := range url {
		waitAsync.Add(1)
		go func(url string) {
			defer waitAsync.Done()

			client := resty.New()
			resp, err := client.R().
				EnableTrace().
				Get(valueUrl)
			if err != nil {
				errorChannel <- fmt.Errorf("error %s: %v", url, err)
				return
			}

			fmt.Println(resp.Status(), resp.Time(), resp.ReceivedAt(), carbon.Parse(start.Format("2006-01-02 15:04:05")).DiffForHumans(), resp)
		}(valueUrl)
	}

	waitAsync.Wait()

	if len(errorChannel) > 0 {
		return "500 ERROR"
	}
	return "200 OK"
}

// async
func GetDataFromApiWithoutAwait(url []string) string {
	start := time.Now()
	for _, valueUrl := range url {
		go func() string {
			client := resty.New()
			resp, err := client.R().
				EnableTrace().
				Get(valueUrl)
			if err != nil {
				return "500 ERROR"
			}

			fmt.Println(resp.Status(), resp.Time(), resp.ReceivedAt(), carbon.Parse(start.Format("2006-01-02 15:04:05")).DiffForHumans(), resp)
			return "200 OK"
		}()
	}

	return "200 OK"
}

func main() {

	// fmt.Println(GetDataFromApiWithoutAwait([]string{"https://freetestapi.com/api/v1/currencies?limit=200", "https://freetestapi.com/api/v1/countries?limit=100"}))

	fmt.Println(GetDataFromApi([]string{"https://freetestapi.com/api/v1/currencies?limit=200", "https://freetestapi.com/api/v1/countries?limit=100"}))

}
