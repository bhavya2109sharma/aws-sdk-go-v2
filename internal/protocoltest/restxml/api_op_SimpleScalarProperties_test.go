// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_SimpleScalarProperties_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *SimpleScalarPropertiesInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes simple scalar properties
		"SimpleScalarProperties": {
			Params: &SimpleScalarPropertiesInput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <stringValue>string</stringValue>
			    <trueBooleanValue>true</trueBooleanValue>
			    <falseBooleanValue>false</falseBooleanValue>
			    <byteValue>1</byteValue>
			    <shortValue>2</shortValue>
			    <integerValue>3</integerValue>
			    <longValue>4</longValue>
			    <floatValue>5.5</floatValue>
			    <DoubleDribble>6.5</DoubleDribble>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Serializes string with escaping
		"SimpleScalarPropertiesWithEscapedCharacter": {
			Params: &SimpleScalarPropertiesInput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("<string>"),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <stringValue>&lt;string&gt;</stringValue>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Serializes string containing white space
		"SimpleScalarPropertiesWithWhiteSpace": {
			Params: &SimpleScalarPropertiesInput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("  string with white    space  "),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <stringValue>  string with white    space  </stringValue>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Serializes string containing exclusively whitespace
		"SimpleScalarPropertiesPureWhiteSpace": {
			Params: &SimpleScalarPropertiesInput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("   "),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <stringValue>   </stringValue>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Supports handling NaN float values.
		"RestXmlSupportsNaNFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.NaN())),
				DoubleValue: ptr.Float64(math.NaN()),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <floatValue>NaN</floatValue>
			    <DoubleDribble>NaN</DoubleDribble>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Supports handling Infinity float values.
		"RestXmlSupportsInfinityFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.Inf(1))),
				DoubleValue: ptr.Float64(math.Inf(1)),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <floatValue>Infinity</floatValue>
			    <DoubleDribble>Infinity</DoubleDribble>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
		// Supports handling -Infinity float values.
		"RestXmlSupportsNegativeInfinityFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.Inf(-1))),
				DoubleValue: ptr.Float64(math.Inf(-1)),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesRequest>
			    <floatValue>-Infinity</floatValue>
			    <DoubleDribble>-Infinity</DoubleDribble>
			</SimpleScalarPropertiesRequest>
			`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actualReq := &http.Request{}
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.SimpleScalarProperties(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_SimpleScalarProperties_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarPropertiesOutput
	}{
		// Serializes simple scalar properties
		"SimpleScalarProperties": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <stringValue>string</stringValue>
			    <trueBooleanValue>true</trueBooleanValue>
			    <falseBooleanValue>false</falseBooleanValue>
			    <byteValue>1</byteValue>
			    <shortValue>2</shortValue>
			    <integerValue>3</integerValue>
			    <longValue>4</longValue>
			    <floatValue>5.5</floatValue>
			    <DoubleDribble>6.5</DoubleDribble>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
		},
		// Serializes string with escaping.
		//
		// This validates the three escape types: literal, decimal and hexadecimal. It
		// also validates that unescaping properly handles the case where unescaping an &
		// produces a newly formed escape sequence (this should not be re-unescaped).
		//
		// Servers may produce different output, this test is designed different unescapes
		// clients must handle
		"SimpleScalarPropertiesComplexEscapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <stringValue>escaped data: &amp;lt;&#xD;&#10;</stringValue>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("escaped data: &lt;\r\n"),
			},
		},
		// Serializes string with escaping
		"SimpleScalarPropertiesWithEscapedCharacter": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <stringValue>&lt;string&gt;</stringValue>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("<string>"),
			},
		},
		// Serializes simple scalar properties with xml preamble, comments and CDATA
		"SimpleScalarPropertiesWithXMLPreamble": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<?xml version = "1.0" encoding = "UTF-8"?>
			<SimpleScalarPropertiesResponse>
			    <![CDATA[characters representing CDATA]]>
			    <stringValue>string</stringValue>
			    <!--xml comment-->
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("string"),
			},
		},
		// Serializes string containing white space
		"SimpleScalarPropertiesWithWhiteSpace": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<?xml version = "1.0" encoding = "UTF-8"?>
			<SimpleScalarPropertiesResponse>
			    <stringValue> string with white    space </stringValue>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String(" string with white    space "),
			},
		},
		// Serializes string containing white space
		"SimpleScalarPropertiesPureWhiteSpace": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<?xml version = "1.0" encoding = "UTF-8"?>
			<SimpleScalarPropertiesResponse>
			    <stringValue>  </stringValue>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:         ptr.String("Foo"),
				StringValue: ptr.String("  "),
			},
		},
		// Supports handling NaN float values.
		"RestXmlSupportsNaNFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <floatValue>NaN</floatValue>
			    <DoubleDribble>NaN</DoubleDribble>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.NaN())),
				DoubleValue: ptr.Float64(math.NaN()),
			},
		},
		// Supports handling Infinity float values.
		"RestXmlSupportsInfinityFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <floatValue>Infinity</floatValue>
			    <DoubleDribble>Infinity</DoubleDribble>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(1))),
				DoubleValue: ptr.Float64(math.Inf(1)),
			},
		},
		// Supports handling -Infinity float values.
		"RestXmlSupportsNegativeInfinityFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesResponse>
			    <floatValue>-Infinity</floatValue>
			    <DoubleDribble>-Infinity</DoubleDribble>
			</SimpleScalarPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(-1))),
				DoubleValue: ptr.Float64(math.Inf(-1)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params SimpleScalarPropertiesInput
			result, err := client.SimpleScalarProperties(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
