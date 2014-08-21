package requests

type GetResultsFromSearch struct {
	Guts       int    `json:"guts"`
	PPOverride bool   `json:"ppOverride"`
	Query      string `json:"query"`
	Type       string `json:"type"`
}
