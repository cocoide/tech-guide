package pathutils

import (
	"fmt"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/google/uuid"
	"os"
)

func StartSignupSessionURL(sessionID *uuid.UUID) string {
	return fmt.Sprintf("%s/api/session/signup?sessionId=%s", os.Getenv("FRONTEND_URL"), sessionID.String())
}

func CompleteSignupSessionURL(tokens *usecase.GenerateTokensResponse) string {
	return fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
}
