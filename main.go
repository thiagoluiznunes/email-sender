package main

import (
	"email-sender/infra/config"
	"email-sender/infra/utils"
	"email-sender/services"
	"fmt"
	"os"
)

func main() {

	destinations, err := utils.ReadCSFile("assets/csv/test-santander.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	// instance a SMPT service
	smptService := services.NewSMPTService(cfg.SMTPProdutos)
	if err != nil {
		fmt.Println(err)
	}

	for _, email := range destinations {
		smptService.SendEmail(email)
	}
}
