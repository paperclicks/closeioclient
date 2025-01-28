package closeio

import "time"

type OptimizerFlareLead struct {
	Lead
	OptimizerCustomFields
}

type OptimizerCustomFields struct {
	ExternalID                 string    `json:"omitempty"`
	FirstPayment               time.Time `json:"omitempty"`
	LastLogin                  time.Time `json:"omitempty"`
	LastPayment                time.Time `json:"omitempty"`
	LeadOwner                  string    `json:"omitempty"`
	PhoneNumber                string    `json:"omitempty"`
	RevenueLast30Days          float64   `json:"omitempty"`
	SignupDate                 time.Time `json:"omitempty"`
	SpendLast30Days            float64   `json:"omitempty"`
	SubscriptionPlan           string    `json:"omitempty"`
	SubscriptionStatus         string    `json:"omitempty"`
	SurveyAvgMonthlySpend      string    `json:"omitempty"`
	SurveyHowDidYouHearAboutUs string    `json:"omitempty"`
	SurveyWhatDescribesYouBest string    `json:"omitempty"`
	TotalPayments              float64   `json:"omitempty"`
	CancellationDate           time.Time `json:"omitempty"`
	UserID                     string    `json:"omitempty"`
}

func (l OptimizerFlareLead) GetName() string {
	return l.Name
}
