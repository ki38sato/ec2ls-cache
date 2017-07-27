package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Params is input parameters for aws session
type Params struct {
	profile string
	region  string
}

// make new aws config
func newAwsSession(p Params) (*session.Session, error) {

	config := aws.NewConfig()
	if p.region != "" {
		config.Region = aws.String(p.region)
	}

	options := session.Options{
		Config:                  *config,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}
	if p.profile != "" {
		options.Profile = p.profile
	}
	sess := session.Must(session.NewSessionWithOptions(options))

	return sess, nil
}
