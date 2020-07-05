package cloudlink

import (
	"edge/pkg/edge/telemetry"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// PushTelemetry .
func PushTelemetry(cloud string, telemetry telemetry.Telemetry) error {
	telem := telemetry.Get()
	if telem.FlightMode == "" {
		log.Println("no telemetry prepared.")
		return nil
	}

	jsonData, _ := json.Marshal(telem)
	log.Printf("Send CLOUD data=%s\n", jsonData)

	req, err := http.NewRequest(
		"POST",
		"http://"+cloud+"/api/v1/communications/"+telem.ID+"/telemetries",
		strings.NewReader(string(jsonData)),
	)
	if err != nil {
		log.Println("telemetry request error:", err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("telemetry request error:", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}
