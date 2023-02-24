package middlewares

import (
	"net/url"
)

type Middlwares struct {
}

func NewMiddlewares() *Middlwares {
	return &Middlwares{}
}

func (mw *Middlwares) CheckUrl(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}

	return true
}


func (mw *Middlwares) CheckMatches(webSites []string, site string) bool {
	for _, webSite := range webSites {
		if webSite == site {
			return true
		}
	}

	return false 
}
