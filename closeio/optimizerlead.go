package closeio

import (
	"time"
)

type OptimizerLead struct {
	Lead
	OptimizerCustomFields
}

type OptimizerCustomFields struct {
	CancellationDate         *time.Time `json:"custom.cf_YXlpcp2Tsx6BUHZh4KAg9s3Xp7n87mrEWnQtcVBozrz,omitempty"`
	CFExternalID             string     `json:"custom.cf_feKyKqKi7EhPCAUhvfEl68SW7i2IwlhqqJ0QH94YZpT,omitempty"`
	ClientTier               string     `json:"custom.cf_52CjvyPw1ONwVxxFjxuDB9q21J2HZAXRLhHbv1EhIWG,omitempty"`
	CustomerSuccessManager   string     `json:"custom.cf_PYQx3pfPCMn17eqnq0fWm4xbgQq5Od1onpCNoUhHTWv,omitempty"`
	DaysLeft                 float64    `json:"custom.cf_JH4XTMMCsrEt8SVeE2P74Ouj9UXZkQRGhX6s4YvMoqm,omitempty"`
	FirstPayment             *time.Time `json:"custom.cf_MJbPmJXJM9pSEvSJOWKYCkrYqN3vGTV5r8HYjLk5N1K,omitempty"`
	LastLogin                *time.Time `json:"custom.cf_WqXaCN7ZkT8jGMx623ysgBFQl1Et2fykX0vxdzar36e,omitempty"`
	LastPayment              *time.Time `json:"custom.cf_7YmJjQd85K2NXaEVDkm885jrWxRAWCrNZEltNrJq7pD,omitempty"`
	PhoneNumber              string     `json:"custom.cf_8JVU7KuWQ7aMukzZEb9buBJqBLlVtlRuUX2JcHkJJuE,omitempty"`
	PreferredContact         *string    `json:"custom.cf_bz5FTvg9lL55DkNE4IseXnSqKDqc5qx4sgrsbT1L8i4,omitempty"`
	PreferredContactMethod   *string    `json:"custom.cf_NzOGLG8NoJIcezhs5hGeTryhlq9RW5uAa2dKlDNANof,omitempty"`
	LeadOwner                string     `json:"custom.cf_2Zi5ZWiSlwNwlvByIvgE3JtoUa8Yqn8FnFdAtmGeUKS,omitempty"`
	SignupDate               *time.Time `json:"custom.cf_DiSq6fOceRNpV3XhXjjByjAdMTbGiirGFXN9PfIrwbd,omitempty"`
	SubscriptionPlan         string     `json:"custom.cf_BiA9ZK6Jhg7NNG6kLrWowAJRTPD3OE7CqNhMmrffKqJ,omitempty"`
	SubscriptionStatus       string     `json:"custom.cf_H1hppSt4cz8VXAGkM1vlsYEmKW7WSCV9BXYTbSAaiIV,omitempty"`
	SurveyAvgMonthlySpend    string     `json:"custom.cf_JFvSwgHW4WuP1lKYuWsKWNc2sGlSKFtt36yZhheAsmF,omitempty"`
	SurveyHowHeard           string     `json:"custom.cf_NE02IUJfGNGlk7YOwOpAJ69FYHkPFtqYJCCKcXxgrtp,omitempty"`
	SurveyBestDescription    string     `json:"custom.cf_SE9WzW4dAxhEAMWzJTPzkESIzUSnvjMLL01iJM6zcJF,omitempty"`
	TopPayingUser            string     `json:"custom.cf_ZCQk9ScxGXOahAq1Yo7RMWT6DsnxyWjwJD9nD2KEQws,omitempty"`
	TotalPaymentsLast3Months float64    `json:"custom.cf_9o1E0Rt0014HZ88Vor16dNloL2L1bJg3VOF7Twf7Ve9,omitempty"`
	TrafficSources           []string   `json:"custom.cf_yIM1O5BtOW3zoOTXaq7flTJsu1h2u9q54PvgtaUrfql,omitempty"`
	Username                 string     `json:"custom.cf_L1oVCZYa3seTdcgR1k71ro8zC4exUkj3TGe2kQ8sHoR,omitempty"`
	TotalPayments            float64    `json:"custom.cf_iF8l4t5IaSXjt8bKhyBl9Nb1IKmWnbFwLL3QJdD6hQI,omitempty"`
}

func (l *OptimizerLead) GetName() string {
	return l.Name
}

func (l *OptimizerLead) GetID() string {
	return l.ID
}

func (l *OptimizerLead) SetID(id string) {
	l.ID = id
}

func (l *OptimizerLead) SetOwner(owner string) {
	l.LeadOwner = owner
}

func (l *OptimizerLead) GetOwner() string {
	return l.LeadOwner
}

func (l *OptimizerLead) GetContacts() []Contact {
	return l.Contacts
}

func (l *OptimizerLead) RemoveDuplicatedContacts(existingContacts []Contact) error {

	if len(existingContacts) == 0 {
		return nil
	}

	newContacts := []Contact{}
	var phone, email, existingEmail, existingPhone string
	var err error
	for _, contact := range l.Contacts {

		contactExists := false

		if len(contact.Emails) > 0 && contact.Emails[0].Email != nil {
			email = *contact.Emails[0].Email
		}

		if len(contact.Phones) > 0 && contact.Phones[0].Phone != nil {
			// phone = *contact.Phones[0].Phone
			// plusIndex := strings.Index(phone, "+")

			// phone = strings.ReplaceAll(phone, " ", "")

			// // Take substring from "+" onwards and remove all spaces
			// phone = phone[plusIndex:]

			phone, err = NormalizePhoneNumber(*contact.Phones[0].Phone, "US")
			if err != nil {
				return err
			}

		}

		//check if there is a contact with the same email and phone number already
		for _, existingContact := range existingContacts {
			if len(existingContact.Emails) > 0 {
				existingEmail = *existingContact.Emails[0].Email
			}
			if len(existingContact.Phones) > 0 {
				existingPhone, err = NormalizePhoneNumber(*existingContact.Phones[0].Phone, "US")
				if err != nil {
					return err
				}

			}

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

	return nil

}
