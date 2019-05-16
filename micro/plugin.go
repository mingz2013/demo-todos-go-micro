package main

import (
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
)

func init() {
	plugin.Register(cors.NewPlugin())
}
