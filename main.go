package main

import (
	"email-sender/infra/config"
	"email-sender/services"
	"fmt"
)

func main() {

	destinations := []string{
		"thiago.luiz@lavid.ufpb.br",
		"tnunes@redventures.com",
		// "testetsunoda3@gmail.com",
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	// instance a SMPT service
	smptService := services.NewSMPTService(cfg.SMTPProdutos)
	for _, email := range destinations {
		err = smptService.SendEmail(email)
		if err != nil {
			fmt.Println(err)
		}
	}
	// list := []string{"testetsunoda@gmail.com"}
	// for i := range list {
	// 	destinations := []*string{
	// 		aws.String(list[i]),
	// 	}
	// 	services.SendEmail(destinations)
	// }
	// services.SendEmail2()
}
