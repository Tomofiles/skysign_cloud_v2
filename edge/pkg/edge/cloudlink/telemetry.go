package cloudlink

import (
	"edge/pkg/edge"
	"edge/pkg/edge/domain/telemetry"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PushTelemetry .
func PushTelemetry(cloud string, telemetry telemetry.Telemetry) (string, *edge.CommandIDs, error) {
	telem := telemetry.Get()
	if telem.State.FlightMode == "" {
		log.Println("no telemetry prepared.")
		return telem.ID, &edge.CommandIDs{CommandIds: make([]string, 0)}, nil
	}

	jsonData, _ := json.Marshal(telem)
	log.Printf("Send CLOUD data=%s\n", jsonData)

	req, err := http.NewRequest(
		"POST",
		cloud+"/api/v1/communications/"+telem.ID+"/telemetry",
		strings.NewReader(string(jsonData)),
	)
	if err != nil {
		log.Println("cloud telemetry request error:", err)
		return "", nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("cloud telemetry request error:", err)
		return "", nil, err
	}

	requestBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("cloud telemetry response error:", err)
		return "", nil, err
	}

	var commandIDs edge.CommandIDs
	err = json.Unmarshal(requestBody, &commandIDs)
	if err != nil {
		log.Println("cloud telemetry response error:", err)
		return "", nil, err
	}

	log.Printf("Receive CLOUD data=%s\n", requestBody)

	return telem.ID, &commandIDs, nil
}
