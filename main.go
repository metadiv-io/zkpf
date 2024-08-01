package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/metadiv-io/zkpf/pkg/zkp"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("ZKP Range Proofer")
	fmt.Println("-----------------")
	fmt.Print("Please enter the path of the proof result file: ")

	path, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	path = strings.ReplaceAll(path, "\n", "")
	s, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Error: ", err, "("+path+")")
	}

	result := zkp.ProofResult{}
	err = result.FromString(string(s))
	if err != nil || result.Verifier == "" || result.Proof1 == "" || result.Proof2 == "" {
		log.Fatalln("Error: Invalid proof result object")
	}

	fmt.Print("Please enter the real data in number: ")
	realData, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	realData = strings.ReplaceAll(realData, "\n", "")
	if realData == "" {
		log.Fatalln("Error: Invalid real data")
	}

	realNumber, err := strconv.Atoi(realData)
	if err != nil {
		log.Fatalln("Error: Invalid real data")
	}

	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("Proof Result:")
	fmt.Println("-------------")
	fmt.Println("InRange: ", result.InRange)
	fmt.Println("MinValue: ", result.MinValue)
	fmt.Println("MaxValue: ", result.MaxValue)
	fmt.Println("RandNum: ", result.RandNum)
	fmt.Println("RealData: ", realNumber)
	fmt.Println("-------------")
	fmt.Println("Result verification: ", zkp.VerifyProofResultWithSecret(result, realNumber))
	fmt.Println("-------------")
	fmt.Println("")
}
