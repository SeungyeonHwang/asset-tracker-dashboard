package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "hwang_personal",
		SharedConfigState: session.SharedConfigEnable,
	}))

	ssmSvc := ssm.New(sess)
	paramName := "/binance/apiKey"

	param, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		fmt.Println("Error getting parameter:", err)
		return
	}

	fmt.Println("Binance API Key:", *param.Parameter.Value)
}
