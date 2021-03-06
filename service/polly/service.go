// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Polly provides the API operation methods for making requests to
// Amazon Polly. See this package's package overview docs
// for details on the service.
//
// Polly methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Polly struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Polly)

// Used for custom request initialization logic
var initRequest func(*Polly, *aws.Request)

// Service information constants
const (
	ServiceName = "polly"     // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Polly client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Polly client from just a config.
//     svc := polly.New(myConfig)
//
//     // Create a Polly client with additional configuration
//     svc := polly.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *Polly {
	var signingName string
	signingRegion := config.Region

	svc := &Polly{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-06-10",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Polly operation and runs any
// custom request initialization.
func (c *Polly) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
