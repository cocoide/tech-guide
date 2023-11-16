package pathutils

import (
	"fmt"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"os"
)

func GetRegisterProfilePath(tokens usecase.GenerateTokensResponse) string {
	return fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
}
func GetRegisterProfilePath(tokens usecase.GenerateTokensResponse) string {
	return fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
}
