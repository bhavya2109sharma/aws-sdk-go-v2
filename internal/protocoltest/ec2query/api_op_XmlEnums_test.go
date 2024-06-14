// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/aws/smithy-go/middleware"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_XmlEnums_awsEc2queryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlEnumsOutput
	}{
		// Serializes simple scalar properties
		"Ec2XmlEnums": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlEnumsResponse xmlns="https://example.com/">
			    <fooEnum1>Foo</fooEnum1>
			    <fooEnum2>0</fooEnum2>
			    <fooEnum3>1</fooEnum3>
			    <fooEnumList>
			        <member>Foo</member>
			        <member>0</member>
			    </fooEnumList>
			    <fooEnumSet>
			        <member>Foo</member>
			        <member>0</member>
			    </fooEnumSet>
			    <fooEnumMap>
			        <entry>
			            <key>hi</key>
			            <value>Foo</value>
			        </entry>
			        <entry>
			            <key>zero</key>
			            <value>0</value>
			        </entry>
			    </fooEnumMap>
			    <RequestId>requestid</RequestId>
			</XmlEnumsResponse>
			`),
			ExpectResult: &XmlEnumsOutput{
				FooEnum1: types.FooEnum("Foo"),
				FooEnum2: types.FooEnum("0"),
				FooEnum3: types.FooEnum("1"),
				FooEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				FooEnumSet: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				FooEnumMap: map[string]types.FooEnum{
					"hi":   types.FooEnum("Foo"),
					"zero": types.FooEnum("0"),
				},
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
			var params XmlEnumsInput
			result, err := client.XmlEnums(context.Background(), &params)
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
