package baseproperty

import (
	"os"
	"strings"
)

type ServerProperties struct {
	DebugMode            bool     `envconfig:"DEBUG_MODE" long:"debug-mode" default:"false"`
	PrintConsoleFormat   bool     `envconfig:"CONSOLE_FORMAT" long:"print-console-format" default:"false"`
	ShutdownTimeout      int64    `envconfig:"SHUTDOWN_TIMEOUT" long:"shutdown-timeout" default:"300"`
	Port                 string   `envconfig:"PORT" long:"port" default:"8080"`
	ProjectID            string   `envconfig:"GOOGLE_CLOUD_PROJECT" long:"project-id" default:""`
	ServiceName          string   `envconfig:"SERVICE_NAME" long:"service-name" default:""`
	ServiceDescription   string   `envconfig:"SERVICE_DESCRIPTION" long:"service-description" default:""`
	RunLocal             bool     `envconfig:"RUN_LOCAL" long:"run-local" default:"false"`
	LogIgnorePaths       string   `envconfig:"LOG_IGNORE_PATHS" long:"log-ignore-paths" default:""`
	ApiDocs              bool     `envconfig:"API_DOCS" long:"api-docs" default:"false"`
	ApiDocsSchema        string   `envconfig:"API_DOCS_SCHEMA" long:"api-docs-schema" default:"https"`
	ApiDocsHost          string   `envconfig:"API_DOCS_HOST" long:"api-docs-host" default:"localhost:8080"`
	ApiDocsVersion       string   `envconfig:"API_DOCS_VERSION" long:"api-docs-version" default:"0.0.1"`
	LogClientIgnorePaths string   `envconfig:"LOG_CLIENT_IGNORE_PATHS" long:"log-client-ignore-paths" default:""`
	Host                 string   `envconfig:"HOST" long:"host" default:"localhost"`
	GinMode              string   `envconfig:"GIN_MODE" long:"gin-mode"`
	ClientLogMasking     bool     `envconfig:"CLIENT_LOG_MASKING" long:"client-log-masking"`
	WAppJwtExpiry        Duration `envconfig:"W_APP_JWT_EXPIRY" long:"w-app-jwt-expiry" default:"15m"`
}

var (
	Server               = ServerProperties{}
	IgnoreLogPaths       = map[string]bool{}
	IgnoreClientLogPaths = map[string]bool{}
	EnvironVariables     = map[string]string{}
)

func InitEvironVariables() {
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i > 0 {
			EnvironVariables[e[:i]] = e[i+1:]
		}
	}
}

func InitIgnoreLogPaths() {
	LogIgnorePathList := strings.Split(Server.LogIgnorePaths, ",")
	for _, paths := range LogIgnorePathList {
		paths = strings.ReplaceAll(paths, " ", "")
		if paths != "" {
			IgnoreLogPaths[paths] = true
		}
	}
}

func InitIgnoreClientLogPaths() {
	LogClientLogPaths := strings.Split(Server.LogClientIgnorePaths, ",")
	for _, paths := range LogClientLogPaths {
		paths = strings.ReplaceAll(paths, " ", "")
		if paths != "" {
			IgnoreClientLogPaths[paths] = true
		}
	}
}
