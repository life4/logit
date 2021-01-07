// +build h_aws !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	logrus_cloudwatchlogs "github.com/kdar/logrus-cloudwatchlogs"
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

func (config AWSHandler) Parse() (Handler, error) {
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
	h.SetHook(hook)
	return h, nil
}

func init() {
	RegisterParser("aws", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewAWSHandler()
		err := meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		handler, err := fconf.Parse()
		if err != nil {
			return nil, fmt.Errorf("init: %v", err)
		}
		return handler, nil
	})
}
