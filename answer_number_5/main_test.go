package main_test

import (
	answernumber5 "answernumber5"
	"testing"
)

func TestGetDataFromApi(t *testing.T) {
	status := answernumber5.GetDataFromApi([]string{"https://freetestapi.com/api/v1/currencies?limit=100", "https://freetestapi.com/api/v1/countries?limit=100"})
	if status != "200 OK" {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", "200 OK", status)
	}
}

func TestGetDataFromApiWithoutAwait(t *testing.T) {
	status := answernumber5.GetDataFromApiWithoutAwait([]string{"https://freetestapi.com/api/v1/currencies?limit=100", "https://freetestapi.com/api/v1/countries?limit=100"})
	if status != "200 OK" {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", "200 OK", status)
	}

}
