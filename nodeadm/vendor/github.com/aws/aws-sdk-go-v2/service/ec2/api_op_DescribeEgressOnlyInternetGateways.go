// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your egress-only internet gateways. The default is to describe all
// your egress-only internet gateways. Alternatively, you can specify specific
// egress-only internet gateway IDs or filter the results to include only the
// egress-only internet gateways that match specific criteria.
func (c *Client) DescribeEgressOnlyInternetGateways(ctx context.Context, params *DescribeEgressOnlyInternetGatewaysInput, optFns ...func(*Options)) (*DescribeEgressOnlyInternetGatewaysOutput, error) {
	if params == nil {
		params = &DescribeEgressOnlyInternetGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEgressOnlyInternetGateways", params, optFns, c.addOperationDescribeEgressOnlyInternetGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEgressOnlyInternetGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEgressOnlyInternetGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The IDs of the egress-only internet gateways.
	EgressOnlyInternetGatewayIds []string

	// The filters.
	//
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEgressOnlyInternetGatewaysOutput struct {

	// Information about the egress-only internet gateways.
	EgressOnlyInternetGateways []types.EgressOnlyInternetGateway

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEgressOnlyInternetGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeEgressOnlyInternetGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeEgressOnlyInternetGateways{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEgressOnlyInternetGateways"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEgressOnlyInternetGateways(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// DescribeEgressOnlyInternetGatewaysPaginatorOptions is the paginator options for
// DescribeEgressOnlyInternetGateways
type DescribeEgressOnlyInternetGatewaysPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEgressOnlyInternetGatewaysPaginator is a paginator for
// DescribeEgressOnlyInternetGateways
type DescribeEgressOnlyInternetGatewaysPaginator struct {
	options   DescribeEgressOnlyInternetGatewaysPaginatorOptions
	client    DescribeEgressOnlyInternetGatewaysAPIClient
	params    *DescribeEgressOnlyInternetGatewaysInput
	nextToken *string
	firstPage bool
}

// NewDescribeEgressOnlyInternetGatewaysPaginator returns a new
// DescribeEgressOnlyInternetGatewaysPaginator
func NewDescribeEgressOnlyInternetGatewaysPaginator(client DescribeEgressOnlyInternetGatewaysAPIClient, params *DescribeEgressOnlyInternetGatewaysInput, optFns ...func(*DescribeEgressOnlyInternetGatewaysPaginatorOptions)) *DescribeEgressOnlyInternetGatewaysPaginator {
	if params == nil {
		params = &DescribeEgressOnlyInternetGatewaysInput{}
	}

	options := DescribeEgressOnlyInternetGatewaysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEgressOnlyInternetGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEgressOnlyInternetGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEgressOnlyInternetGateways page.
func (p *DescribeEgressOnlyInternetGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEgressOnlyInternetGatewaysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeEgressOnlyInternetGateways(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeEgressOnlyInternetGatewaysAPIClient is a client that implements the
// DescribeEgressOnlyInternetGateways operation.
type DescribeEgressOnlyInternetGatewaysAPIClient interface {
	DescribeEgressOnlyInternetGateways(context.Context, *DescribeEgressOnlyInternetGatewaysInput, ...func(*Options)) (*DescribeEgressOnlyInternetGatewaysOutput, error)
}

var _ DescribeEgressOnlyInternetGatewaysAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEgressOnlyInternetGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEgressOnlyInternetGateways",
	}
}
