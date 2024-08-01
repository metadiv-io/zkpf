package zkp

import "encoding/json"

type ProofResult struct {
	InRange  bool   `json:"in_range"`
	MinValue int    `json:"min_value"`
	MaxValue int    `json:"max_value"`
	RandNum  int    `json:"rand_num"`
	Verifier string `json:"verifier"`
	Proof1   string `json:"proof1"`
	Proof2   string `json:"proof2"`
}

func (p *ProofResult) String() string {
	s, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func (p *ProofResult) FromString(s string) error {
	err := json.Unmarshal([]byte(s), p)
	if err != nil {
		return err
	}
	return nil
}
