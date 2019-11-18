package alldebrid

import (
	"net/http"

	"github.com/auriou/godebridaria/config"
)

type Queries map[string]string
type ClientAlldebrid struct {
	Base          string
	Agent         string
	HTTPClient    *http.Client
	ContextConfig *config.ClientConfig
}
