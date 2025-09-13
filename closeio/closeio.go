package closeio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
)

// HttpCloseIoClient represents a Close.io client
type HttpCloseIoClient struct {
	apiKey string
}

// NewCloseIoClient creates a new Close.io client
func NewCloseIoClient(apiKey string) *HttpCloseIoClient {
	return &HttpCloseIoClient{apiKey: apiKey}
}

// LeadResponse represents the Close.io API response
type LeadResponse struct {
	Data []Lead `json:"data"`
}

// UserResponse represents the Close.io API response
type UserResponse struct {
	Data []User `json:"data"`
}

// CreateOrUpdateLeads handles creating or updating leads in bulk
func (c *HttpCloseIoClient) CreateOrUpdateLead(lead LeadInterface, leadOwner string) error {
	existingLead, err := c.SearchLead(lead.GetName())
	if err != nil {
		return fmt.Errorf("failed to search for lead: %v", err)
	}

	if existingLead != nil {
		// Update existing lead
		lead.SetID(existingLead.GetID())

		//if the lead that we are updating exists but does not have an owner, and also a leadOwner has been generated, then set the lead owner
		if existingLead.GetOwner() == "" && leadOwner != "" {
			lead.SetOwner(leadOwner)
		}

		//make sure to remove any duplicate contacts that already exists
		lead.RemoveDuplicatedContacts(existingLead.GetContacts())

		if err := c.UpdateLead(lead); err != nil {
			return fmt.Errorf("failed to update lead: %v", err)
		}
	} else {
		// Create new lead
		// Set a lead owner
		lead.SetOwner(leadOwner)
		if err := c.CreateLead(lead); err != nil {
			return fmt.Errorf("failed to create lead: %v", err)
		}
	}

	return nil
}

// CreateOrUpdateLeadsV2 does not search for an existing lead but instead expects the existingLead object as an input.
// If existing lead is nil then it creates a new lead, otherwise updates the existingLead with the details contained in "lead"
func (c *HttpCloseIoClient) CreateOrUpdateLeadV2(lead LeadInterface, existingLead LeadInterface, leadOwner string) error {

	if existingLead != nil && existingLead.GetID() != "" {
		// Update existing lead
		lead.SetID(existingLead.GetID())

		//if the lead that we are updating exists but does not have an owner, and also a leadOwner has been generated, then set the lead owner
		if existingLead.GetOwner() == "" && leadOwner != "" {
			lead.SetOwner(leadOwner)
		}

		//make sure to remove any duplicate contacts that already exists
		lead.RemoveDuplicatedContacts(existingLead.GetContacts())

		if err := c.UpdateLead(lead); err != nil {
			return fmt.Errorf("failed to update lead: %v", err)
		}
	} else {
		// Create new lead
		// Set a lead owner
		lead.SetOwner(leadOwner)
		if err := c.CreateLead(lead); err != nil {
			return fmt.Errorf("failed to create lead: %v", err)
		}
	}

	return nil
}

// FindLeadByName searches for a lead by its name
func (c *HttpCloseIoClient) FindLeadByName(name string) (*Lead, error) {
	url := fmt.Sprintf("https://api.close.com/api/v1/lead/?query=%s", url.QueryEscape(name))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result LeadResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return &result.Data[0], nil
	}
	return nil, nil
}

// FindLeadByNameOrContact searches for a lead by its name or contact name
func (c *HttpCloseIoClient) FindLeadByNameOrContact(name string) (*Lead, error) {
	query := fmt.Sprintf("name:\"%s\" OR contacts.name:\"%s\"", name, name)
	url := fmt.Sprintf("https://api.close.com/api/v1/lead/?query=%s", query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result LeadResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return &result.Data[0], nil
	}
	return nil, nil
}

// CreateLead creates a new lead
func (c *HttpCloseIoClient) CreateLead(lead LeadInterface) error {
	data, err := json.Marshal(lead)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.close.com/api/v1/lead/", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create lead: %s", body)
	}

	return nil
}

// UpdateLead updates an existing lead
func (c *HttpCloseIoClient) UpdateLead(lead LeadInterface) error {
	url := fmt.Sprintf("https://api.close.com/api/v1/lead/%s/", lead.GetID())
	data, err := json.Marshal(lead)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update lead: %s", body)
	}

	return nil
}

// UpdateLead updates an existing lead
func (c *HttpCloseIoClient) GetUsers() ([]User, error) {

	url := "https://api.close.com/api/v1/user/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return result.Data, nil
	}

	return result.Data, err
}

// sendRequestWithRateLimit handles HTTP requests with rate limit management
func (c *HttpCloseIoClient) sendRequestWithRateLimit(req *http.Request) (*http.Response, error) {
	for {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("sendRequestWithRateLimit error: %s", err.Error())

			return nil, err
		}

		if resp.StatusCode == 429 {
			resetSeconds := parseRateLimitResetHeader(resp.Header.Get("RateLimit"))
			if resetSeconds > 0 {
				fmt.Printf("Rate limit reached. Waiting for %d seconds before retrying...\n", resetSeconds)
				time.Sleep(time.Duration(resetSeconds) * time.Second)
				continue
			}
		}

		return resp, nil
	}
}

// parseRateLimitResetHeader extracts the reset time in seconds from the RateLimit header
func parseRateLimitResetHeader(header string) int {

	parts := strings.Split(header, ";")
	for _, part := range parts {
		if strings.Contains(part, "rate_reset=") {
			value := strings.TrimSpace(strings.Split(part, "=")[1])
			resetSeconds, err := strconv.Atoi(value)
			if err != nil {
				log.Printf("Error parsing rate_reset value from RateLimit header: %v", err)
				return 0
			}
			log.Printf("RATE LIMIT HEADERS: %s - rate_reset: %d", header, resetSeconds)
			return resetSeconds
		}
	}
	return 0
}

// SearchLead searches for a lead or contact by name and returns the first matching lead ID
func (c *HttpCloseIoClient) SearchLead(name string) (LeadInterface, error) {
	url := "https://api.close.com/api/v1/data/search/"

	queryPayload := fmt.Sprintf(`{
		"limit": null,
		"query": {
			"negate": false,
			"queries": [
				{
					"negate": false,
					"object_type": "lead",
					"type": "object_type"
				},
				{
					"negate": false,
					"queries": [
						{
							"negate": false,
							"queries": [
								{
									"negate": false,
									"queries": [
										{
											"condition": {
												"mode": "full_words",
												"type": "text",
												"value": "%s"
											},
											"field": {
												"field_name": "name",
												"object_type": "lead",
												"type": "regular_field"
											},
											"negate": false,
											"type": "field_condition"
										}
									],
									"type": "and"
								},
								{
									"negate": false,
									"related_object_type": "contact",
									"related_query": {
										"negate": false,
										"queries": [
											{
												"condition": {
													"mode": "full_words",
													"type": "text",
													"value": "%s"
												},
												"field": {
													"field_name": "name",
													"object_type": "contact",
													"type": "regular_field"
												},
												"negate": false,
												"type": "field_condition"
											}
										],
										"type": "and"
									},
									"this_object_type": "lead",
									"type": "has_related"
								}
							],
							"type": "or"
						}
					],
					"type": "and"
				}
			],
			"type": "and"
		},
		      "_fields": {
    "lead": ["contacts","custom", "display_name","description","url","status_id","organization_id","tasks","status_label","name","id","addresses","contacts","opportunities","custom","html_url","integration_links"]
  },
		"results_limit": null,
		"sort": []
	}`, name, name)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(queryPayload)))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to search lead: %s", body)
	}

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return result.Data[0], nil
	}
	return nil, nil
}

// Search gets in input a Close.io JSON search object, searches for a leads or contacts and returns the first matching lead ID
func (c *HttpCloseIoClient) Search(query string, target interface{}) error {
	url := "https://api.close.com/api/v1/data/search/"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.sendRequestWithRateLimit(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to search lead - status code: %d - body: %s", resp.StatusCode, body)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// normalizePhoneNumber tries to parse a phone number without a region. If it fails, it assumes a default region.
func NormalizePhoneNumber(number string, defaultRegion string) (string, error) {
	// Remove spaces and unnecessary characters
	number = strings.ReplaceAll(number, " ", "")

	// If the number starts with "+", parse without region
	var parsed *phonenumbers.PhoneNumber
	var err error
	if strings.HasPrefix(number, "+") {
		parsed, err = phonenumbers.Parse(number, "")
	} else {
		parsed, err = phonenumbers.Parse(number, defaultRegion)
	}

	if err != nil {
		return "", err
	}

	// Format to E.164 (e.g., +66969371167)
	return phonenumbers.Format(parsed, phonenumbers.E164), nil
}

func (c *HttpCloseIoClient) MergeLead(sourceID, destinationID string) error {
	mergeBody := map[string]string{
		"source":      sourceID,
		"destination": destinationID,
	}
	body, _ := json.Marshal(mergeBody)

	var maxRetries = 5
	var retryDelay = 2 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest("POST", "https://api.close.com/api/v1/lead/merge/", bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(c.apiKey, "")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		if resp.StatusCode == http.StatusOK {
			return nil
		}

		// Check for "participating in another merge" error
		if strings.Contains(string(respBody), "participating in another merge") {
			if attempt < maxRetries {
				//log.Printf("Merge delayed (attempt %d/%d), retrying after %v...", attempt, maxRetries, retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			return fmt.Errorf("merge failed after retries: %s", respBody)
		}

		// Other error (not recoverable)
		return fmt.Errorf("merge failed: %s", respBody)
	}

	return fmt.Errorf("unexpected merge failure for source %s to destination %s", sourceID, destinationID)
}
