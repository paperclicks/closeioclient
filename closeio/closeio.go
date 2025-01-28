package closeio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HttpCloseIoClient represents a Close.io client
type HttpCloseIoClient struct {
	apiKey string
}

// NewCloseIoClient creates a new Close.io client
func NewCloseIoClient(apiKey string) *HttpCloseIoClient {
	return &HttpCloseIoClient{apiKey: apiKey}
}

// Response represents the Close.io API response
type Response struct {
	Data []Lead `json:"data"`
}

// CreateOrUpdateLeads handles creating or updating leads in bulk
func (c *HttpCloseIoClient) CreateOrUpdateLead(lead ClickFlareLead) error {
	existingLead, err := c.FindLeadByName(lead.Name)
	if err != nil {
		return fmt.Errorf("failed to search for lead: %v", err)
	}

	if existingLead != nil {
		// Update existing lead
		lead.ID = existingLead.ID
		if err := c.UpdateLead(lead); err != nil {
			return fmt.Errorf("failed to update lead: %v", err)
		}
	} else {
		// Create new lead
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

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return &result.Data[0], nil
	}
	return nil, nil
}

// CreateLead creates a new lead
func (c *HttpCloseIoClient) CreateLead(lead ClickFlareLead) error {
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

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create lead: %s", body)
	}

	return nil
}

// UpdateLead updates an existing lead
func (c *HttpCloseIoClient) UpdateLead(lead ClickFlareLead) error {
	url := fmt.Sprintf("https://api.close.com/api/v1/lead/%s/", lead.ID)
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
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to update lead: %s", body)
	}

	return nil
}

// sendRequestWithRateLimit handles HTTP requests with rate limit management
func (c *HttpCloseIoClient) sendRequestWithRateLimit(req *http.Request) (*http.Response, error) {
	for {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == 429 {
			resetSeconds := parseRateLimitResetHeader(resp.Header.Get("RateLimit"))
			if resetSeconds > 0 {
				log.Printf("Rate limit reached. Waiting for %d seconds before retrying...\n", resetSeconds)
				time.Sleep(time.Duration(resetSeconds) * time.Second)
				continue
			}
		}

		return resp, nil
	}
}

// parseRateLimitResetHeader extracts the reset time in seconds from the RateLimit header
func parseRateLimitResetHeader(header string) int {
	parts := strings.Split(header, ",")
	for _, part := range parts {
		if strings.Contains(part, "reset=") {
			value := strings.TrimSpace(strings.Split(part, "=")[1])
			resetSeconds, err := strconv.Atoi(value)
			if err != nil {
				log.Printf("Error parsing reset value from RateLimit header: %v", err)
				return 0
			}
			return resetSeconds
		}
	}
	return 0
}
