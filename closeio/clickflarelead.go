package closeio

import "time"

type ClickFlareLead struct {
	Lead
	ClickFlareCustomFields
}

type ClickFlareCustomFields struct {
	ExternalID                 string     `json:"custom.cf_7F7RY6AIklNjtHqhn5Gu2Q8GDGqBReKBRVQODBUq50f,omitempty"`
	FirstPayment               *time.Time `json:"custom.cf_aOaS3DfR4FYxEZlh2G5NZMb37MC9VTkuNkyiFiWn65u,omitempty"`
	LastLogin                  *time.Time `json:"custom.cf_mC7R0tBBYw934DSb8gD0kqiHwaakSkVcdGIUHGHzv0H,omitempty"`
	LastPayment                *time.Time `json:"custom.cf_NDjlZfvi4rXzDjOZrwq907RmhoHLJzrlMbQKhRMFku1,omitempty"`
	LeadOwner                  string     `json:"custom.cf_uO1dswz1QHSN64tA843s7aU7medggOe6usq6bBaFgyu,omitempty"`
	PhoneNumber                string     `json:"custom.cf_xL9fxsSPyBltQB62IomdftovC5tYmpAnB8O2qBhgGHb,omitempty"`
	RevenueLast30Days          float64    `json:"custom.cf_kj1b7kQ8QAhkqkI2o7jZFslthmbnPIAIh2SvNMOT82x,omitempty"`
	SignupDate                 *time.Time `json:"custom.cf_PZrkOw9eac6xWR8nKn8RPegFqxICZ1Gh2E9lzvh6c74,omitempty"`
	SpendLast30Days            float64    `json:"custom.cf_pBPpZ01j78I75dnEjDu8wHwmjkVa7krbpRrcVIspB1K,omitempty"`
	SubscriptionPlan           string     `json:"custom.cf_Qv1RFfEskwzrM1koV1pXAWsIxem0SkRLSVTwoXcpIN0,omitempty"`
	SubscriptionStatus         string     `json:"custom.cf_g3jOWIe5IIzoljjZHzbWr3l9gPT0uOLQhEUQmM55dPX,omitempty"`
	SurveyAvgMonthlySpend      string     `json:"custom.cf_raFhZCFYgCtsStQk1DXtVRbVcgoyJLhx48sXmVmlLwG,omitempty"`
	SurveyHowDidYouHearAboutUs string     `json:"custom.cf_2uoA7IJH3E0tgeEhdxtCGV86u0MDv6uHqIe0CX6ilrT,omitempty"`
	SurveyWhatDescribesYouBest string     `json:"custom.cf_tojmUXWfTzCk2AzEw1cvmSmUXg45N5Txfr6rYn1CP4D,omitempty"`
	TotalPayments              float64    `json:"custom.cf_Xm7071TrrPBV4HFAEVt5x9XE1KGWOyVBrkIfni2jSKL,omitempty"`
	CancellationDate           *time.Time `json:"custom.cf_iXiJfR6ADj7jjd9LkKzokkh1BDxdHMzpld9fcmHf0Vb,omitempty"`
	UserID                     string     `json:"custom.cf_KUj0DbIykJYX5cRhRat69HsoZ52ALUOoqqB9xiCLuJX,omitempty"`
}

func (l *ClickFlareLead) GetName() string {
	return l.Name
}

func (l *ClickFlareLead) GetID() string {
	return l.ID
}

func (l *ClickFlareLead) SetID(id string) {
	l.ID = id
}

func (l *ClickFlareLead) SetOwner(owner string) {
	l.ClickFlareCustomFields.LeadOwner = owner
}

func (l *ClickFlareLead) GetOwner() string {
	return l.ClickFlareCustomFields.LeadOwner
}

func (l *ClickFlareLead) RemoveDuplicatedContacts(existingContacts []Contact) {

	if len(existingContacts) == 0 {
		return
	}

	newContacts := []Contact{}

	for _, contact := range l.Contacts {

		contactExists := false

		email := *contact.Emails[0].Email
		phone := *contact.Phones[0].Phone

		//check if there is a contact with the same email and phone number already
		for _, existingContact := range existingContacts {
			existingEmail := *existingContact.Emails[0].Email
			existingPhone := *existingContact.Phones[0].Phone

			if email == existingEmail && phone == existingPhone {
				contactExists = true
			}
		}
		//if there is not a contact with same email and phone then add it to the list of new contacts
		if !contactExists {
			newContacts = append(newContacts, contact)
		}
	}

	// set contacts of the lead to the list of new contacts so that only non existing contacts will be updated.
	l.Contacts = newContacts

}
