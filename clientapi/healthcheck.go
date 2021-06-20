package clientapi

import (
	"encoding/json"
	"fmt"
)

// HealthService handles communication with the healthcheck
type HealthService service

// Status represents a healthcheck status.
type Status struct {
	Status string `json:"status,omitempty"`
}

// Check fetches endpoint health status.
func (s *HealthService) Check() (*Status, error) {
	location := "v1/health"
	url := fmt.Sprintf("%v%v", s.client.BaseURL, location)

	body, err := s.client.sendRequest("GET", url, HealthOK, nil)
	if err != nil {
		return nil, err
	}

	responseData := new(Status)
	json.Unmarshal(body, &responseData)

	return responseData, nil
}
