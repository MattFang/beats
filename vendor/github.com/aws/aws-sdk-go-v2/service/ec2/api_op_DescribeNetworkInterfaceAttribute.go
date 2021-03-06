// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeNetworkInterfaceAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfaceAttributeRequest
type DescribeNetworkInterfaceAttributeInput struct {
	_ struct{} `type:"structure"`

	// The attribute of the network interface. This parameter is required.
	Attribute NetworkInterfaceAttribute `locationName:"attribute" type:"string" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNetworkInterfaceAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkInterfaceAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNetworkInterfaceAttributeInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeNetworkInterfaceAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfaceAttributeResult
type DescribeNetworkInterfaceAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The attachment (if any) of the network interface.
	Attachment *NetworkInterfaceAttachment `locationName:"attachment" type:"structure"`

	// The description of the network interface.
	Description *AttributeValue `locationName:"description" type:"structure"`

	// The security groups associated with the network interface.
	Groups []GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`

	// The ID of the network interface.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`

	// Indicates whether source/destination checking is enabled.
	SourceDestCheck *AttributeBooleanValue `locationName:"sourceDestCheck" type:"structure"`
}

// String returns the string representation
func (s DescribeNetworkInterfaceAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNetworkInterfaceAttribute = "DescribeNetworkInterfaceAttribute"

// DescribeNetworkInterfaceAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes a network interface attribute. You can specify only one attribute
// at a time.
//
//    // Example sending a request using DescribeNetworkInterfaceAttributeRequest.
//    req := client.DescribeNetworkInterfaceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfaceAttribute
func (c *Client) DescribeNetworkInterfaceAttributeRequest(input *DescribeNetworkInterfaceAttributeInput) DescribeNetworkInterfaceAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeNetworkInterfaceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNetworkInterfaceAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeNetworkInterfaceAttributeOutput{})
	return DescribeNetworkInterfaceAttributeRequest{Request: req, Input: input, Copy: c.DescribeNetworkInterfaceAttributeRequest}
}

// DescribeNetworkInterfaceAttributeRequest is the request type for the
// DescribeNetworkInterfaceAttribute API operation.
type DescribeNetworkInterfaceAttributeRequest struct {
	*aws.Request
	Input *DescribeNetworkInterfaceAttributeInput
	Copy  func(*DescribeNetworkInterfaceAttributeInput) DescribeNetworkInterfaceAttributeRequest
}

// Send marshals and sends the DescribeNetworkInterfaceAttribute API request.
func (r DescribeNetworkInterfaceAttributeRequest) Send(ctx context.Context) (*DescribeNetworkInterfaceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNetworkInterfaceAttributeResponse{
		DescribeNetworkInterfaceAttributeOutput: r.Request.Data.(*DescribeNetworkInterfaceAttributeOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNetworkInterfaceAttributeResponse is the response type for the
// DescribeNetworkInterfaceAttribute API operation.
type DescribeNetworkInterfaceAttributeResponse struct {
	*DescribeNetworkInterfaceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNetworkInterfaceAttribute request.
func (r *DescribeNetworkInterfaceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
