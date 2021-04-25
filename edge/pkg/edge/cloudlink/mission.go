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
	log.Printf("Send CLOUD UploadMission data=%s\n", "{}")

	req, err := http.NewRequest(
		"POST",
		cloud+"/api/v1/communications/"+vehicleID+"/uploadmissions/"+commandID,
		strings.NewReader("{}"),
	)
	if err != nil {
		log.Println("cloud uploadMission request error:", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("cloud uploadMission request error:", err)
		return nil, err
	}

	requestBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("cloud uploadMission response error:", err)
		return nil, err
	}

	var uploadMission edge.UploadMission
	err = json.Unmarshal(requestBody, &uploadMission)
	if err != nil {
		log.Println("cloud uploadMission response error:", err)
		return nil, err
	}

	log.Printf("Receive CLOUD UploadMission data=%s\n", requestBody)

	log.Printf("Send CLOUD Mission data=%s\n", "{}")

	req, err = http.NewRequest(
		"GET",
		cloud+"/api/v1/missions/"+uploadMission.MissionID,
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
