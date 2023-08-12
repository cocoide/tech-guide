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
	res, err := og.GetAnswerFromPrompt("[要約: Terraformを使用してDatadogのモニタリング監視を構築・改善した話。構築前の状況や目的、Datadogタグの目的、エージェント導入、コンテナ環境変数追加などの手順を紹介。特にCPU監視のモニター設定の例も示し、Queryの設定方法について説明。], 以下のTagの中から要約内容に関連するものを選んで[Observation,Datadog,Frontend, SRE, Golang, Ruby, 個人開発, Businness, Marketing, IAC] Weightの範囲は1~5。 ただし、Weightが3以上のもので、かつ、 `<TagName>: <Weight>このフォーマットの値だけを出力して", 0.01)
	if err != nil {
		t.Errorf("func error: %v", err)
	}
	t.Logf("answer: %s", res)
}
