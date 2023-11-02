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
var metric_name string
var port string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use: "status",
	Run: func(cmd *cobra.Command, args []string) {
		reg := prometheus.NewRegistry()

		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: metric_name,
		})

		err := reg.Register(gauge)

		if err != nil {
			fmt.Println(err)
		}

		mux := http.NewServeMux()

		mux.Handle("/metrics", middlewares.StatusMiddleware(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), url, gauge))

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

	statusCmd.Flags().StringVar(&metric_name, "metric-name", "", "Name of the metric")
	statusCmd.MarkFlagRequired("metric-name")

	statusCmd.Flags().StringVar(&port, "port", "8081", "Port to expose metrics (default: 8081)")
}
