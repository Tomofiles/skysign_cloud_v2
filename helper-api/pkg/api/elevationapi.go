package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang/glog"
)

const (
	elevationURL = "https://cyberjapandata2.gsi.go.jp/general/dem/scripts/getelevation.php"
)

// GetElevation .
func GetElevation(latitude, longitude float64) float64 {
	req, err := http.NewRequest("GET", elevationURL, nil)
	if err != nil {
		glog.Fatal(err)
	}

	params := req.URL.Query()
	params.Add("outtype", "JSON")
	params.Add("lat", strconv.FormatFloat(latitude, 'f', 2, 64))
	params.Add("lon", strconv.FormatFloat(longitude, 'f', 2, 64))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		glog.Fatal(err)
	}
	defer resp.Body.Close()

	byteResp, _ := ioutil.ReadAll(resp.Body)

	var elevation elevation
	var f float64
	if err := json.Unmarshal(byteResp, &elevation); err == nil {
		f = elevation.Elevation
	}

	return f
}

type elevation struct {
	Elevation float64 `json:"elevation"`
}
