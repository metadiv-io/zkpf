package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/zkpf/handler"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.UI)
	r.POST("/proof", handler.Proof)

	go func() {
		if err := r.Run(":5001"); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	fmt.Println("Server is running on http://localhost:5001")
	openBrowser("http://localhost:5001")

	select {}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Error opening browser: %v", err)
	}
}
