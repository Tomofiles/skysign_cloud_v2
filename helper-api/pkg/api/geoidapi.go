package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang/glog"
)

const (
	geoidURL = "https://vldb.gsi.go.jp/sokuchi/surveycalc/geoid/calcgh/cgi/geoidcalc.pl"
)

// GetGeoidHeight .
func GetGeoidHeight(latitude, longitude float64) float64 {
	req, err := http.NewRequest("GET", geoidURL, nil)
	if err != nil {
		glog.Fatal(err)
	}

	params := req.URL.Query()
	params.Add("outputType", "json")
	params.Add("latitude", strconv.FormatFloat(latitude, 'f', 2, 64))
	params.Add("longitude", strconv.FormatFloat(longitude, 'f', 2, 64))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		glog.Fatal(err)
	}
	defer resp.Body.Close()

	byteResp, _ := ioutil.ReadAll(resp.Body)

	var geoidHeight geoidHeight
	if err := json.Unmarshal(byteResp, &geoidHeight); err != nil {
		glog.Fatal(err)
	}

	f, _ := strconv.ParseFloat(geoidHeight.OutputData.GeoidHeight, 64)

	return f
}

type geoidHeight struct {
	OutputData struct {
		GeoidHeight string `json:"geoidHeight"`
	} `json:"OutputData"`
}
