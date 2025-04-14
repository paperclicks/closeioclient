package closeio

import "time"

type LanderLabLead struct {
	Lead
	LanderLabCustomFields
}

type LanderLabCustomFields struct {
	ActiveDomains             int        `json:"custom.cf_6Xh3qRiC4mMacLjepnBxpEEYZYURFGRi4nbnFkSZcDf,omitempty"`
	ActiveIntegrations        string     `json:"custom.cf_ipkBqxsZrCqQKdKuSG9rdjc4VEFyrQjqMwWGEmZOeF4,omitempty"`
	CancellationDate          *time.Time `json:"custom.cf_83J7o0Hon3LnOgh0GMxRLgwZpXnw0xI4nxXDVgUpYRJ,omitempty"`
	FirstPayment              *time.Time `json:"custom.cf_NUnxmj02CWcwAF2Zcco486pM6t4dWzeqTapkb88G6zU,omitempty"`
	IsFreeTrial               string     `json:"custom.cf_sr1m3Lakq40kwp3vNgDGQfprgHBEOrsRl42CQEntDG8,omitempty"`
	Language                  string     `json:"custom.cf_OYezLzcPvIkMdRl4m3c20m73nmxaLqjfgOkhdFW3VTz,omitempty"`
	LastLogin                 *time.Time `json:"custom.cf_1xMxx2jhQFWi6jLeCvUz0i4nbwMGDQgJN5jBKipHgpP,omitempty"`
	LastPayment               *time.Time `json:"custom.cf_2cayyuVTuUUTWmparCkZY4Vp78NTbgMkfMO0yTddM3D,omitempty"`
	LeadOwner                 string     `json:"custom.cf_TuqTJaOCR8t8bw2XWzDeydM37VqkgUPA0G59m5FR2VR,omitempty"`
	OnboardingExperienceLevel string     `json:"custom.cf_OskMNeSpHhbc2xOWK15ZnXyi48UrTZNGmRycvsmhkVH,omitempty"`
	OnboardingIntegrations    string     `json:"custom.cf_I2QK09KWWwXF2Z86uNMiF7uFxzgdFqGxNj6MgqN5v8X,omitempty"`
	OnboardingOccupation      string     `json:"custom.cf_oXjFydRV9K2Kv0qBrNunH4SJ7YvtCNqrubi0ZPfktAz,omitempty"`
	OnboardingPurpose         string     `json:"custom.cf_kAizsNqdglTJ23FNnCVZm346Mbp1jrFfxAF4t54GtWv,omitempty"`
	PreferredContact          string     `json:"custom.cf_JGud216nhflYTauWpkumJZn6MUT3Arhf1mqFXzQtcUp,omitempty"`
	PreferredContactMethod    string     `json:"custom.cf_mp9J6481WDi6V5T9vetje0Kvp0sGsNWSsfwQAPDFkgA,omitempty"`
	SignupDate                *time.Time `json:"custom.cf_Z9vma76NJFCi7K13jVyfBRQaco2z5uTv6GYLDh4NFAm,omitempty"`
	SubscriptionPlan          string     `json:"custom.cf_r8F5q4ulyXLv1YigseTo7pbx10kQXW2Gnzpl3rS2AET,omitempty"`
	SubscriptionStatus        string     `json:"custom.cf_sziRp7jKr6BqbAFELqLqiS6F8f40HXo6vD264Sv76pm,omitempty"`
	TotalPayments             int        `json:"custom.cf_WCtcQbYL4Jen8z3XT5kQ0IIY0gTi8oKWk76Pt49ymc0,omitempty"`
	Username                  string     `json:"custom.cf_ot7qBeR8O2wYxcFyVFIjGIRZjQKL6TAHuVPsBSDNLJ4,omitempty"`
	Websites                  int        `json:"custom.cf_bnaDqczt7t8HYTRfUv3AKsLPRv7OdKWa4ltqIJzHGFn,omitempty"`
}

func (l *LanderLabLead) GetName() string {
	return l.Name
}

func (l *LanderLabLead) GetID() string {
	return l.ID
}

func (l *LanderLabLead) SetID(id string) {
	l.ID = id
}

func (l *LanderLabLead) SetOwner(owner string) {
	l.LanderLabCustomFields.LeadOwner = owner
}

func (l *LanderLabLead) GetOwner() string {
	return l.LanderLabCustomFields.LeadOwner
}

func (l *LanderLabLead) RemoveDuplicatedContacts(existingContacts []Contact) error {

	if len(existingContacts) == 0 {
		return nil
	}

	newContacts := []Contact{}
	var phone, email, existingEmail, existingPhone string
	//var err error
	for _, contact := range l.Contacts {

		contactExists := false

		email = *contact.Emails[0].Email

		phone, _ = NormalizePhoneNumber(*contact.Phones[0].Phone, "US")
		// if err != nil {
		// 	return err
		// }

		//check if there is a contact with the same email and phone number already
		for _, existingContact := range existingContacts {
			if len(existingContact.Emails) > 0 {
				existingEmail = *existingContact.Emails[0].Email
			}
			if len(existingContact.Phones) > 0 {
				existingPhone, _ = NormalizePhoneNumber(*existingContact.Phones[0].Phone, "US")
				// if err != nil {
				// 	return err
				// }

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
