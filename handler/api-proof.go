package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/zkpf/model"
	"github.com/metadiv-io/zkpf/pkg/zkp"
)

func Proof(c *gin.Context) {
	var req model.ProofRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	ok := zkp.VerifyProofResultWithSecret(req.Proof, req.ActualValue)
	result := model.ProofResponse{
		ActualValue: req.ActualValue,
		MaxValue:    req.Proof.MaxValue,
		MinValue:    req.Proof.MinValue,
		Valid:       ok,
	}

	c.JSON(200, result)
}
