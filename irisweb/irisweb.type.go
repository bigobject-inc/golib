package irisweb

import (
	"github.com/kataras/iris/v12"
)

type SSLConfig struct {
	Port     int    `json:"port"`
	KeyFile  string `json:"keyFile"`
	CertFile string `json:"certFile"`
}

// Configure web configure
type Configure struct {
	Host string             `json:"host"`
	Port int                `json:"port"`
	SSL  *SSLConfig         `json:"ssl"`
	Iris iris.Configuration `json:"iris"`
}

// HandleErrorParams web handle error params
type HandleErrorParams struct {
	Code           string
	DisplayMessage string
	HandleName     string
	Message        string
	StatusCode     int
	Params         map[string]interface{}
}

// HandleErrorResponse web handle error response
type HandleErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Params  map[string]interface{} `json:"params"`
}

// Paths Web paths
type Paths struct {
	Root        string
	StaticPaths map[string]StaticPath
}

// Route Web route
type Route struct {
	Handle  []iris.Handler
	Method  string
	Name    string
	Path    string
	Summary string
}

// StaticFile Web static file
type StaticFile struct {
	Pathname string `json:"path"`
	Filename string `json:"file"`
	Absolute string `json:"absolute"`
	Relative string `json:"relative"`
}

// StaticPath Web AR(absolute and relative)Path
type StaticPath struct {
	Absolute string `json:"absolute"`
	Relative string `json:"relative"`
}
