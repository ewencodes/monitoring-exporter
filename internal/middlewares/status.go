/*
Copyright © 2023 ewencodes
*/
package middlewares

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func StatusMiddleware(next http.Handler, url string, gauge prometheus.Gauge) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(url)

		if err != nil || response.StatusCode > 399 || response.StatusCode < 199 {
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Error while requesting %s: %d \n", url, response.StatusCode)
			}

			gauge.Set(0)
		} else {
			gauge.Set(1)
		}

		next.ServeHTTP(w, r)
	})
}
