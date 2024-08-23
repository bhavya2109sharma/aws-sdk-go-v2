// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// For region aws-global with FIPS disabled and DualStack disabled
func TestEndpointCase0(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-global"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-east-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-global with FIPS enabled and DualStack disabled
func TestEndpointCase1(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-global"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-east-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase2(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase3(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-east-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase4(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase5(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-east-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-cn-global with FIPS disabled and DualStack disabled
func TestEndpointCase6(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-cn-global"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.cn-northwest-1.amazonaws.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "cn-northwest-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-north-1 with FIPS enabled and DualStack enabled
func TestEndpointCase7(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-north-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.cn-north-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-north-1 with FIPS enabled and DualStack disabled
func TestEndpointCase8(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-north-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.cn-north-1.amazonaws.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-north-1 with FIPS disabled and DualStack enabled
func TestEndpointCase9(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-north-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.cn-north-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-north-1 with FIPS disabled and DualStack disabled
func TestEndpointCase10(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("cn-north-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.cn-northwest-1.amazonaws.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "cn-northwest-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-us-gov-global with FIPS disabled and DualStack disabled
func TestEndpointCase11(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-us-gov-global"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-gov-west-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-us-gov-global with FIPS enabled and DualStack disabled
func TestEndpointCase12(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-us-gov-global"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-gov-west-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase13(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-gov-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase14(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-gov-west-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase15(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-gov-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase16(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-gov-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-gov-west-1.amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-iso-global with FIPS disabled and DualStack disabled
func TestEndpointCase17(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-iso-global"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-iso-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase18(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-iso-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase19(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase20(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-iso-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase21(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-iso-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-iso-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region aws-iso-b-global with FIPS disabled and DualStack disabled
func TestEndpointCase22(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("aws-iso-b-global"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-isob-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase23(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isob-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase24(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations-fips.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase25(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-isob-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase26(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isob-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-isob-east-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isof-south-1 with FIPS disabled and DualStack disabled
func TestEndpointCase27(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-isof-south-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://organizations.us-isof-south-1.csp.hci.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningName(&sp, "organizations")
						smithyhttp.SetSigV4ASigningName(&sp, "organizations")

						smithyhttp.SetSigV4SigningRegion(&sp, "us-isof-south-1")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For custom endpoint with region set and fips disabled and dualstack disabled
func TestEndpointCase28(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
		Endpoint:     ptr.String("https://example.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://example.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For custom endpoint with region not set and fips disabled and dualstack disabled
func TestEndpointCase29(t *testing.T) {
	var params = EndpointParameters{
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(false),
		Endpoint:     ptr.String("https://example.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://example.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For custom endpoint with fips enabled and dualstack disabled
func TestEndpointCase30(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(true),
		UseDualStack: ptr.Bool(false),
		Endpoint:     ptr.String("https://example.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For custom endpoint with fips disabled and dualstack enabled
func TestEndpointCase31(t *testing.T) {
	var params = EndpointParameters{
		Region:       ptr.String("us-east-1"),
		UseFIPS:      ptr.Bool(false),
		UseDualStack: ptr.Bool(true),
		Endpoint:     ptr.String("https://example.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: Dualstack and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// Missing region
func TestEndpointCase32(t *testing.T) {
	var params = EndpointParameters{}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: Missing Region", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}
