package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var t = time.Now().Format("2006-01-02-150405")

func awsSession() (error, *session.Session) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatal("Failed establishing AWS session ", err)
		return err, &session.Session{}
	}
	return nil, sess
}

func terminateStack(session *session.Session) error {
	svc := cloudformation.New(session)

	params := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackname),
	}

	_, err := svc.DeleteStack(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Fatalln("Deleting failed: ", awsErr.Code(), awsErr.Message())
		} else {
			log.Fatalln("Deleting failed ", err)
			return err
		}
	}
	return err
}

func deployStack(t string, session *session.Session) error {
	// file, err := os.Open(tmpl)
	// if err != nil {
	// 	log.Fatal("failed to open template", err)
	// 	return err
	// }

	svc := cloudformation.New(session)

	params := &cloudformation.CreateStackInput{
		StackName:       aws.String(stackname),
		DisableRollback: aws.Bool(true), // no rollback by default
		TemplateBody:    aws.String(t),
	}

	resp, err := svc.CreateStack(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Fatalln("Deploying failed: ", awsErr.Code(), awsErr.Message())
		} else {
			log.Fatalln("Deploying failed ", err)
			return err
		}
	}

	log.Println("Deployment successful: ", resp)
	return nil
}
