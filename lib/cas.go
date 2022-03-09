package lib

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"gopkg.in/cas.v2"
)

var CasClient *cas.Client

var SessionsStore cas.SessionStore
var TicketStore *cas.MemoryStore

func InitCas() {

	SessionsStore = cas.NewMemorySessionStore()
	TicketStore = &cas.MemoryStore{}

	url, _ := url.Parse(viper.GetString("CAS_URL"))
	CasClient = cas.NewClient(&cas.Options{
		Store:        TicketStore,
		URL:          url,
		SessionStore: SessionsStore,
		Cookie: &http.Cookie{
			Path:     fmt.Sprint("/", viper.GetString("APPNAME")),
			MaxAge:   86400,
			HttpOnly: false,
			Secure:   false,
		},
	})

}
