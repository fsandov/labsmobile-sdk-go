package main

import (
	"context"
	"fmt"

	"github.com/fsandov/labsmobile-sdk-go/pkg/sms"
)

func main() {
	ctx := context.Background()

	apiKey := ""
	username := ""

	client, err := sms.NewClient(apiKey, username)
	if err != nil {
		fmt.Println("error: ", err)
	}

	smsData := sms.SendSMSRequest{
		Message: "Prueba desde SDK con requirements",
		Recipient: []sms.Recipient{
			{
				Msisdn: "56982601777",
			},
		},
	}

	smsResponse, err := client.SendSMS(ctx, smsData)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(smsResponse)

	balance, err := client.GetBalance(ctx)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(balance)

}
