package server

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/controller"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
)

type Server struct {
	*ogent.OgentHandler
	client *ent.Client
	ac     *controller.AccessController
}

func StartServer(cfg *config.Config) {
	authHandler := ogauth.NewSecurityHandler(NewSecurityHandler(cfg.Sessions, client))
	// Start listening.
	srv, err := ogent.NewServer(
		c,
		c.GetAuthorizationHandler(),
		ogent.WithPathPrefix("/api/v1"),
	)
	if err != nil {
		log.Fatal(err)
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
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
