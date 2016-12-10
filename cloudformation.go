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

func deployStack(t string, session *session.Session) error {
	svc := cloudformation.New(session)

	createParams := &cloudformation.CreateStackInput{
		StackName:       aws.String(stackname),
		DisableRollback: aws.Bool(true), // no rollback by default
		TemplateBody:    aws.String(t),
	}

	out, err := svc.CreateStack(createParams)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Fatal("Deploying failed:", awsErr.Code(), awsErr.Message())
		} else {
			log.Fatal("Deploying failed", err)
			return err
		}
	}

	describeStacksInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}
	if err := svc.WaitUntilStackCreateComplete(describeStacksInput); err != nil {
		// FIXME this works in so far that we wait until the stack is
		// completed and capture errors, but it doesn't really tail
		// cloudroamtion events.
		log.Fatal(err)
	}

	log.Println("Deployment successful:", out)
	return nil
}

func terminateStack(session *session.Session) error {
	svc := cloudformation.New(session)

	params := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackname),
	}

	out, err := svc.DeleteStack(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Fatalln("Deleting failed: ", awsErr.Code(), awsErr.Message())
		} else {
			log.Fatalln("Deleting failed ", err)
			return err
		}
	}

	describeStacksInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	if err := svc.WaitUntilStackDeleteComplete(describeStacksInput); err != nil {
		// FIXME this works in so far that we wait until the stack is
		// completed and capture errors, but it doesn't really tail
		// cloudroamtion events.
		log.Fatal(err)
	}

	log.Println("Deletion successful:", out)
	return nil
}
