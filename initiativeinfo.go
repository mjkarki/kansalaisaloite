package main

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
