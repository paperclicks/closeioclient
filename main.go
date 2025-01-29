package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paperclicks/closeioclient/closeio"
)

func main() {
	client := closeio.NewCloseIoClient("api_6ofO4BdefApUqQNHTg66kb.4UDSTDXWXka8vy8CHKXqZW")

	// 	leadBody := `{
	//     "tasks": [],
	//     "html_url": "https://app.close.com/lead/lead_lKJosXmlzl9PUmVix7Wm715Y9N6ABZC3tsvy3ihM01v/",
	//     "url": "http://www.paperclicks.net",
	//     "integration_links": [],
	//     "opportunities": [],
	//     "contacts": [
	//         {
	//             "id": "cont_jyK9N3CuAcPsH8uWAyiRJaW23urGEtqhu1xwq8H1m4g",
	//             "organization_id": "orga_n7pQjacgbTaEMGIyAi16g07RclEM2ZyoC4N7E9QuQND",
	//             "lead_id": "lead_lKJosXmlzl9PUmVix7Wm715Y9N6ABZC3tsvy3ihM01v",
	//             "created_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//             "updated_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//             "date_created": "2025-01-22T21:13:49.107000+00:00",
	//             "date_updated": "2025-01-26T15:51:53.211000+00:00",
	//             "name": "Ervin Hoxha",
	//             "title": "",
	//             "display_name": "Ervin Hoxha",
	//             "integration_links": [],
	//             "urls": [],
	//             "emails": [
	//                 {
	//                     "type": "office",
	//                     "email": "ervin@paperclicks.net",
	//                     "is_unsubscribed": false
	//                 }
	//             ],
	//             "phones": [
	//                 {
	//                     "type": "office",
	//                     "phone": "+355696867650",
	//                     "phone_formatted": "+355 69 686 7650",
	//                     "country": "AL"
	//                 }
	//             ]
	//         },
	//         {
	//             "id": "cont_5Khpu0NcfyCoB9VNl60LvE4zL5f7owPFKRYDxEx6NSH",
	//             "organization_id": "orga_n7pQjacgbTaEMGIyAi16g07RclEM2ZyoC4N7E9QuQND",
	//             "lead_id": "lead_lKJosXmlzl9PUmVix7Wm715Y9N6ABZC3tsvy3ihM01v",
	//             "created_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//             "updated_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//             "date_created": "2025-01-22T20:50:34.938000+00:00",
	//             "date_updated": "2025-01-26T15:51:53.218000+00:00",
	//             "name": "Ervin Hoxha",
	//             "title": "",
	//             "display_name": "Ervin Hoxha",
	//             "integration_links": [],
	//             "urls": [],
	//             "emails": [
	//                 {
	//                     "type": "office",
	//                     "email": "ervin@paperclicks.net",
	//                     "is_unsubscribed": false
	//                 }
	//             ],
	//             "phones": [
	//                 {
	//                     "type": "office",
	//                     "phone": "+355696867660",
	//                     "phone_formatted": "+355 69 686 7660",
	//                     "country": "AL"
	//                 }
	//             ]
	//         }
	//     ],
	//     "date_created": "2025-01-22T20:50:34.911000+00:00",
	//     "addresses": [],
	//     "created_by_name": "Ervin Hoxha",
	//     "organization_id": "orga_n7pQjacgbTaEMGIyAi16g07RclEM2ZyoC4N7E9QuQND",
	//     "date_updated": "2025-01-26T15:54:02.911000+00:00",
	//     "status_label": "Customer",
	//     "updated_by_name": "Ervin Hoxha",
	//     "id": "lead_lKJosXmlzl9PUmVix7Wm715Y9N6ABZC3tsvy3ihM01v",
	//     "name": "ClickFlare",
	//     "description": "",
	//     "display_name": "ClickFlare",
	//     "created_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//     "updated_by": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//     "status_id": "stat_TkMiJDUeDWKewUTD6b25ZlLaFeSTY9KAb2ua6v5cjQo",
	//     "custom.cf_iXiJfR6ADj7jjd9LkKzokkh1BDxdHMzpld9fcmHf0Vb": "2025-01-26T11:00:00+00:00",
	//     "custom.cf_7F7RY6AIklNjtHqhn5Gu2Q8GDGqBReKBRVQODBUq50f": "xxxxxxxy",
	//     "custom.cf_aOaS3DfR4FYxEZlh2G5NZMb37MC9VTkuNkyiFiWn65u": "2025-01-22T11:00:00+00:00",
	//     "custom.cf_mC7R0tBBYw934DSb8gD0kqiHwaakSkVcdGIUHGHzv0H": "2025-01-22T11:00:00+00:00",
	//     "custom.cf_NDjlZfvi4rXzDjOZrwq907RmhoHLJzrlMbQKhRMFku1": "2025-01-27T11:00:00+00:00",
	//     "custom.cf_uO1dswz1QHSN64tA843s7aU7medggOe6usq6bBaFgyu": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
	//     "custom.cf_xL9fxsSPyBltQB62IomdftovC5tYmpAnB8O2qBhgGHb": "00355696867660",
	//     "custom.cf_kj1b7kQ8QAhkqkI2o7jZFslthmbnPIAIh2SvNMOT82x": 200000,
	//     "custom.cf_PZrkOw9eac6xWR8nKn8RPegFqxICZ1Gh2E9lzvh6c74": "2025-01-22T11:00:00+00:00",
	//     "custom.cf_pBPpZ01j78I75dnEjDu8wHwmjkVa7krbpRrcVIspB1K": 100000,
	//     "custom.cf_Qv1RFfEskwzrM1koV1pXAWsIxem0SkRLSVTwoXcpIN0": "ClickFlare Agency",
	//     "custom.cf_g3jOWIe5IIzoljjZHzbWr3l9gPT0uOLQhEUQmM55dPX": "active",
	//     "custom.cf_raFhZCFYgCtsStQk1DXtVRbVcgoyJLhx48sXmVmlLwG": "10000",
	//     "custom.cf_2uoA7IJH3E0tgeEhdxtCGV86u0MDv6uHqIe0CX6ilrT": "Google",
	//     "custom.cf_tojmUXWfTzCk2AzEw1cvmSmUXg45N5Txfr6rYn1CP4D": "Affiliate",
	//     "custom.cf_Xm7071TrrPBV4HFAEVt5x9XE1KGWOyVBrkIfni2jSKL": 89.5,
	//     "custom.cf_KUj0DbIykJYX5cRhRat69HsoZ52ALUOoqqB9xiCLuJX": "123456789"
	// }`

	newLeadBody := `{
    "contacts": [
        {
            "name": "Ervin Hoxha New",
            "title": "",
            "display_name": "Ervin Hoxha new",
            "emails": [
                {
                    "type": "office",
                    "email": "ervin2@paperclicks.net"
                }
            ],
            "phones": [
                {
                    "type": "office",
                    "phone": "+355696867650"
                }
            ]
        }
    ],
    "name": "TheOptimizer",
    "description": "created lead",
    "display_name": "ClickFlare",
    "status_id": "stat_TkMiJDUeDWKewUTD6b25ZlLaFeSTY9KAb2ua6v5cjQo",
    "custom.cf_iXiJfR6ADj7jjd9LkKzokkh1BDxdHMzpld9fcmHf0Vb": "2025-01-26T11:00:00+00:00",
    "custom.cf_7F7RY6AIklNjtHqhn5Gu2Q8GDGqBReKBRVQODBUq50f": "xxxxxxxy",
    "custom.cf_aOaS3DfR4FYxEZlh2G5NZMb37MC9VTkuNkyiFiWn65u": "2025-01-22T11:00:00+00:00",
    "custom.cf_mC7R0tBBYw934DSb8gD0kqiHwaakSkVcdGIUHGHzv0H": "2025-01-22T11:00:00+00:00",
    "custom.cf_NDjlZfvi4rXzDjOZrwq907RmhoHLJzrlMbQKhRMFku1": "2025-01-27T11:00:00+00:00",
    "custom.cf_uO1dswz1QHSN64tA843s7aU7medggOe6usq6bBaFgyu": "user_oGHxhqyE3cr2GyfR1ztuLpJnQqdXb44bSyOEflE6Lj0",
    "custom.cf_xL9fxsSPyBltQB62IomdftovC5tYmpAnB8O2qBhgGHb": "00355696867660",
    "custom.cf_kj1b7kQ8QAhkqkI2o7jZFslthmbnPIAIh2SvNMOT82x": 200000,
    "custom.cf_PZrkOw9eac6xWR8nKn8RPegFqxICZ1Gh2E9lzvh6c74": "2025-01-22T11:00:00+00:00",
    "custom.cf_pBPpZ01j78I75dnEjDu8wHwmjkVa7krbpRrcVIspB1K": 100000,
    "custom.cf_Qv1RFfEskwzrM1koV1pXAWsIxem0SkRLSVTwoXcpIN0": "ClickFlare Agency",
    "custom.cf_g3jOWIe5IIzoljjZHzbWr3l9gPT0uOLQhEUQmM55dPX": "active",
    "custom.cf_raFhZCFYgCtsStQk1DXtVRbVcgoyJLhx48sXmVmlLwG": "10000",
    "custom.cf_2uoA7IJH3E0tgeEhdxtCGV86u0MDv6uHqIe0CX6ilrT": "Google",
    "custom.cf_tojmUXWfTzCk2AzEw1cvmSmUXg45N5Txfr6rYn1CP4D": "Affiliate",
    "custom.cf_Xm7071TrrPBV4HFAEVt5x9XE1KGWOyVBrkIfni2jSKL": 89.5,
    "custom.cf_KUj0DbIykJYX5cRhRat69HsoZ52ALUOoqqB9xiCLuJX": "123456789"
}`

	var newLead closeio.ClickFlareLead

	existingLead := closeio.ClickFlareLead{}
	existingLead.Name = "test test"
	if err := client.CreateOrUpdateLead(&existingLead); err != nil {
		log.Fatalf("Error creating or updating leads: %v", err)
	}

	users, err := client.GetUsers()
	fmt.Println(users)

	err = json.Unmarshal([]byte(newLeadBody), &newLead)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if err := client.CreateOrUpdateLead(&newLead); err != nil {
		log.Fatalf("Error creating or updating leads: %v", err)
	}
}
