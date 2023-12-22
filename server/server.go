package server

import (
	"crypto/tls"
	"net/http"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/controller"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	ogauth "github.com/orbit-ops/launchpad-core/ent/ogentauth"
	"github.com/orbit-ops/launchpad-core/providers"
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
	ac     *controller.AccessController
}

func Start(cfg *config.Config, prov providers.Provider, client *ent.Client) error {
	authHandler := ogauth.NewSecurityHandler(NewSecurityHandler(cfg.Sessions, client))

	ac, err := controller.NewAccessController(prov, client)
	if err != nil {
		return err
	}

	// Start listening.
	h := &handler{
		ac: ac,
	}

	srv, err := ogent.NewServer(
		h,
		ogauth.NewSecurityHandler(authHandler),
		ogent.WithPathPrefix("/api/v1"),
	)
	if err != nil {
		return err
	}

	corsOpts := []handlers.CORSOption{
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"POST", "PATCH", "GET", "DELETE"}),
		handlers.AllowedHeaders([]string{"content-type", "x-page", "x-items-per-page"}),
		handlers.ExposedHeaders([]string{"x-total"}),
	}

	// var origin string
	// if cfg.Web.ServeFrontend {
	// 	origin = cfg.Web.ExternalUrl
	// } else {
	// 	origin = cfg.Web.FrontendUrl
	// }
	// corsOpts = append(corsOpts, handlers.AllowedOrigins([]string{origin}))

	server := http.Server{
		Handler: handlers.CORS(corsOpts...)(srv),
		Addr:    cfg.ApiUrl,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	log.Debugf("Serving at %s\n", server.Addr)
	// err = server.ListenAndServeTLS(cfg.Web.Tls.Certificate, cfg.Web.Tls.Key)
	return server.ListenAndServe()
}
