package pathutils

import (
	"fmt"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/google/uuid"
	"os"
)

func StartSignupURL(sessionID uuid.UUID) string {
	return fmt.Sprintf("%s/api/session/signup?sessionId=%s", os.Getenv("FRONTEND_URL"), sessionID.String())
}

func CompleteSignupURL(tokens *usecase.GenerateTokensResponse) string {
	return fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
}

func LoginURL(tokens *usecase.GenerateTokensResponse) string {
	return fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
}
