package web

import (
	"embed"
)

//go:embed swagger-ui/*
var WebUI embed.FS
