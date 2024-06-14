// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/document"
	internaldocument "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/internal/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"math"
)

type awsRestjson1_serializeOpInvokeAgent struct {
}

func (*awsRestjson1_serializeOpInvokeAgent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpInvokeAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*InvokeAgentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/agents/{agentId}/agentAliases/{agentAliasId}/sessions/{sessionId}/text")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsInvokeAgentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentInvokeAgentInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsInvokeAgentInput(v *InvokeAgentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AgentAliasId == nil || len(*v.AgentAliasId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member agentAliasId must not be empty")}
	}
	if v.AgentAliasId != nil {
		if err := encoder.SetURI("agentAliasId").String(*v.AgentAliasId); err != nil {
			return err
		}
	}

	if v.AgentId == nil || len(*v.AgentId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member agentId must not be empty")}
	}
	if v.AgentId != nil {
		if err := encoder.SetURI("agentId").String(*v.AgentId); err != nil {
			return err
		}
	}

	if v.SessionId == nil || len(*v.SessionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member sessionId must not be empty")}
	}
	if v.SessionId != nil {
		if err := encoder.SetURI("sessionId").String(*v.SessionId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentInvokeAgentInput(v *InvokeAgentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnableTrace != nil {
		ok := object.Key("enableTrace")
		ok.Boolean(*v.EnableTrace)
	}

	if v.EndSession != nil {
		ok := object.Key("endSession")
		ok.Boolean(*v.EndSession)
	}

	if v.InputText != nil {
		ok := object.Key("inputText")
		ok.String(*v.InputText)
	}

	if v.SessionState != nil {
		ok := object.Key("sessionState")
		if err := awsRestjson1_serializeDocumentSessionState(v.SessionState, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpRetrieve struct {
}

func (*awsRestjson1_serializeOpRetrieve) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpRetrieve) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RetrieveInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/knowledgebases/{knowledgeBaseId}/retrieve")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsRetrieveInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentRetrieveInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsRetrieveInput(v *RetrieveInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.KnowledgeBaseId == nil || len(*v.KnowledgeBaseId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member knowledgeBaseId must not be empty")}
	}
	if v.KnowledgeBaseId != nil {
		if err := encoder.SetURI("knowledgeBaseId").String(*v.KnowledgeBaseId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentRetrieveInput(v *RetrieveInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.RetrievalConfiguration != nil {
		ok := object.Key("retrievalConfiguration")
		if err := awsRestjson1_serializeDocumentKnowledgeBaseRetrievalConfiguration(v.RetrievalConfiguration, ok); err != nil {
			return err
		}
	}

	if v.RetrievalQuery != nil {
		ok := object.Key("retrievalQuery")
		if err := awsRestjson1_serializeDocumentKnowledgeBaseQuery(v.RetrievalQuery, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpRetrieveAndGenerate struct {
}

func (*awsRestjson1_serializeOpRetrieveAndGenerate) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpRetrieveAndGenerate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RetrieveAndGenerateInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/retrieveAndGenerate")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentRetrieveAndGenerateInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsRetrieveAndGenerateInput(v *RetrieveAndGenerateInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentRetrieveAndGenerateInput(v *RetrieveAndGenerateInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Input != nil {
		ok := object.Key("input")
		if err := awsRestjson1_serializeDocumentRetrieveAndGenerateInput(v.Input, ok); err != nil {
			return err
		}
	}

	if v.RetrieveAndGenerateConfiguration != nil {
		ok := object.Key("retrieveAndGenerateConfiguration")
		if err := awsRestjson1_serializeDocumentRetrieveAndGenerateConfiguration(v.RetrieveAndGenerateConfiguration, ok); err != nil {
			return err
		}
	}

	if v.SessionConfiguration != nil {
		ok := object.Key("sessionConfiguration")
		if err := awsRestjson1_serializeDocumentRetrieveAndGenerateSessionConfiguration(v.SessionConfiguration, ok); err != nil {
			return err
		}
	}

	if v.SessionId != nil {
		ok := object.Key("sessionId")
		ok.String(*v.SessionId)
	}

	return nil
}

func awsRestjson1_serializeDocumentAdditionalModelRequestFields(v map[string]document.Interface, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentAdditionalModelRequestFieldsValue(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentAdditionalModelRequestFieldsValue(v document.Interface, value smithyjson.Value) error {
	if v == nil {
		return nil
	}
	if !internaldocument.IsInterface(v) {
		return fmt.Errorf("%T is not a compatible document type", v)
	}
	db, err := v.MarshalSmithyDocument()
	if err != nil {
		return err
	}
	value.Write(db)
	return nil
}

func awsRestjson1_serializeDocumentApiResult(v *types.ApiResult, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ActionGroup != nil {
		ok := object.Key("actionGroup")
		ok.String(*v.ActionGroup)
	}

	if v.ApiPath != nil {
		ok := object.Key("apiPath")
		ok.String(*v.ApiPath)
	}

	if v.HttpMethod != nil {
		ok := object.Key("httpMethod")
		ok.String(*v.HttpMethod)
	}

	if v.HttpStatusCode != nil {
		ok := object.Key("httpStatusCode")
		ok.Integer(*v.HttpStatusCode)
	}

	if v.ResponseBody != nil {
		ok := object.Key("responseBody")
		if err := awsRestjson1_serializeDocumentResponseBody(v.ResponseBody, ok); err != nil {
			return err
		}
	}

	if len(v.ResponseState) > 0 {
		ok := object.Key("responseState")
		ok.String(string(v.ResponseState))
	}

	return nil
}

func awsRestjson1_serializeDocumentByteContentDoc(v *types.ByteContentDoc, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContentType != nil {
		ok := object.Key("contentType")
		ok.String(*v.ContentType)
	}

	if v.Data != nil {
		ok := object.Key("data")
		ok.Base64EncodeBytes(v.Data)
	}

	if v.Identifier != nil {
		ok := object.Key("identifier")
		ok.String(*v.Identifier)
	}

	return nil
}

func awsRestjson1_serializeDocumentContentBody(v *types.ContentBody, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Body != nil {
		ok := object.Key("body")
		ok.String(*v.Body)
	}

	return nil
}

func awsRestjson1_serializeDocumentExternalSource(v *types.ExternalSource, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ByteContent != nil {
		ok := object.Key("byteContent")
		if err := awsRestjson1_serializeDocumentByteContentDoc(v.ByteContent, ok); err != nil {
			return err
		}
	}

	if v.S3Location != nil {
		ok := object.Key("s3Location")
		if err := awsRestjson1_serializeDocumentS3ObjectDoc(v.S3Location, ok); err != nil {
			return err
		}
	}

	if len(v.SourceType) > 0 {
		ok := object.Key("sourceType")
		ok.String(string(v.SourceType))
	}

	return nil
}

func awsRestjson1_serializeDocumentExternalSources(v []types.ExternalSource, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentExternalSource(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentExternalSourcesGenerationConfiguration(v *types.ExternalSourcesGenerationConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AdditionalModelRequestFields != nil {
		ok := object.Key("additionalModelRequestFields")
		if err := awsRestjson1_serializeDocumentAdditionalModelRequestFields(v.AdditionalModelRequestFields, ok); err != nil {
			return err
		}
	}

	if v.GuardrailConfiguration != nil {
		ok := object.Key("guardrailConfiguration")
		if err := awsRestjson1_serializeDocumentGuardrailConfiguration(v.GuardrailConfiguration, ok); err != nil {
			return err
		}
	}

	if v.InferenceConfig != nil {
		ok := object.Key("inferenceConfig")
		if err := awsRestjson1_serializeDocumentInferenceConfig(v.InferenceConfig, ok); err != nil {
			return err
		}
	}

	if v.PromptTemplate != nil {
		ok := object.Key("promptTemplate")
		if err := awsRestjson1_serializeDocumentPromptTemplate(v.PromptTemplate, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentExternalSourcesRetrieveAndGenerateConfiguration(v *types.ExternalSourcesRetrieveAndGenerateConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GenerationConfiguration != nil {
		ok := object.Key("generationConfiguration")
		if err := awsRestjson1_serializeDocumentExternalSourcesGenerationConfiguration(v.GenerationConfiguration, ok); err != nil {
			return err
		}
	}

	if v.ModelArn != nil {
		ok := object.Key("modelArn")
		ok.String(*v.ModelArn)
	}

	if v.Sources != nil {
		ok := object.Key("sources")
		if err := awsRestjson1_serializeDocumentExternalSources(v.Sources, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentFilterAttribute(v *types.FilterAttribute, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("value")
		if err := awsRestjson1_serializeDocumentFilterValue(v.Value, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentFilterValue(v document.Interface, value smithyjson.Value) error {
	if v == nil {
		return nil
	}
	if !internaldocument.IsInterface(v) {
		return fmt.Errorf("%T is not a compatible document type", v)
	}
	db, err := v.MarshalSmithyDocument()
	if err != nil {
		return err
	}
	value.Write(db)
	return nil
}

func awsRestjson1_serializeDocumentFunctionResult(v *types.FunctionResult, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ActionGroup != nil {
		ok := object.Key("actionGroup")
		ok.String(*v.ActionGroup)
	}

	if v.Function != nil {
		ok := object.Key("function")
		ok.String(*v.Function)
	}

	if v.ResponseBody != nil {
		ok := object.Key("responseBody")
		if err := awsRestjson1_serializeDocumentResponseBody(v.ResponseBody, ok); err != nil {
			return err
		}
	}

	if len(v.ResponseState) > 0 {
		ok := object.Key("responseState")
		ok.String(string(v.ResponseState))
	}

	return nil
}

func awsRestjson1_serializeDocumentGenerationConfiguration(v *types.GenerationConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AdditionalModelRequestFields != nil {
		ok := object.Key("additionalModelRequestFields")
		if err := awsRestjson1_serializeDocumentAdditionalModelRequestFields(v.AdditionalModelRequestFields, ok); err != nil {
			return err
		}
	}

	if v.GuardrailConfiguration != nil {
		ok := object.Key("guardrailConfiguration")
		if err := awsRestjson1_serializeDocumentGuardrailConfiguration(v.GuardrailConfiguration, ok); err != nil {
			return err
		}
	}

	if v.InferenceConfig != nil {
		ok := object.Key("inferenceConfig")
		if err := awsRestjson1_serializeDocumentInferenceConfig(v.InferenceConfig, ok); err != nil {
			return err
		}
	}

	if v.PromptTemplate != nil {
		ok := object.Key("promptTemplate")
		if err := awsRestjson1_serializeDocumentPromptTemplate(v.PromptTemplate, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentGuardrailConfiguration(v *types.GuardrailConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GuardrailId != nil {
		ok := object.Key("guardrailId")
		ok.String(*v.GuardrailId)
	}

	if v.GuardrailVersion != nil {
		ok := object.Key("guardrailVersion")
		ok.String(*v.GuardrailVersion)
	}

	return nil
}

func awsRestjson1_serializeDocumentInferenceConfig(v *types.InferenceConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TextInferenceConfig != nil {
		ok := object.Key("textInferenceConfig")
		if err := awsRestjson1_serializeDocumentTextInferenceConfig(v.TextInferenceConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentInvocationResultMember(v types.InvocationResultMember, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.InvocationResultMemberMemberApiResult:
		av := object.Key("apiResult")
		if err := awsRestjson1_serializeDocumentApiResult(&uv.Value, av); err != nil {
			return err
		}

	case *types.InvocationResultMemberMemberFunctionResult:
		av := object.Key("functionResult")
		if err := awsRestjson1_serializeDocumentFunctionResult(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentKnowledgeBaseQuery(v *types.KnowledgeBaseQuery, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Text != nil {
		ok := object.Key("text")
		ok.String(*v.Text)
	}

	return nil
}

func awsRestjson1_serializeDocumentKnowledgeBaseRetrievalConfiguration(v *types.KnowledgeBaseRetrievalConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.VectorSearchConfiguration != nil {
		ok := object.Key("vectorSearchConfiguration")
		if err := awsRestjson1_serializeDocumentKnowledgeBaseVectorSearchConfiguration(v.VectorSearchConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentKnowledgeBaseRetrieveAndGenerateConfiguration(v *types.KnowledgeBaseRetrieveAndGenerateConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GenerationConfiguration != nil {
		ok := object.Key("generationConfiguration")
		if err := awsRestjson1_serializeDocumentGenerationConfiguration(v.GenerationConfiguration, ok); err != nil {
			return err
		}
	}

	if v.KnowledgeBaseId != nil {
		ok := object.Key("knowledgeBaseId")
		ok.String(*v.KnowledgeBaseId)
	}

	if v.ModelArn != nil {
		ok := object.Key("modelArn")
		ok.String(*v.ModelArn)
	}

	if v.RetrievalConfiguration != nil {
		ok := object.Key("retrievalConfiguration")
		if err := awsRestjson1_serializeDocumentKnowledgeBaseRetrievalConfiguration(v.RetrievalConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentKnowledgeBaseVectorSearchConfiguration(v *types.KnowledgeBaseVectorSearchConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filter != nil {
		ok := object.Key("filter")
		if err := awsRestjson1_serializeDocumentRetrievalFilter(v.Filter, ok); err != nil {
			return err
		}
	}

	if v.NumberOfResults != nil {
		ok := object.Key("numberOfResults")
		ok.Integer(*v.NumberOfResults)
	}

	if len(v.OverrideSearchType) > 0 {
		ok := object.Key("overrideSearchType")
		ok.String(string(v.OverrideSearchType))
	}

	return nil
}

func awsRestjson1_serializeDocumentPromptSessionAttributesMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentPromptTemplate(v *types.PromptTemplate, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TextPromptTemplate != nil {
		ok := object.Key("textPromptTemplate")
		ok.String(*v.TextPromptTemplate)
	}

	return nil
}

func awsRestjson1_serializeDocumentRAGStopSequences(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentResponseBody(v map[string]types.ContentBody, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		mapVar := v[key]
		if err := awsRestjson1_serializeDocumentContentBody(&mapVar, om); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentRetrievalFilter(v types.RetrievalFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.RetrievalFilterMemberAndAll:
		av := object.Key("andAll")
		if err := awsRestjson1_serializeDocumentRetrievalFilterList(uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberEquals:
		av := object.Key("equals")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberGreaterThan:
		av := object.Key("greaterThan")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberGreaterThanOrEquals:
		av := object.Key("greaterThanOrEquals")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberIn:
		av := object.Key("in")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberLessThan:
		av := object.Key("lessThan")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberLessThanOrEquals:
		av := object.Key("lessThanOrEquals")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberListContains:
		av := object.Key("listContains")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberNotEquals:
		av := object.Key("notEquals")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberNotIn:
		av := object.Key("notIn")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberOrAll:
		av := object.Key("orAll")
		if err := awsRestjson1_serializeDocumentRetrievalFilterList(uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberStartsWith:
		av := object.Key("startsWith")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	case *types.RetrievalFilterMemberStringContains:
		av := object.Key("stringContains")
		if err := awsRestjson1_serializeDocumentFilterAttribute(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentRetrievalFilterList(v []types.RetrievalFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentRetrievalFilter(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentRetrieveAndGenerateConfiguration(v *types.RetrieveAndGenerateConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExternalSourcesConfiguration != nil {
		ok := object.Key("externalSourcesConfiguration")
		if err := awsRestjson1_serializeDocumentExternalSourcesRetrieveAndGenerateConfiguration(v.ExternalSourcesConfiguration, ok); err != nil {
			return err
		}
	}

	if v.KnowledgeBaseConfiguration != nil {
		ok := object.Key("knowledgeBaseConfiguration")
		if err := awsRestjson1_serializeDocumentKnowledgeBaseRetrieveAndGenerateConfiguration(v.KnowledgeBaseConfiguration, ok); err != nil {
			return err
		}
	}

	if len(v.Type) > 0 {
		ok := object.Key("type")
		ok.String(string(v.Type))
	}

	return nil
}

func awsRestjson1_serializeDocumentRetrieveAndGenerateInput(v *types.RetrieveAndGenerateInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Text != nil {
		ok := object.Key("text")
		ok.String(*v.Text)
	}

	return nil
}

func awsRestjson1_serializeDocumentRetrieveAndGenerateSessionConfiguration(v *types.RetrieveAndGenerateSessionConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.KmsKeyArn != nil {
		ok := object.Key("kmsKeyArn")
		ok.String(*v.KmsKeyArn)
	}

	return nil
}

func awsRestjson1_serializeDocumentReturnControlInvocationResults(v []types.InvocationResultMember, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentInvocationResultMember(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentS3ObjectDoc(v *types.S3ObjectDoc, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Uri != nil {
		ok := object.Key("uri")
		ok.String(*v.Uri)
	}

	return nil
}

func awsRestjson1_serializeDocumentSessionAttributesMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentSessionState(v *types.SessionState, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.InvocationId != nil {
		ok := object.Key("invocationId")
		ok.String(*v.InvocationId)
	}

	if v.PromptSessionAttributes != nil {
		ok := object.Key("promptSessionAttributes")
		if err := awsRestjson1_serializeDocumentPromptSessionAttributesMap(v.PromptSessionAttributes, ok); err != nil {
			return err
		}
	}

	if v.ReturnControlInvocationResults != nil {
		ok := object.Key("returnControlInvocationResults")
		if err := awsRestjson1_serializeDocumentReturnControlInvocationResults(v.ReturnControlInvocationResults, ok); err != nil {
			return err
		}
	}

	if v.SessionAttributes != nil {
		ok := object.Key("sessionAttributes")
		if err := awsRestjson1_serializeDocumentSessionAttributesMap(v.SessionAttributes, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentTextInferenceConfig(v *types.TextInferenceConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxTokens != nil {
		ok := object.Key("maxTokens")
		ok.Integer(*v.MaxTokens)
	}

	if v.StopSequences != nil {
		ok := object.Key("stopSequences")
		if err := awsRestjson1_serializeDocumentRAGStopSequences(v.StopSequences, ok); err != nil {
			return err
		}
	}

	if v.Temperature != nil {
		ok := object.Key("temperature")
		switch {
		case math.IsNaN(float64(*v.Temperature)):
			ok.String("NaN")

		case math.IsInf(float64(*v.Temperature), 1):
			ok.String("Infinity")

		case math.IsInf(float64(*v.Temperature), -1):
			ok.String("-Infinity")

		default:
			ok.Float(*v.Temperature)

		}
	}

	if v.TopP != nil {
		ok := object.Key("topP")
		switch {
		case math.IsNaN(float64(*v.TopP)):
			ok.String("NaN")

		case math.IsInf(float64(*v.TopP), 1):
			ok.String("Infinity")

		case math.IsInf(float64(*v.TopP), -1):
			ok.String("-Infinity")

		default:
			ok.Float(*v.TopP)

		}
	}

	return nil
}
