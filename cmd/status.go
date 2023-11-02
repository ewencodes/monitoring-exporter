/*
Copyright © 2023 ewencodes
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/ewencodes/monitoring-job/internal/middlewares"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var url string
var metric_prefix string
var port string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use: "status",
	Run: func(cmd *cobra.Command, args []string) {
		reg := prometheus.NewRegistry()

		status_gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metric_prefix,
			Name:      "status",
		})
		response_time_gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metric_prefix,
			Name:      "response_time",
		})

		err := reg.Register(status_gauge)

		if err != nil {
			fmt.Println(err)
		}

		err = reg.Register(response_time_gauge)

		if err != nil {
			fmt.Println(err)
		}

		mux := http.NewServeMux()

		mux.Handle("/metrics", middlewares.StatusMiddleware(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), url, status_gauge, response_time_gauge))

		err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringVar(&url, "url", "", "Url to get status")
	statusCmd.MarkFlagRequired("url")

	statusCmd.Flags().StringVar(&metric_prefix, "metric-prefix", "", "Prefix of the metric")
	statusCmd.MarkFlagRequired("metric-name")

	statusCmd.Flags().StringVar(&port, "port", "8081", "Port to expose metrics (default: 8081)")
}
