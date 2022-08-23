package common

type DestinationType string

const (
	CoralogixDestinationType  DestinationType = "coralogix"
	GrafanaDestinationType    DestinationType = "grafana"
	DatadogDestinationType    DestinationType = "datadog"
	HoneycombDestinationType  DestinationType = "honeycomb"
	NewRelicDestinationType   DestinationType = "newrelic"
	LogzioDestinationType     DestinationType = "logzio"
	PrometheusDestinationType DestinationType = "prometheus"
	LokiDestinationType       DestinationType = "loki"
	TempoDestinationType      DestinationType = "tempo"
)
