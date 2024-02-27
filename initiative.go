package main

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
