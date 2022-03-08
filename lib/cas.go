package lib

import (
	"net/url"

	"github.com/spf13/viper"
	"gopkg.in/cas.v2"
)

var CasClient *cas.Client

func InitCas() {

	url, _ := url.Parse(viper.GetString("CAS_URL"))
	CasClient = cas.NewClient(&cas.Options{
		URL: url,
	})
}
