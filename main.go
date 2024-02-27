package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const URL string = "https://www.kansalaisaloite.fi/api/v1/initiatives?limit=50&orderBy=createdNewest&offset="

type Initiative struct {
	ID                       string                 `json:"id"`
	Modified                 string                 `json:"modified"`
	State                    string                 `json:"state"`
	StateDate                string                 `json:"stateDate"`
	AcceptanceIdentifier     interface{}            `json:"acceptanceIdentifier"`
	SupportCount             int                    `json:"supportCount"`
	SentSupportCount         int                    `json:"sentSupportCount"`
	VerificationPassed       bool                   `json:"verificationPassed"`
	ExternalSupportCount     int                    `json:"externalSupportCount"`
	VerifiedSupportCount     int                    `json:"verifiedSupportCount"`
	Verified                 interface{}            `json:"verified"`
	Name                     map[string]interface{} `json:"name"`
	StartDate                string                 `json:"startDate"`
	EndDate                  string                 `json:"endDate"`
	ProposalType             string                 `json:"proposalType"`
	PrimaryLanguage          string                 `json:"primaryLanguage"`
	FinancialSupport         bool                   `json:"financialSupport"`
	FinancialSupportURL      interface{}            `json:"financialSupportURL"`
	SupportStatementsOnPaper bool                   `json:"supportStatementsOnPaper"`
	SupportStatementsInWeb   bool                   `json:"supportStatementsInWeb"`
	SupportStatementsRemoved interface{}            `json:"supportStatementsRemoved"`
	VotingInProgress         bool                   `json:"votingInProgress"`
	URL                      map[string]string      `json:"url"`
	TotalSupportCount        int                    `json:"totalSupportCount"`
}

type InitiativeInfo struct {
	ID                       string                 `json:"id"`
	Modified                 string                 `json:"modified"`
	State                    string                 `json:"state"`
	StateDate                string                 `json:"stateDate"`
	AcceptanceIdentifier     interface{}            `json:"acceptanceIdentifier"`
	SupportCount             int                    `json:"supportCount"`
	SentSupportCount         int                    `json:"sentSupportCount"`
	VerificationPassed       bool                   `json:"verificationPassed"`
	ExternalSupportCount     int                    `json:"externalSupportCount"`
	VerifiedSupportCount     int                    `json:"verifiedSupportCount"`
	Verified                 interface{}            `json:"verified"`
	Name                     map[string]interface{} `json:"name"`
	StartDate                string                 `json:"startDate"`
	EndDate                  string                 `json:"endDate"`
	ProposalType             string                 `json:"proposalType"`
	PrimaryLanguage          string                 `json:"primaryLanguage"`
	FinancialSupport         bool                   `json:"financialSupport"`
	FinancialSupportURL      interface{}            `json:"financialSupportURL"`
	SupportStatementsOnPaper bool                   `json:"supportStatementsOnPaper"`
	SupportStatementsInWeb   bool                   `json:"supportStatementsInWeb"`
	SupportStatementsRemoved interface{}            `json:"supportStatementsRemoved"`
	Links                    []map[string]string    `json:"links"`
	Proposal                 map[string]interface{} `json:"proposal"`
	Rationale                map[string]interface{} `json:"rationale"`
	Initiators               []interface{}          `json:"initiators"`
	Representatives          []interface{}          `json:"representatives"`
	Reserves                 []interface{}          `json:"reserves"`
	Accountables             []struct {
		FirstNames       string `json:"firstNames"`
		LastName         string `json:"lastName"`
		HomeMunicipality struct {
			Fi string `json:"fi"`
			Sv string `json:"sv"`
		} `json:"homeMunicipality"`
		ContactInfo struct {
			Email   string `json:"email"`
			Phone   string `json:"phone"`
			Address string `json:"address"`
		} `json:"contactInfo"`
	} `json:"accountables"`
	URL struct {
		Fi string `json:"fi"`
		Sv string `json:"sv"`
	} `json:"url"`
	VotingInProgress  bool `json:"votingInProgress"`
	TotalSupportCount int  `json:"totalSupportCount"`
}

type InitiativeStruct interface {
	[]Initiative | InitiativeInfo
}

func getData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Tried to fetch data from URL"+url+":", err)
	}
	return data
}

func getInitiativeStructFromJSON[T InitiativeStruct](data []byte) T {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal("Tried to convert JSON to structured data:", err)
	}
	return result
}

func main() {
	var data []byte
	var aloitteet []Initiative = []Initiative{}
	var listall = false
	var offset = 0
	var nimi string

	if len(os.Args) < 2 {
		fmt.Println("Käyttö:")
		fmt.Println("    " + os.Args[0] + ` ”hakuteksti"`)
		fmt.Println("  tai")
		fmt.Println("    " + os.Args[0] + ` -a`)
		os.Exit(0)
	}

	search := os.Args[1]

	if search == "-a" {
		fmt.Println("Listataan kaikki aloitteet:")
		listall = true
	}

	data = getData(fmt.Sprintf("%s%d", URL, offset))
	for bytes.Compare(data, []byte{91, 93}) != 0 {
		a := getInitiativeStructFromJSON[[]Initiative](data)
		aloitteet = append(aloitteet, a...)
		offset += 50
		data = getData(fmt.Sprintf("%s%d", URL, offset))
	}

	for _, initiative := range aloitteet {
		if initiative.Name["fi"] == nil {
			nimi = initiative.Name["sv"].(string)
		} else {
			nimi = initiative.Name["fi"].(string)
		}
		if listall == false {
			nimi_l := strings.ToLower(nimi)
			if strings.Contains(nimi_l, search) {
				url := initiative.ID
				data = getData(url)
				initiativeStruct := getInitiativeStructFromJSON[InitiativeInfo](data)
				fmt.Println("- " + nimi + ":")
				fmt.Println("  Kannatusilmoituksia:", initiativeStruct.SupportCount)
				fmt.Println("  Keräys päättyy:     ", initiativeStruct.EndDate)
			}
		} else {
			fmt.Println("- " + nimi)
		}
	}
}
