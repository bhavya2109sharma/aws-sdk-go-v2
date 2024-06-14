// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_XmlTimestamps_awsEc2queryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlTimestampsOutput
	}{
		// Tests how normal timestamps are serialized
		"Ec2XmlTimestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <normal>2014-04-29T18:30:38Z</normal>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"Ec2XmlTimestampsWithDateTimeFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <dateTime>2014-04-29T18:30:38Z</dateTime>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time on the target shape works like
		// normal timestamps
		"Ec2XmlTimestampsWithDateTimeOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <dateTimeOnTarget>2014-04-29T18:30:38Z</dateTimeOnTarget>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				DateTimeOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"Ec2XmlTimestampsWithEpochSecondsFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <epochSeconds>1398796238</epochSeconds>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds on the target shape works
		"Ec2XmlTimestampsWithEpochSecondsOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <epochSecondsOnTarget>1398796238</epochSecondsOnTarget>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				EpochSecondsOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date works
		"Ec2XmlTimestampsWithHttpDateFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <httpDate>Tue, 29 Apr 2014 18:30:38 GMT</httpDate>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date on the target shape works
		"Ec2XmlTimestampsWithHttpDateOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsResponse xmlns="https://example.com/">
			    <httpDateOnTarget>Tue, 29 Apr 2014 18:30:38 GMT</httpDateOnTarget>
			    <RequestId>requestid</RequestId>
			</XmlTimestampsResponse>
			`),
			ExpectResult: &XmlTimestampsOutput{
				HttpDateOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
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
			var params XmlTimestampsInput
			result, err := client.XmlTimestamps(context.Background(), &params)
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
