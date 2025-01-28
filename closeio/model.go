package closeio

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Lead represents the structure of a Close.io lead type
type Lead struct {
	ID       string                 `json:"id,omitempty"`
	Name     string                 `json:"name"`
	Contacts []Contact              `json:"contacts,omitempty"`
	Custom   map[string]interface{} `json:"custom,omitempty"`
}

// Contact represents a contact in Close.io
// type Contact struct {
// 	ID     string                 `json:"id,omitempty"`
// 	Name   string                 `json:"name,omitempty"`
// 	Emails []string               `json:"emails,omitempty"`
// 	Phones []string               `json:"phones,omitempty"`
// 	Custom map[string]interface{} `json:"custom,omitempty"`
// }

func JSONToLead(ser []byte) (*Lead, error) {

	var lead Lead
	err := json.Unmarshal(ser, &lead)
	if err != nil {
		return nil, err
	}

	lead.Custom = make(map[string]interface{})
	var keyValues map[string]interface{}
	err = json.Unmarshal(ser, &keyValues)
	if err != nil {
		return nil, err
	}

	for key, value := range keyValues {
		if strings.HasPrefix(key, "custom.") {
			lead.Custom[strings.Replace(key, "custom.", "", 1)] = value
		}
	}

	return &lead, nil
}

func LeadToJSON(lead Lead) ([]byte, error) {

	bytes, _ := json.Marshal(lead)
	var dynamicFields map[string]interface{}
	json.Unmarshal(bytes, &dynamicFields)

	for key, value := range lead.Custom {
		dynamicFields[fmt.Sprintf("custom.%s", key)] = value
	}

	return json.Marshal(dynamicFields)
}
