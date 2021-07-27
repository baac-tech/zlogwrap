package main

import (
	"time"

	"github.com/buildingwatsize/zlogwrap"
)

func seeMoreUsage() {
	time.Sleep(1000)

	// You can define the service name as header of log message
	serviceName := "[Header in message]:"
	logger2 := zlogwrap.New()
	logger2.Debug()
	logger2.Debug(serviceName, "Debug Log")
	logger2.Info(serviceName, "Info Log")
	logger2.Error(serviceName, "Error Log")

	time.Sleep(1000)

	// The better way to define the service name
	serviceNameWithConfig := "LOGGER3"
	logger3 := zlogwrap.New(zlogwrap.Config{
		ServiceName: serviceNameWithConfig,
	})
	logger3.Debug()
	logger3.Debug("Debug Log")
	logger3.Info("Info Log")
	logger3.Error("Error Log")
}
