package closeio

import "time"

type ClickFlareLead struct {
	DisplayName                string                 `json:"display_name,omitempty"`
	Description                string                 `json:"description,omitempty"`
	URL                        string                 `json:"url,omitempty"`
	StatusID                   string                 `json:"status_id,omitempty"`
	OrganizationID             string                 `json:"organization_id,omitempty"`
	Tasks                      []interface{}          `json:"tasks,omitempty"`
	StatusLabel                string                 `json:"status_label,omitempty"`
	Name                       string                 `json:"name,omitempty"`
	ID                         string                 `json:"id,omitempty"`
	Addresses                  []interface{}          `json:"addresses,omitempty"`
	Contacts                   []Contact              `json:"contacts,omitempty"`
	Opportunities              []interface{}          `json:"opportunities,omitempty"`
	Custom                     map[string]interface{} `json:"custom,omitempty"`
	HTMLURL                    string                 `json:"html_url,omitempty"`
	IntegrationLinks           []interface{}          `json:"integration_links,omitempty"`
	ExternalID                 string                 `json:"custom.cf_7F7RY6AIklNjtHqhn5Gu2Q8GDGqBReKBRVQODBUq50f,omitempty"`
	FirstPayment               time.Time              `json:"custom.cf_aOaS3DfR4FYxEZlh2G5NZMb37MC9VTkuNkyiFiWn65u,omitempty"`
	LastLogin                  time.Time              `json:"custom.cf_mC7R0tBBYw934DSb8gD0kqiHwaakSkVcdGIUHGHzv0H,omitempty"`
	LastPayment                time.Time              `json:"custom.cf_NDjlZfvi4rXzDjOZrwq907RmhoHLJzrlMbQKhRMFku1,omitempty"`
	LeadOwner                  string                 `json:"custom.cf_uO1dswz1QHSN64tA843s7aU7medggOe6usq6bBaFgyu,omitempty"`
	PhoneNumber                string                 `json:"custom.cf_xL9fxsSPyBltQB62IomdftovC5tYmpAnB8O2qBhgGHb,omitempty"`
	RevenueLast30Days          float64                `json:"custom.cf_kj1b7kQ8QAhkqkI2o7jZFslthmbnPIAIh2SvNMOT82x,omitempty"`
	SignupDate                 time.Time              `json:"custom.cf_PZrkOw9eac6xWR8nKn8RPegFqxICZ1Gh2E9lzvh6c74,omitempty"`
	SpendLast30Days            float64                `json:"custom.cf_pBPpZ01j78I75dnEjDu8wHwmjkVa7krbpRrcVIspB1K,omitempty"`
	SubscriptionPlan           string                 `json:"custom.cf_Qv1RFfEskwzrM1koV1pXAWsIxem0SkRLSVTwoXcpIN0,omitempty"`
	SubscriptionStatus         string                 `json:"custom.cf_g3jOWIe5IIzoljjZHzbWr3l9gPT0uOLQhEUQmM55dPX,omitempty"`
	SurveyAvgMonthlySpend      string                 `json:"custom.cf_raFhZCFYgCtsStQk1DXtVRbVcgoyJLhx48sXmVmlLwG,omitempty"`
	SurveyHowDidYouHearAboutUs string                 `json:"custom.cf_2uoA7IJH3E0tgeEhdxtCGV86u0MDv6uHqIe0CX6ilrT,omitempty"`
	SurveyWhatDescribesYouBest string                 `json:"custom.cf_tojmUXWfTzCk2AzEw1cvmSmUXg45N5Txfr6rYn1CP4D,omitempty"`
	TotalPayments              float64                `json:"custom.cf_Xm7071TrrPBV4HFAEVt5x9XE1KGWOyVBrkIfni2jSKL,omitempty"`
	CancellationDate           time.Time              `json:"custom.cf_iXiJfR6ADj7jjd9LkKzokkh1BDxdHMzpld9fcmHf0Vb,omitempty"`
	UserID                     string                 `json:"custom.cf_KUj0DbIykJYX5cRhRat69HsoZ52ALUOoqqB9xiCLuJX,omitempty"`
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
	Type           string `json:"type,omitempty"`
	Email          string `json:"email,omitempty"`
	IsUnsubscribed bool   `json:"is_unsubscribed,omitempty"`
}

type Phone struct {
	Type           string      `json:"type,omitempty"`
	Phone          string      `json:"phone,omitempty"`
	PhoneFormatted string      `json:"phone_formatted,omitempty"`
	Country        interface{} `json:"country,omitempty"`
}

type Custom struct {
	CancellationDate           time.Time `json:"Cancellation date"`
	ExternalID                 string    `json:"External ID,omitempty"`
	FirstPayment               time.Time `json:"First payment,omitempty"`
	LastLogin                  time.Time `json:"Last login,omitempty"`
	LastPayment                time.Time `json:"Last payment,omitempty"`
	LeadOwner                  string    `json:"Lead owner,omitempty"`
	PhoneNumber                string    `json:"Phone number,omitempty"`
	RevenueLast30Days          float64   `json:"Revenue last 30 days,omitempty"`
	SignupDate                 time.Time `json:"Signup date,omitempty"`
	SpendLast30Days            float64   `json:"Spend last 30 days,omitempty"`
	SubscriptionPlan           string    `json:"Subscription plan,omitempty"`
	SubscriptionStatus         string    `json:"Subscription status,omitempty"`
	SurveyAvgMonthlySpend      string    `json:"Survey - Avg. Monthly Spend,omitempty"`
	SurveyHowDidYouHearAboutUs string    `json:"Survey - How did you hear about us,omitempty"`
	SurveyWhatDescribesYouBest string    `json:"Survey - What describes you best,omitempty"`
	TotalPayments              float64   `json:"Total payments,omitempty"`
	UserID                     string    `json:"User ID,omitempty"`
}
