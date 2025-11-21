package swagger

import (
	"testing"

	"github.com/duantiao/go-zero/tools/goctl/api/spec"
)

type Context struct {
	UseDefinitions            bool
	WrapCodeMsg               bool
	BizCodeEnumDescription    string
	YAPIParametersWithExample bool
}

func testingContext(_ *testing.T) Context {
	return Context{}
}

func contextFromApi(info spec.Info) Context {
	if len(info.Properties) == 0 {
		return Context{}
	}
	return Context{
		UseDefinitions:            getBoolFromKVOrDefault(info.Properties, propertyKeyUseDefinitions, defaultValueOfPropertyUseDefinition),
		WrapCodeMsg:               getBoolFromKVOrDefault(info.Properties, propertyKeyWrapCodeMsg, false),
		BizCodeEnumDescription:    getStringFromKVOrDefault(info.Properties, propertyKeyBizCodeEnumDescription, "业务响应码"),
		YAPIParametersWithExample: getBoolFromKVOrDefault(info.Properties, propertyKeyYAPIParametersWithExample, false),
	}
}
