package radarr

import (
	// "fmt"
	"net/http"

	cfg "github.com/PFNASS/go-arrs/pkg/config"
)

type RadarrConfig struct {
	IP		string
	Port	string
	Token	string
	BaseURL string
}

type RadarrClient struct {
	httpClient		*http.Client
	Config			*RadarrConfig
	BaseURL			string
}