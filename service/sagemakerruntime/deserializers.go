// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

func deserializeS3Expires(v string) (*time.Time, error) {
	t, err := smithytime.ParseHTTPDate(v)
	if err != nil {
		return nil, nil
	}
	return &t, nil
}

type awsRestjson1_deserializeOpInvokeEndpoint struct {
}

func (*awsRestjson1_deserializeOpInvokeEndpoint) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpInvokeEndpoint) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorInvokeEndpoint(response, &metadata)
	}
	output := &InvokeEndpointOutput{}
	out.Result = output

	err = awsRestjson1_deserializeOpHttpBindingsInvokeEndpointOutput(output, response)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
	}

	err = awsRestjson1_deserializeOpDocumentInvokeEndpointOutput(output, response.Body, response.ContentLength)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorInvokeEndpoint(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalDependencyException", errorCode):
		return awsRestjson1_deserializeErrorInternalDependencyException(response, errorBody)

	case strings.EqualFold("InternalFailure", errorCode):
		return awsRestjson1_deserializeErrorInternalFailure(response, errorBody)

	case strings.EqualFold("ModelError", errorCode):
		return awsRestjson1_deserializeErrorModelError(response, errorBody)

	case strings.EqualFold("ModelNotReadyException", errorCode):
		return awsRestjson1_deserializeErrorModelNotReadyException(response, errorBody)

	case strings.EqualFold("ServiceUnavailable", errorCode):
		return awsRestjson1_deserializeErrorServiceUnavailable(response, errorBody)

	case strings.EqualFold("ValidationError", errorCode):
		return awsRestjson1_deserializeErrorValidationError(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpHttpBindingsInvokeEndpointOutput(v *InvokeEndpointOutput, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-Closed-Session-Id"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ClosedSessionId = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Content-Type"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ContentType = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-Custom-Attributes"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.CustomAttributes = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("x-Amzn-Invoked-Production-Variant"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.InvokedProductionVariant = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-New-Session-Id"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.NewSessionId = ptr.String(headerValues[0])
	}

	return nil
}
func awsRestjson1_deserializeOpDocumentInvokeEndpointOutput(v *InvokeEndpointOutput, body io.ReadCloser, contentLength int64) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization of nil %T", v)
	}

	var buf bytes.Buffer
	if contentLength > 0 {
		buf.Grow(int(contentLength))
	} else {
		buf.Grow(512)
	}

	_, err := buf.ReadFrom(body)
	if err != nil {
		return err
	}
	if buf.Len() > 0 {
		v.Body = buf.Bytes()
	}
	return nil
}

type awsRestjson1_deserializeOpInvokeEndpointAsync struct {
}

func (*awsRestjson1_deserializeOpInvokeEndpointAsync) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpInvokeEndpointAsync) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorInvokeEndpointAsync(response, &metadata)
	}
	output := &InvokeEndpointAsyncOutput{}
	out.Result = output

	err = awsRestjson1_deserializeOpHttpBindingsInvokeEndpointAsyncOutput(output, response)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentInvokeEndpointAsyncOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorInvokeEndpointAsync(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalFailure", errorCode):
		return awsRestjson1_deserializeErrorInternalFailure(response, errorBody)

	case strings.EqualFold("ServiceUnavailable", errorCode):
		return awsRestjson1_deserializeErrorServiceUnavailable(response, errorBody)

	case strings.EqualFold("ValidationError", errorCode):
		return awsRestjson1_deserializeErrorValidationError(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpHttpBindingsInvokeEndpointAsyncOutput(v *InvokeEndpointAsyncOutput, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-FailureLocation"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.FailureLocation = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-OutputLocation"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.OutputLocation = ptr.String(headerValues[0])
	}

	return nil
}
func awsRestjson1_deserializeOpDocumentInvokeEndpointAsyncOutput(v **InvokeEndpointAsyncOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *InvokeEndpointAsyncOutput
	if *v == nil {
		sv = &InvokeEndpointAsyncOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "InferenceId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Header to be of type string, got %T instead", value)
				}
				sv.InferenceId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpInvokeEndpointWithResponseStream struct {
}

func (*awsRestjson1_deserializeOpInvokeEndpointWithResponseStream) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpInvokeEndpointWithResponseStream) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorInvokeEndpointWithResponseStream(response, &metadata)
	}
	output := &InvokeEndpointWithResponseStreamOutput{}
	out.Result = output

	err = awsRestjson1_deserializeOpHttpBindingsInvokeEndpointWithResponseStreamOutput(output, response)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorInvokeEndpointWithResponseStream(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalFailure", errorCode):
		return awsRestjson1_deserializeErrorInternalFailure(response, errorBody)

	case strings.EqualFold("InternalStreamFailure", errorCode):
		return awsRestjson1_deserializeErrorInternalStreamFailure(response, errorBody)

	case strings.EqualFold("ModelError", errorCode):
		return awsRestjson1_deserializeErrorModelError(response, errorBody)

	case strings.EqualFold("ModelStreamError", errorCode):
		return awsRestjson1_deserializeErrorModelStreamError(response, errorBody)

	case strings.EqualFold("ServiceUnavailable", errorCode):
		return awsRestjson1_deserializeErrorServiceUnavailable(response, errorBody)

	case strings.EqualFold("ValidationError", errorCode):
		return awsRestjson1_deserializeErrorValidationError(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpHttpBindingsInvokeEndpointWithResponseStreamOutput(v *InvokeEndpointWithResponseStreamOutput, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-Content-Type"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ContentType = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("X-Amzn-SageMaker-Custom-Attributes"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.CustomAttributes = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("x-Amzn-Invoked-Production-Variant"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.InvokedProductionVariant = ptr.String(headerValues[0])
	}

	return nil
}
func awsRestjson1_deserializeEventStreamResponseStream(v *types.ResponseStream, msg *eventstream.Message) error {
	if v == nil {
		return fmt.Errorf("unexpected serialization of nil %T", v)
	}

	eventType := msg.Headers.Get(eventstreamapi.EventTypeHeader)
	if eventType == nil {
		return fmt.Errorf("%s event header not present", eventstreamapi.EventTypeHeader)
	}

	switch {
	case strings.EqualFold("PayloadPart", eventType.String()):
		vv := &types.ResponseStreamMemberPayloadPart{}
		if err := awsRestjson1_deserializeEventMessagePayloadPart(&vv.Value, msg); err != nil {
			return err
		}
		*v = vv
		return nil

	default:
		buffer := bytes.NewBuffer(nil)
		eventstream.NewEncoder().Encode(buffer, *msg)
		*v = &types.UnknownUnionMember{
			Tag:   eventType.String(),
			Value: buffer.Bytes(),
		}
		return nil

	}
}

func awsRestjson1_deserializeEventStreamExceptionResponseStream(msg *eventstream.Message) error {
	exceptionType := msg.Headers.Get(eventstreamapi.ExceptionTypeHeader)
	if exceptionType == nil {
		return fmt.Errorf("%s event header not present", eventstreamapi.ExceptionTypeHeader)
	}

	switch {
	case strings.EqualFold("InternalStreamFailure", exceptionType.String()):
		return awsRestjson1_deserializeEventMessageExceptionInternalStreamFailure(msg)

	case strings.EqualFold("ModelStreamError", exceptionType.String()):
		return awsRestjson1_deserializeEventMessageExceptionModelStreamError(msg)

	default:
		br := bytes.NewReader(msg.Payload)
		var buff [1024]byte
		ringBuffer := smithyio.NewRingBuffer(buff[:])

		body := io.TeeReader(br, ringBuffer)
		decoder := json.NewDecoder(body)
		decoder.UseNumber()
		code, message, err := restjson.GetErrorInfo(decoder)
		if err != nil {
			return err
		}
		errorCode := "UnknownError"
		errorMessage := errorCode
		if ev := exceptionType.String(); len(ev) > 0 {
			errorCode = ev
		} else if ev := code; len(ev) > 0 {
			errorCode = ev
		}
		if ev := message; len(ev) > 0 {
			errorMessage = ev
		}
		return &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	}
}

func awsRestjson1_deserializeEventMessagePayloadPart(v *types.PayloadPart, msg *eventstream.Message) error {
	if v == nil {
		return fmt.Errorf("unexpected serialization of nil %T", v)
	}

	if msg.Payload != nil {
		bsv := make([]byte, len(msg.Payload))
		copy(bsv, msg.Payload)

		v.Bytes = bsv
	}
	return nil
}

func awsRestjson1_deserializeEventMessageExceptionModelStreamError(msg *eventstream.Message) error {
	br := bytes.NewReader(msg.Payload)
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(br, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	v := &types.ModelStreamError{}
	if err := awsRestjson1_deserializeDocumentModelStreamError(&v, shape); err != nil {
		if err != nil {
			var snapshot bytes.Buffer
			io.Copy(&snapshot, ringBuffer)
			err = &smithy.DeserializationError{
				Err:      fmt.Errorf("failed to decode response body, %w", err),
				Snapshot: snapshot.Bytes(),
			}
			return err
		}

	}
	return v
}

func awsRestjson1_deserializeEventMessageExceptionInternalStreamFailure(msg *eventstream.Message) error {
	br := bytes.NewReader(msg.Payload)
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(br, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	v := &types.InternalStreamFailure{}
	if err := awsRestjson1_deserializeDocumentInternalStreamFailure(&v, shape); err != nil {
		if err != nil {
			var snapshot bytes.Buffer
			io.Copy(&snapshot, ringBuffer)
			err = &smithy.DeserializationError{
				Err:      fmt.Errorf("failed to decode response body, %w", err),
				Snapshot: snapshot.Bytes(),
			}
			return err
		}

	}
	return v
}

func awsRestjson1_deserializeDocumentInternalStreamFailure(v **types.InternalStreamFailure, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalStreamFailure
	if *v == nil {
		sv = &types.InternalStreamFailure{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentModelStreamError(v **types.ModelStreamError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ModelStreamError
	if *v == nil {
		sv = &types.ModelStreamError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ErrorCode":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorCode to be of type string, got %T instead", value)
				}
				sv.ErrorCode_ = ptr.String(jtv)
			}

		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorInternalDependencyException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalDependencyException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalDependencyException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalFailure(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalFailure{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalFailure(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalStreamFailure(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalStreamFailure{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalStreamFailure(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorModelError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ModelError{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentModelError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorModelNotReadyException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ModelNotReadyException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentModelNotReadyException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorModelStreamError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ModelStreamError{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentModelStreamError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorServiceUnavailable(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ServiceUnavailable{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentServiceUnavailable(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorValidationError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ValidationError{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentValidationError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentInternalDependencyException(v **types.InternalDependencyException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalDependencyException
	if *v == nil {
		sv = &types.InternalDependencyException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInternalFailure(v **types.InternalFailure, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalFailure
	if *v == nil {
		sv = &types.InternalFailure{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentModelError(v **types.ModelError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ModelError
	if *v == nil {
		sv = &types.ModelError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "LogStreamArn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected LogStreamArn to be of type string, got %T instead", value)
				}
				sv.LogStreamArn = ptr.String(jtv)
			}

		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		case "OriginalMessage":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.OriginalMessage = ptr.String(jtv)
			}

		case "OriginalStatusCode":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected StatusCode to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.OriginalStatusCode = ptr.Int32(int32(i64))
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentModelNotReadyException(v **types.ModelNotReadyException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ModelNotReadyException
	if *v == nil {
		sv = &types.ModelNotReadyException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentServiceUnavailable(v **types.ServiceUnavailable, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ServiceUnavailable
	if *v == nil {
		sv = &types.ServiceUnavailable{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentValidationError(v **types.ValidationError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationError
	if *v == nil {
		sv = &types.ValidationError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
