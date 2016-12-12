package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func deployStack(t string, session *session.Session) error {
	svc := cloudformation.New(session)

	createParams := &cloudformation.CreateStackInput{
		StackName:       aws.String(stackname),
		DisableRollback: aws.Bool(true), // no rollback by default
		TemplateBody:    aws.String(t),
	}

	updateParams := &cloudformation.UpdateStackInput{
		StackName:    aws.String(stackname),
		TemplateBody: aws.String(t),
	}

	if stackExists(session) {
		log.Println("Stack exists, updating...")

		out, err := svc.UpdateStack(updateParams)

		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				log.Fatal("Updating stack failed: ", awsErr.Code(), awsErr.Message())
			} else {
				log.Fatal("Updating stack failed ", err)
				return err
			}
		}

		describeStacksInput := &cloudformation.DescribeStacksInput{
			StackName: aws.String(stackname),
		}
		if err := svc.WaitUntilStackUpdateComplete(describeStacksInput); err != nil {
			// FIXME this works in so far that we wait until the stack is
			// completed and capture errors, but it doesn't really tail
			// cloudroamtion events.
			log.Fatal(err)
		}

		log.Println("Stack update successful:", out)

	} else {

		out, err := svc.CreateStack(createParams)

		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				log.Fatal("Deploying failed: ", awsErr.Code(), awsErr.Message())
			} else {
				log.Fatal("Deploying failed ", err)
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
	}

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
			log.Fatal("Deleting failed: ", awsErr.Code(), awsErr.Message())
		} else {
			log.Fatal("Deleting failed ", err)
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

	log.Print("Deletion successful: ", out)
	return nil
}

func stackExists(session *session.Session) bool {
	svc := cloudformation.New(session)

	describeStacksInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	_, err := svc.DescribeStacks(describeStacksInput)

	if err == nil {
		return true
	}

	return false
}
