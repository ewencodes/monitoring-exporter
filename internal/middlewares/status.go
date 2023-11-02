/*
Copyright © 2023 ewencodes
*/
package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func StatusMiddleware(next http.Handler, url string, status_gauge prometheus.Gauge, response_time_gauge prometheus.Gauge) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		response, err := http.Get(url)

		elapsed := time.Since(start).Milliseconds()

		if err != nil || response.StatusCode > 399 || response.StatusCode < 199 {
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Error while requesting %s: %d \n", url, response.StatusCode)
			}

			status_gauge.Set(0)
			response_time_gauge.Set(float64(elapsed))
		} else {
			status_gauge.Set(1)
			response_time_gauge.Set(float64(elapsed))
		}

		next.ServeHTTP(w, r)
	})
}
