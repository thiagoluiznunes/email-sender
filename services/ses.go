package services

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/gomail.v2"
)

const (
	Sender = "contato@produtos.cartoes.com.br"
	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	ConfigurationSet = "production-hermes-configuration-set"
	Subject          = "Amazon SES Test (AWS SDK for Go)"
	HtmlBody         = ``
	TextBody         = "This email was sent with Amazon SES using the AWS SDK for Go."
	CharSet          = "UTF-8"
)

func SendEmail(destinations []*string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1")},
	)
	svc := ses.New(sess)
	//msg := gomail.NewMessage(gomail.SetCharset("UTF-8"), gomail.SetEncoding(gomail.QuotedPrintable))
	msg := gomail.NewMessage()
	msg.SetHeader("From", Sender)
	msg.SetHeader("To", *destinations[0])
	msg.SetHeader("Subject", "Teste Gabriel")
	msg.SetHeader("X-SES-CONFIGURATION-SET", "production-hermes-configuration-set")
	//msg.SetHeader("List-Unsubscribe", fmt.Sprintf("https://email-engage.br.originhosting.io/unsubscribe/index.html?email=%s", *destinations[0]))
	//msg.SetHeader("Reply-To", "contato@produtos.cartoes.com.br")
	//msg.SetHeader("X-Report-Abuse-To", "contato@produtos.cartoes.com.br")
	msg.SetBody("text/plain", "OI")
	msg.AddAlternative("text/html", "OI")
	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)
	source := aws.String(Sender)
	/*destinations := []*string{
		aws.String(Recipient),
	}*/
	message := ses.RawMessage{Data: emailRaw.Bytes()}
	input := ses.SendRawEmailInput{Source: source, Destinations: destinations, RawMessage: &message}
	output, err := svc.SendRawEmail(&input)
	if err != nil {
		fmt.Println("Response from SES", err)
	}
	fmt.Println("Response from SES", output)
	return
}

func SendEmail2() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1")},
	)
	// Create an SES session.
	svc := ses.New(sess)
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String("testetsunoda@gmail.com"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String("Ola"),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}
	// Attempt to send the email.
	result, err := svc.SendEmail(input)
	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("Email Sent to address: " + "testetsunoda@gmail.com")
	fmt.Println(result)
}
