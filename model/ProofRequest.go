package model

import "github.com/metadiv-io/zkpf/pkg/zkp"

type ProofRequest struct {
	ActualValue int             `json:"actual_value"`
	Proof       zkp.ProofResult `json:"proof"`
}
