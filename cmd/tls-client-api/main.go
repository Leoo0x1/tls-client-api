package main

import (
	"github.com/Leoo0x1/tls-client-api/internal/tls-client-api/api"
	"github.com/justtrackio/gosoline/pkg/application"
	"github.com/justtrackio/gosoline/pkg/httpserver"
	"os"
	"path/filepath"
)

func main() {
	ex, _ := os.Executable()
	configFilePath := filepath.Join(filepath.Dir(ex), "config.dist.yml")
	// configFilePath := "config.dist.yml"

	application.Run(
		application.WithConfigFile(configFilePath, "yml"),
		application.WithConfigFileFlag,
		application.WithModuleFactory("tls-client-api", httpserver.New("tls-client-api", api.DefineRouter)),
	)
}
