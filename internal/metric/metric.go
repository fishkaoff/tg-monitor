package metric

import (
	"net/http"
)

const EXPECTEDSTATUSCODE = 200

type Metricer interface {
	CheckSites(websites []string) map[string]int
}

type Metric struct {
}

func NewMetric() *Metric {
	return &Metric{}
}

func (m *Metric) CheckSites(webSites []string) map[string]int {
	result := make(map[string]int)

	for i := 0; i < len(webSites); i++ {
		status, err := http.Get(webSites[i])
		if err != nil {
			result[webSites[i]] = 404
			continue
		}
		result[webSites[i]] = status.StatusCode
	}

	return result
}
