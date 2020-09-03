package main

import (
	"email-sender/infra/config"
	"email-sender/services"
	"fmt"
)

func main() {

	// destinations := []string{"thiago.luiz@lavid.ufpb.br", "tnunes@redventures.com"}
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	err = services.SendEmailBySMPT(cfg.SMTPProdutos)
	if err != nil {
		fmt.Println(err)
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
