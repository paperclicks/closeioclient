package closeio

import "time"

type LeadInterface interface {
	GetName() string // Example method for common behavior
	GetID() string
	SetID(id string)
	SetOwner(owner string)
	GetOwner() string
}

type Lead struct {
	DisplayName      string                 `json:"display_name,omitempty"`
	Description      string                 `json:"description,omitempty"`
	URL              string                 `json:"url,omitempty"`
	StatusID         string                 `json:"status_id,omitempty"`
	OrganizationID   string                 `json:"organization_id,omitempty"`
	Tasks            []interface{}          `json:"tasks,omitempty"`
	StatusLabel      string                 `json:"status_label,omitempty"`
	Name             string                 `json:"name,omitempty"`
	ID               string                 `json:"id,omitempty"`
	Addresses        []interface{}          `json:"addresses,omitempty"`
	Contacts         []Contact              `json:"contacts,omitempty"`
	Opportunities    []interface{}          `json:"opportunities,omitempty"`
	Custom           map[string]interface{} `json:"custom,omitempty"`
	HTMLURL          string                 `json:"html_url,omitempty"`
	IntegrationLinks []interface{}          `json:"integration_links,omitempty"`
	ClickFlareCustomFields
}

type Contact struct {
	ID               string        `json:"id,omitempty"`
	OrganizationID   string        `json:"organization_id,omitempty"`
	LeadID           string        `json:"lead_id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Title            string        `json:"title,omitempty"`
	DisplayName      string        `json:"display_name,omitempty"`
	IntegrationLinks []interface{} `json:"integration_links,omitempty"`
	Urls             []interface{} `json:"urls,omitempty"`
	Emails           []Email       `json:"emails,omitempty"`
	Phones           []Phone       `json:"phones,omitempty"`
}

type Email struct {
	Type           *string `json:"type,omitempty"`
	Email          *string `json:"email,omitempty"`
	IsUnsubscribed bool    `json:"is_unsubscribed,omitempty"`
}

type Phone struct {
	Type           *string     `json:"type,omitempty"`
	Phone          *string     `json:"phone,omitempty"`
	PhoneFormatted *string     `json:"phone_formatted,omitempty"`
	Country        interface{} `json:"country,omitempty"`
}

type Custom struct {
	CancellationDate           *time.Time `json:"Cancellation date"`
	ExternalID                 string     `json:"External ID,omitempty"`
	FirstPayment               *time.Time `json:"First payment,omitempty"`
	LastLogin                  *time.Time `json:"Last login,omitempty"`
	LastPayment                *time.Time `json:"Last payment,omitempty"`
	LeadOwner                  string     `json:"Lead owner,omitempty"`
	PhoneNumber                string     `json:"Phone number,omitempty"`
	RevenueLast30Days          float64    `json:"Revenue last 30 days,omitempty"`
	SignupDate                 *time.Time `json:"Signup date,omitempty"`
	SpendLast30Days            float64    `json:"Spend last 30 days,omitempty"`
	SubscriptionPlan           string     `json:"Subscription plan,omitempty"`
	SubscriptionStatus         string     `json:"Subscription status,omitempty"`
	SurveyAvgMonthlySpend      string     `json:"Survey - Avg. Monthly Spend,omitempty"`
	SurveyHowDidYouHearAboutUs string     `json:"Survey - How did you hear about us,omitempty"`
	SurveyWhatDescribesYouBest string     `json:"Survey - What describes you best,omitempty"`
	TotalPayments              float64    `json:"Total payments,omitempty"`
	UserID                     string     `json:"User ID,omitempty"`
}

type User struct {
	DateCreated           *time.Time  `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated"`
	Email                 string      `json:"email"`
	EmailVerifiedAt       interface{} `json:"email_verified_at"`
	FirstName             interface{} `json:"first_name"`
	GoogleProfileImageURL interface{} `json:"google_profile_image_url"`
	ID                    string      `json:"id"`
	Image                 interface{} `json:"image"`
	LastName              interface{} `json:"last_name"`
	LastUsedTimezone      interface{} `json:"last_used_timezone"`
	Organizations         []string    `json:"organizations"`
}

type SearchResponse struct {
	Data []Lead `json:"data"`
}
