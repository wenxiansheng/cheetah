package initialize

import (
	_ "github.com/test-instructor/cheetah/server/source/example"
	_ "github.com/test-instructor/cheetah/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
