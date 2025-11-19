package swagger

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func descriptionFromCommentAndOptions(ctx Context, comment string, options []string, apiType spec.Type) string {
	desc := formatComment(comment)
	if ctx.YAPIParametersWithExample {
		if exampleVal := exampleValueFromOptions(ctx, options, apiType); exampleVal != nil {
			desc = desc + "\n示例值: " + fmt.Sprint(exampleVal)
		}
	}
	return desc
}
