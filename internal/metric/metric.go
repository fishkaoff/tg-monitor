package metric

import (
	"net/http"
)

const EXPECTEDSTATUSCODE = 200

type Metricer interface {
	CheckSites(site int) (string, int)
}

type Metric struct {
	webSites []string
}

func NewMetric(webSites []string) *Metric {
	return &Metric{webSites}
}

func (m *Metric) CheckSites() map[string]int {
	result := make(map[string]int)

	for i := 0; i < len(m.webSites); i++ {
		status, err := http.Get(m.webSites[i])
		if err != nil {
			result[m.webSites[i]] = 404
			continue
		}
		result[m.webSites[i]] = status.StatusCode
	}

	return result
}
