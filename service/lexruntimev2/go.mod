module github.com/aws/aws-sdk-go-v2/service/lexruntimev2

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.26.2
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.6
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.6
	github.com/aws/smithy-go v1.20.2
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
