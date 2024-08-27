package model

type ProofResponse struct {
	ActualValue int  `json:"actual_value"`
	MaxValue    int  `json:"max_value"`
	MinValue    int  `json:"min_value"`
	Valid       bool `json:"valid"`
}
