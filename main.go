package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

const (
	START = "start"
	STOP  = "stop"
)

type Request struct {
	Type   string `json:"type"`
	DbName string `json:"dbName"`
}

func HandleRequest(ctx context.Context, request Request) {
	mySession := session.Must(session.NewSession())
	RDS := rds.New(mySession)
	if request.Type == STOP {
		dbConfig := rds.StopDBInstanceInput{DBInstanceIdentifier: &request.DbName}
		if output, err := RDS.StopDBInstance(&dbConfig); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(output.GoString())
		}
		return
	}
	dbConfig := rds.StartDBInstanceInput{DBInstanceIdentifier: &request.DbName}
	if output, err := RDS.StartDBInstance(&dbConfig); err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(output.GoString())
	}
}

func main() {
	lambda.Start(HandleRequest)
}
