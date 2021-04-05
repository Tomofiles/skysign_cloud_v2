package cloudlink

import (
	"edge/pkg/edge"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PullMission .
func PullMission(cloud string, vehicleID string, commandID string) (*edge.Mission, error) {
	log.Printf("Send CLOUD Communication data=%s\n", "{}")

	req, err := http.NewRequest(
		"POST",
		cloud+"/api/v1/communications/"+vehicleID+"/uploadmissions/"+commandID,
		strings.NewReader("{}"),
	)
	if err != nil {
		log.Println("cloud communication request error:", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("cloud communication request error:", err)
		return nil, err
	}

	requestBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("cloud communication response error:", err)
		return nil, err
	}

	var communication edge.Communication
	err = json.Unmarshal(requestBody, &communication)
	if err != nil {
		log.Println("cloud communication response error:", err)
		return nil, err
	}

	log.Printf("Receive CLOUD Communication data=%s\n", requestBody)

	log.Printf("Send CLOUD Mission data=%s\n", "{}")

	req, err = http.NewRequest(
		"GET",
		cloud+"/api/v1/missions/"+communication.MissionID,
		strings.NewReader("{}"),
	)
	if err != nil {
		log.Println("cloud mission request error:", err)
		return nil, err
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("cloud mission request error:", err)
		return nil, err
	}

	requestBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("cloud mission response error:", err)
		return nil, err
	}

	var mission edge.Mission
	err = json.Unmarshal(requestBody, &mission)
	if err != nil {
		log.Println("cloud mission response error:", err)
		return nil, err
	}

	log.Printf("Receive CLOUD mission data=%s\n", requestBody)

	return &mission, nil
}
