package metrik

import (
	"net/http"
)

const EXPECTEDSTATUSCODE = 200


type Metriker interface {
	CheckSites(site int) (string, int)
}


type Metrik struct {
	webSites []string
}


func NewMetrik(webSites []string) *Metrik {
	return &Metrik{webSites}
}

func (m *Metrik) CheckSites() (map[string]int) {
	result := make(map[string]int)

	for i := 0; i<len(m.webSites); i++ {
		status, err := http.Get(m.webSites[i])
		if err != nil {
			result[m.webSites[i]] = 404
			continue 
		}
		result[m.webSites[i]] = status.StatusCode
	}

	return result
}