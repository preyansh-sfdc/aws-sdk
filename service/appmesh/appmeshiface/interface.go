// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appmeshiface provides an interface to enable mocking the AWS App Mesh service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appmeshiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appmesh"
)

// AppMeshAPI provides an interface to enable mocking the
// appmesh.AppMesh service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS App Mesh.
//    func myFunc(svc appmeshiface.AppMeshAPI) bool {
//        // Make svc.CreateGatewayRoute request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appmesh.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppMeshClient struct {
//        appmeshiface.AppMeshAPI
//    }
//    func (m *mockAppMeshClient) CreateGatewayRoute(input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppMeshClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AppMeshAPI interface {
	CreateGatewayRoute(*appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error)
	CreateGatewayRouteWithContext(aws.Context, *appmesh.CreateGatewayRouteInput, ...request.Option) (*appmesh.CreateGatewayRouteOutput, error)
	CreateGatewayRouteRequest(*appmesh.CreateGatewayRouteInput) (*request.Request, *appmesh.CreateGatewayRouteOutput)

	CreateMesh(*appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error)
	CreateMeshWithContext(aws.Context, *appmesh.CreateMeshInput, ...request.Option) (*appmesh.CreateMeshOutput, error)
	CreateMeshRequest(*appmesh.CreateMeshInput) (*request.Request, *appmesh.CreateMeshOutput)

	CreateRoute(*appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error)
	CreateRouteWithContext(aws.Context, *appmesh.CreateRouteInput, ...request.Option) (*appmesh.CreateRouteOutput, error)
	CreateRouteRequest(*appmesh.CreateRouteInput) (*request.Request, *appmesh.CreateRouteOutput)

	CreateVirtualGateway(*appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error)
	CreateVirtualGatewayWithContext(aws.Context, *appmesh.CreateVirtualGatewayInput, ...request.Option) (*appmesh.CreateVirtualGatewayOutput, error)
	CreateVirtualGatewayRequest(*appmesh.CreateVirtualGatewayInput) (*request.Request, *appmesh.CreateVirtualGatewayOutput)

	CreateVirtualNode(*appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error)
	CreateVirtualNodeWithContext(aws.Context, *appmesh.CreateVirtualNodeInput, ...request.Option) (*appmesh.CreateVirtualNodeOutput, error)
	CreateVirtualNodeRequest(*appmesh.CreateVirtualNodeInput) (*request.Request, *appmesh.CreateVirtualNodeOutput)

	CreateVirtualRouter(*appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error)
	CreateVirtualRouterWithContext(aws.Context, *appmesh.CreateVirtualRouterInput, ...request.Option) (*appmesh.CreateVirtualRouterOutput, error)
	CreateVirtualRouterRequest(*appmesh.CreateVirtualRouterInput) (*request.Request, *appmesh.CreateVirtualRouterOutput)

	CreateVirtualService(*appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error)
	CreateVirtualServiceWithContext(aws.Context, *appmesh.CreateVirtualServiceInput, ...request.Option) (*appmesh.CreateVirtualServiceOutput, error)
	CreateVirtualServiceRequest(*appmesh.CreateVirtualServiceInput) (*request.Request, *appmesh.CreateVirtualServiceOutput)

	DeleteGatewayRoute(*appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error)
	DeleteGatewayRouteWithContext(aws.Context, *appmesh.DeleteGatewayRouteInput, ...request.Option) (*appmesh.DeleteGatewayRouteOutput, error)
	DeleteGatewayRouteRequest(*appmesh.DeleteGatewayRouteInput) (*request.Request, *appmesh.DeleteGatewayRouteOutput)

	DeleteMesh(*appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error)
	DeleteMeshWithContext(aws.Context, *appmesh.DeleteMeshInput, ...request.Option) (*appmesh.DeleteMeshOutput, error)
	DeleteMeshRequest(*appmesh.DeleteMeshInput) (*request.Request, *appmesh.DeleteMeshOutput)

	DeleteRoute(*appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error)
	DeleteRouteWithContext(aws.Context, *appmesh.DeleteRouteInput, ...request.Option) (*appmesh.DeleteRouteOutput, error)
	DeleteRouteRequest(*appmesh.DeleteRouteInput) (*request.Request, *appmesh.DeleteRouteOutput)

	DeleteVirtualGateway(*appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error)
	DeleteVirtualGatewayWithContext(aws.Context, *appmesh.DeleteVirtualGatewayInput, ...request.Option) (*appmesh.DeleteVirtualGatewayOutput, error)
	DeleteVirtualGatewayRequest(*appmesh.DeleteVirtualGatewayInput) (*request.Request, *appmesh.DeleteVirtualGatewayOutput)

	DeleteVirtualNode(*appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error)
	DeleteVirtualNodeWithContext(aws.Context, *appmesh.DeleteVirtualNodeInput, ...request.Option) (*appmesh.DeleteVirtualNodeOutput, error)
	DeleteVirtualNodeRequest(*appmesh.DeleteVirtualNodeInput) (*request.Request, *appmesh.DeleteVirtualNodeOutput)

	DeleteVirtualRouter(*appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error)
	DeleteVirtualRouterWithContext(aws.Context, *appmesh.DeleteVirtualRouterInput, ...request.Option) (*appmesh.DeleteVirtualRouterOutput, error)
	DeleteVirtualRouterRequest(*appmesh.DeleteVirtualRouterInput) (*request.Request, *appmesh.DeleteVirtualRouterOutput)

	DeleteVirtualService(*appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error)
	DeleteVirtualServiceWithContext(aws.Context, *appmesh.DeleteVirtualServiceInput, ...request.Option) (*appmesh.DeleteVirtualServiceOutput, error)
	DeleteVirtualServiceRequest(*appmesh.DeleteVirtualServiceInput) (*request.Request, *appmesh.DeleteVirtualServiceOutput)

	DescribeGatewayRoute(*appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error)
	DescribeGatewayRouteWithContext(aws.Context, *appmesh.DescribeGatewayRouteInput, ...request.Option) (*appmesh.DescribeGatewayRouteOutput, error)
	DescribeGatewayRouteRequest(*appmesh.DescribeGatewayRouteInput) (*request.Request, *appmesh.DescribeGatewayRouteOutput)

	DescribeMesh(*appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error)
	DescribeMeshWithContext(aws.Context, *appmesh.DescribeMeshInput, ...request.Option) (*appmesh.DescribeMeshOutput, error)
	DescribeMeshRequest(*appmesh.DescribeMeshInput) (*request.Request, *appmesh.DescribeMeshOutput)

	DescribeRoute(*appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error)
	DescribeRouteWithContext(aws.Context, *appmesh.DescribeRouteInput, ...request.Option) (*appmesh.DescribeRouteOutput, error)
	DescribeRouteRequest(*appmesh.DescribeRouteInput) (*request.Request, *appmesh.DescribeRouteOutput)

	DescribeVirtualGateway(*appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error)
	DescribeVirtualGatewayWithContext(aws.Context, *appmesh.DescribeVirtualGatewayInput, ...request.Option) (*appmesh.DescribeVirtualGatewayOutput, error)
	DescribeVirtualGatewayRequest(*appmesh.DescribeVirtualGatewayInput) (*request.Request, *appmesh.DescribeVirtualGatewayOutput)

	DescribeVirtualNode(*appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error)
	DescribeVirtualNodeWithContext(aws.Context, *appmesh.DescribeVirtualNodeInput, ...request.Option) (*appmesh.DescribeVirtualNodeOutput, error)
	DescribeVirtualNodeRequest(*appmesh.DescribeVirtualNodeInput) (*request.Request, *appmesh.DescribeVirtualNodeOutput)

	DescribeVirtualRouter(*appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error)
	DescribeVirtualRouterWithContext(aws.Context, *appmesh.DescribeVirtualRouterInput, ...request.Option) (*appmesh.DescribeVirtualRouterOutput, error)
	DescribeVirtualRouterRequest(*appmesh.DescribeVirtualRouterInput) (*request.Request, *appmesh.DescribeVirtualRouterOutput)

	DescribeVirtualService(*appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error)
	DescribeVirtualServiceWithContext(aws.Context, *appmesh.DescribeVirtualServiceInput, ...request.Option) (*appmesh.DescribeVirtualServiceOutput, error)
	DescribeVirtualServiceRequest(*appmesh.DescribeVirtualServiceInput) (*request.Request, *appmesh.DescribeVirtualServiceOutput)

	ListGatewayRoutes(*appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error)
	ListGatewayRoutesWithContext(aws.Context, *appmesh.ListGatewayRoutesInput, ...request.Option) (*appmesh.ListGatewayRoutesOutput, error)
	ListGatewayRoutesRequest(*appmesh.ListGatewayRoutesInput) (*request.Request, *appmesh.ListGatewayRoutesOutput)

	ListGatewayRoutesPages(*appmesh.ListGatewayRoutesInput, func(*appmesh.ListGatewayRoutesOutput, bool) bool) error
	ListGatewayRoutesPagesWithContext(aws.Context, *appmesh.ListGatewayRoutesInput, func(*appmesh.ListGatewayRoutesOutput, bool) bool, ...request.Option) error

	ListMeshes(*appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error)
	ListMeshesWithContext(aws.Context, *appmesh.ListMeshesInput, ...request.Option) (*appmesh.ListMeshesOutput, error)
	ListMeshesRequest(*appmesh.ListMeshesInput) (*request.Request, *appmesh.ListMeshesOutput)

	ListMeshesPages(*appmesh.ListMeshesInput, func(*appmesh.ListMeshesOutput, bool) bool) error
	ListMeshesPagesWithContext(aws.Context, *appmesh.ListMeshesInput, func(*appmesh.ListMeshesOutput, bool) bool, ...request.Option) error

	ListRoutes(*appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error)
	ListRoutesWithContext(aws.Context, *appmesh.ListRoutesInput, ...request.Option) (*appmesh.ListRoutesOutput, error)
	ListRoutesRequest(*appmesh.ListRoutesInput) (*request.Request, *appmesh.ListRoutesOutput)

	ListRoutesPages(*appmesh.ListRoutesInput, func(*appmesh.ListRoutesOutput, bool) bool) error
	ListRoutesPagesWithContext(aws.Context, *appmesh.ListRoutesInput, func(*appmesh.ListRoutesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *appmesh.ListTagsForResourceInput, ...request.Option) (*appmesh.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*appmesh.ListTagsForResourceInput) (*request.Request, *appmesh.ListTagsForResourceOutput)

	ListTagsForResourcePages(*appmesh.ListTagsForResourceInput, func(*appmesh.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *appmesh.ListTagsForResourceInput, func(*appmesh.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListVirtualGateways(*appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error)
	ListVirtualGatewaysWithContext(aws.Context, *appmesh.ListVirtualGatewaysInput, ...request.Option) (*appmesh.ListVirtualGatewaysOutput, error)
	ListVirtualGatewaysRequest(*appmesh.ListVirtualGatewaysInput) (*request.Request, *appmesh.ListVirtualGatewaysOutput)

	ListVirtualGatewaysPages(*appmesh.ListVirtualGatewaysInput, func(*appmesh.ListVirtualGatewaysOutput, bool) bool) error
	ListVirtualGatewaysPagesWithContext(aws.Context, *appmesh.ListVirtualGatewaysInput, func(*appmesh.ListVirtualGatewaysOutput, bool) bool, ...request.Option) error

	ListVirtualNodes(*appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error)
	ListVirtualNodesWithContext(aws.Context, *appmesh.ListVirtualNodesInput, ...request.Option) (*appmesh.ListVirtualNodesOutput, error)
	ListVirtualNodesRequest(*appmesh.ListVirtualNodesInput) (*request.Request, *appmesh.ListVirtualNodesOutput)

	ListVirtualNodesPages(*appmesh.ListVirtualNodesInput, func(*appmesh.ListVirtualNodesOutput, bool) bool) error
	ListVirtualNodesPagesWithContext(aws.Context, *appmesh.ListVirtualNodesInput, func(*appmesh.ListVirtualNodesOutput, bool) bool, ...request.Option) error

	ListVirtualRouters(*appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error)
	ListVirtualRoutersWithContext(aws.Context, *appmesh.ListVirtualRoutersInput, ...request.Option) (*appmesh.ListVirtualRoutersOutput, error)
	ListVirtualRoutersRequest(*appmesh.ListVirtualRoutersInput) (*request.Request, *appmesh.ListVirtualRoutersOutput)

	ListVirtualRoutersPages(*appmesh.ListVirtualRoutersInput, func(*appmesh.ListVirtualRoutersOutput, bool) bool) error
	ListVirtualRoutersPagesWithContext(aws.Context, *appmesh.ListVirtualRoutersInput, func(*appmesh.ListVirtualRoutersOutput, bool) bool, ...request.Option) error

	ListVirtualServices(*appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error)
	ListVirtualServicesWithContext(aws.Context, *appmesh.ListVirtualServicesInput, ...request.Option) (*appmesh.ListVirtualServicesOutput, error)
	ListVirtualServicesRequest(*appmesh.ListVirtualServicesInput) (*request.Request, *appmesh.ListVirtualServicesOutput)

	ListVirtualServicesPages(*appmesh.ListVirtualServicesInput, func(*appmesh.ListVirtualServicesOutput, bool) bool) error
	ListVirtualServicesPagesWithContext(aws.Context, *appmesh.ListVirtualServicesInput, func(*appmesh.ListVirtualServicesOutput, bool) bool, ...request.Option) error

	TagResource(*appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *appmesh.TagResourceInput, ...request.Option) (*appmesh.TagResourceOutput, error)
	TagResourceRequest(*appmesh.TagResourceInput) (*request.Request, *appmesh.TagResourceOutput)

	UntagResource(*appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *appmesh.UntagResourceInput, ...request.Option) (*appmesh.UntagResourceOutput, error)
	UntagResourceRequest(*appmesh.UntagResourceInput) (*request.Request, *appmesh.UntagResourceOutput)

	UpdateGatewayRoute(*appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error)
	UpdateGatewayRouteWithContext(aws.Context, *appmesh.UpdateGatewayRouteInput, ...request.Option) (*appmesh.UpdateGatewayRouteOutput, error)
	UpdateGatewayRouteRequest(*appmesh.UpdateGatewayRouteInput) (*request.Request, *appmesh.UpdateGatewayRouteOutput)

	UpdateMesh(*appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error)
	UpdateMeshWithContext(aws.Context, *appmesh.UpdateMeshInput, ...request.Option) (*appmesh.UpdateMeshOutput, error)
	UpdateMeshRequest(*appmesh.UpdateMeshInput) (*request.Request, *appmesh.UpdateMeshOutput)

	UpdateRoute(*appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error)
	UpdateRouteWithContext(aws.Context, *appmesh.UpdateRouteInput, ...request.Option) (*appmesh.UpdateRouteOutput, error)
	UpdateRouteRequest(*appmesh.UpdateRouteInput) (*request.Request, *appmesh.UpdateRouteOutput)

	UpdateVirtualGateway(*appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error)
	UpdateVirtualGatewayWithContext(aws.Context, *appmesh.UpdateVirtualGatewayInput, ...request.Option) (*appmesh.UpdateVirtualGatewayOutput, error)
	UpdateVirtualGatewayRequest(*appmesh.UpdateVirtualGatewayInput) (*request.Request, *appmesh.UpdateVirtualGatewayOutput)

	UpdateVirtualNode(*appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error)
	UpdateVirtualNodeWithContext(aws.Context, *appmesh.UpdateVirtualNodeInput, ...request.Option) (*appmesh.UpdateVirtualNodeOutput, error)
	UpdateVirtualNodeRequest(*appmesh.UpdateVirtualNodeInput) (*request.Request, *appmesh.UpdateVirtualNodeOutput)

	UpdateVirtualRouter(*appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error)
	UpdateVirtualRouterWithContext(aws.Context, *appmesh.UpdateVirtualRouterInput, ...request.Option) (*appmesh.UpdateVirtualRouterOutput, error)
	UpdateVirtualRouterRequest(*appmesh.UpdateVirtualRouterInput) (*request.Request, *appmesh.UpdateVirtualRouterOutput)

	UpdateVirtualService(*appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error)
	UpdateVirtualServiceWithContext(aws.Context, *appmesh.UpdateVirtualServiceInput, ...request.Option) (*appmesh.UpdateVirtualServiceOutput, error)
	UpdateVirtualServiceRequest(*appmesh.UpdateVirtualServiceInput) (*request.Request, *appmesh.UpdateVirtualServiceOutput)
}

var _ AppMeshAPI = (*appmesh.AppMesh)(nil)
