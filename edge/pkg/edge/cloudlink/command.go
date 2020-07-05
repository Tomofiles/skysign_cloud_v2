package cloudlink

import (
	"edge/pkg/edge"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PullCommand .
func PullCommand(cloud string, vehicleID, commandID string) (*edge.Command, error) {
	log.Printf("Send CLOUD data=%s\n", "{}")

	req, err := http.NewRequest(
		"POST",
		"http://"+cloud+"/api/v1/communications/"+vehicleID+"/commands/"+commandID,
		strings.NewReader("{}"),
	)
	if err != nil {
		log.Println("cloud command request error:", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("cloud command request error:", err)
		return nil, err
	}

	requestBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("cloud command response error:", err)
		return nil, err
	}

	var command edge.Command
	err = json.Unmarshal(requestBody, &command)
	if err != nil {
		log.Println("cloud command response error:", err)
		return nil, err
	}

	log.Printf("Receive CLOUD data=%s\n", requestBody)

	return &command, nil
}
