// various small aws helper funcitons in this file
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func awsSession() (error, *session.Session) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatal("Failed establishing AWS session ", err)
		return err, &session.Session{}
	}
	return nil, sess 
}

// helper functions for finding ids of AWS resources
func subnetFinder(s string) (error, string) { return nil, "" }
func sgFinder(s string) (error, string)     { return nil, "" }
func amiFinder(s string) (error, string)    { return nil, "" }
func vpcFinder(s string) (error, string)    { return nil, "" }
