// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateMonitor struct {
}

func (*validateOpCreateMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMonitor struct {
}

func (*validateOpDeleteMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetHealthEvent struct {
}

func (*validateOpGetHealthEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetHealthEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetHealthEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetHealthEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInternetEvent struct {
}

func (*validateOpGetInternetEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInternetEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInternetEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInternetEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMonitor struct {
}

func (*validateOpGetMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueryResults struct {
}

func (*validateOpGetQueryResults) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueryResults) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueryResultsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueryResultsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueryStatus struct {
}

func (*validateOpGetQueryStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueryStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueryStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueryStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListHealthEvents struct {
}

func (*validateOpListHealthEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListHealthEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListHealthEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListHealthEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartQuery struct {
}

func (*validateOpStartQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopQuery struct {
}

func (*validateOpStopQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMonitor struct {
}

func (*validateOpUpdateMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMonitor{}, middleware.After)
}

func addOpDeleteMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMonitor{}, middleware.After)
}

func addOpGetHealthEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetHealthEvent{}, middleware.After)
}

func addOpGetInternetEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInternetEvent{}, middleware.After)
}

func addOpGetMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMonitor{}, middleware.After)
}

func addOpGetQueryResultsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueryResults{}, middleware.After)
}

func addOpGetQueryStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueryStatus{}, middleware.After)
}

func addOpListHealthEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListHealthEvents{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartQuery{}, middleware.After)
}

func addOpStopQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopQuery{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMonitor{}, middleware.After)
}

func validateOpCreateMonitorInput(v *CreateMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMonitorInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMonitorInput(v *DeleteMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMonitorInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetHealthEventInput(v *GetHealthEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetHealthEventInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if v.EventId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInternetEventInput(v *GetInternetEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInternetEventInput"}
	if v.EventId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMonitorInput(v *GetMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMonitorInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueryResultsInput(v *GetQueryResultsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueryResultsInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueryStatusInput(v *GetQueryStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueryStatusInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListHealthEventsInput(v *ListHealthEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListHealthEventsInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartQueryInput(v *StartQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartQueryInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if len(v.QueryType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("QueryType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopQueryInput(v *StopQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopQueryInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if v.QueryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMonitorInput(v *UpdateMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMonitorInput"}
	if v.MonitorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MonitorName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
