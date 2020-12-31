package logit

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/kdar/logrus-cloudwatchlogs"
)

type AWSHandler struct {
	BaseHandler
	Region     string
	Group      string
	Stream     string
	MaxRetries int
}

func NewAWSHandler() AWSHandler {
	return AWSHandler{
		BaseHandler: NewBaseHandler(),
		Region:      "us-east-1",
	}
}

func (config AWSHandler) Parse() (*Handler, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:     &config.Region,
			MaxRetries: &config.MaxRetries,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create session: %v", err)
	}

	_, err = sts.New(sess).GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, fmt.Errorf("cannot authorize: %v", err)
	}

	hook, err := logrus_cloudwatchlogs.NewHook(
		config.Group,
		config.Stream,
		sess,
	)
	if err != nil {
		return nil, err
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}
