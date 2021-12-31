package main

import (
	"os"
	"strings"

	passport "engram-flow-api-v1/internal"

	vparse "engram-flow-api-v1/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func init() {
	if "LOCAL" == strings.ToUpper(os.Getenv("ENV")) {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// ===========================================================================
	// Load environment variables
	// ===========================================================================
	var (
		env     = strings.ToUpper(os.Getenv("ENV")) // LOCAL, DEV, STG, PRD
		port    = os.Getenv("PORT")                 // server traffic on this port
		version = os.Getenv("VERSION")              // path to VERSION file
	)
	// ===========================================================================
	// Read version information
	// ===========================================================================
	version, err := vparse.ParseVersionFile(version)
	if err != nil {
		log.WithFields(log.Fields{
			"env":  env,
			"err":  err,
			"path": os.Getenv("VERSION"),
		}).Fatal("Can't find a VERSION file")
		return
	}
	log.WithFields(log.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
	}).Info("Loaded VERSION file")

	// ===========================================================================
	// Initialise application context
	// ===========================================================================
	appEnv := passport.AppEnv{
		Render:  render.New(),
		Version: version,
		Env:     env,
		Port:    port,
	}
	// ===========================================================================
	// Start application
	// ===========================================================================
	passport.StartServer(appEnv)
}
