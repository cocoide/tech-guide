package gateway_test

import (
	"context"
	"testing"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/gateway"
)

func Test_OpenAIGateway(t *testing.T) {
	conf.NewEnv()
	ctx := context.Background()
	og := gateway.NewOpenAIGateway(ctx)
	res, err := og.GetAnswerFromPrompt("[Title: Tailwind CSSはCSS設計に何をもたらすか], Select tags related to the title from [Frontend, SRE, Golang, Ruby, 個人開発, Businness, Markeging, 設計] with weights ranging from 1 to 5. Output only the tags with a weight of 3 or higher in the format `<TagName>: <Weight>,`(do not output anything for tags with weights below 3)", 0)
	if err != nil {
		t.Errorf("func error: %v", err)
	}
	t.Logf("answer: %s", res)
}
