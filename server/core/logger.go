package core

import (
	"god2admin/global"

	"github.com/sirupsen/logrus"
)

func initLogger() {
	logger := logrus.New()
	global.LOGGER = logger
}
