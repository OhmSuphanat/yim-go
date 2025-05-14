package backofficeproperty

import (
	"log"
	baseproperty "yim-go/shared/external/configuration/base-property"

	"github.com/kelseyhightower/envconfig"
)

type ServerProperties baseproperty.ServerProperties

var server ServerProperties

func Get() ServerProperties {
	return server
}

var ignoreLogPaths map[string]bool

func IgnoreLogPaths(s string) bool {
	return ignoreLogPaths[s]
}

func GetIgnoreLogPathsMap() map[string]bool {
	return ignoreLogPaths
}

var ignoreClientLogPaths map[string]bool

func IgnoreClientLogPaths(s string) bool {
	return ignoreClientLogPaths[s]
}

var environVariables map[string]string

func EnvironVariable(s string) string {
	return environVariables[s]
}

func EnvironVariables() map[string]string {
	m := make(map[string]string)
	for k, v := range environVariables {
		m[k] = v
	}
	return m
}

func InitServerProperties() {
	// parse server config
	if err := envconfig.Process("", &baseproperty.Server); err != nil {
		log.Fatalf("read env error: %s", err.Error())
	}
	server = ServerProperties(baseproperty.Server)

	// init env var
	baseproperty.InitEvironVariables()
	environVariables = baseproperty.EnvironVariables

	// init ignore log path
	baseproperty.InitIgnoreLogPaths()
	ignoreLogPaths = baseproperty.IgnoreLogPaths

	// init ignore client log path
	baseproperty.InitIgnoreClientLogPaths()
	ignoreClientLogPaths = baseproperty.IgnoreClientLogPaths

}
