package httpserver

import (
	"context"
	"time"

	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/apikey"
	authz "github.com/tiagoposse/go-auth/authorization"
	sessions "github.com/tiagoposse/go-auth/sessions"
)

type SessionInfo struct {
	ID        string       `json:"id"`
	PhotoURL  string       `json:"photo_url"`
	Provider  string       `json:"provider"`
	Email     string       `json:"email"`
	Firstname string       `json:"firstname"`
	Lastname  string       `json:"lastname"`
	Group     string       `json:"group"`
	Scopes    authz.Scopes `json:"scopes"`
}

type SecurityHandler struct {
	*sessions.SessionsController
	client *ent.Client
}

func NewSecurityHandler(cfg *config.SessionsConfig, client *ent.Client) *SecurityHandler {
	return &SecurityHandler{
		client:             client,
		SessionsController: sessions.NewSessionsController(*cfg.SessionKey.Value, time.Duration(cfg.JwtExpiration)),
	}
}

func (h *SecurityHandler) ValidateApiKeyAuth(ctx context.Context, key string) (authz.ScopedSession, error) {
	// .WithUser(func(uq *ent.UserQuery) { uq.WithGroup() })
	_, err := h.client.ApiKey.Query().Where(apikey.KeyEQ(key)).First(ctx)
	if err != nil {
		return nil, err
	}
	// if ak.Edges.User == nil {
	// 	return nil, errors.New("user not found")
	// }

	return &sessions.Session{
		SessionInfo: SessionInfo{
			// ID:        ak.Edges.User.ID,
			// PhotoURL:  ak.Edges.User.PhotoURL,
			// Provider:  ak.Edges.User.Provider,
			// Email:     ak.Edges.User.Email,
			// Firstname: ak.Edges.User.Firstname,
			// Lastname:  ak.Edges.User.Lastname,
			// Group:     ak.Edges.User.Edges.Group.ID,
			// Scopes:    ak.Scopes,
		},
	}, nil
}

func (h *SecurityHandler) ValidateCookieAuth(ctx context.Context, key string) (authz.ScopedSession, error) {
	isess, err := h.ValidateSessionToken(ctx, key)
	info := isess.(*sessions.Session).SessionInfo.(map[string]interface{})

	scopes := authz.Scopes{}
	for _, s := range info["scopes"].([]interface{}) {
		scopes = append(scopes, authz.Scope(s.(string)))
	}

	return &sessions.Session{SessionInfo: SessionInfo{
		ID:        info["id"].(string),
		Email:     info["email"].(string),
		Provider:  info["provider"].(string),
		Firstname: info["firstname"].(string),
		Lastname:  info["lastname"].(string),
		PhotoURL:  info["photo_url"].(string),
		Scopes:    scopes,
		Group:     info["group"].(string),
	}}, err
}
