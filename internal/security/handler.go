package security

import (
	"context"

	"github.com/tiagoposse/connect/ent/ogent"
	authz "github.com/tiagoposse/go-auth/authorization"
	"github.com/tiagoposse/go-auth/sessions"
	"github.com/tiagoposse/ogent-auth/authorization"
)


func NewAuthorizationHandler(val authorization.IAuthTokenValidator, scopes map[string]authz.Scopes) *authorization.AuthorizationHandler {
	ctrl := authz.NewAuthzController(scopes)

  return &authorization.AuthorizationHandler{
    AuthzController: ctrl,
    IAuthTokenValidator: val,
  }
}

func (h *authorization.AuthorizationHandler) HandleApiKeyAuth(c context.Context, operationName string, t ogent.ApiKeyAuth) (context.Context, error) {
  session, err := h.ValidateApiKeyAuth(c, t.GetAPIKey())
	if err != nil {
		return c, err
	}

  ctx := context.WithValue(c, sessions.ContextSessionKey{}, session)
  if h.ValidateScopes(ctx, operationName, session.GetScopes()); err != nil {
    return ctx, nil
  }

	return ctx, nil
}

func (h *authorization.AuthorizationHandler) HandleCookieAuth(c context.Context, operationName string, t ogent.ApiKeyAuth) (context.Context, error) {
  session, err := h.ValidateCookieAuth(c, t.GetAPIKey())
	if err != nil {
		return c, err
	}

  ctx := context.WithValue(c, sessions.ContextSessionKey{}, session)
  if h.ValidateScopes(ctx, operationName, session.GetScopes()); err != nil {
    return ctx, nil
  }

	return ctx, nil
}
