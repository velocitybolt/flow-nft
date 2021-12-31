package internal

import (
	"github.com/unrolled/render"
)

// AppEnv holds application configuration data
type AppEnv struct {
	Render  *render.Render
	Version string
	Env     string
	Port    string
}
