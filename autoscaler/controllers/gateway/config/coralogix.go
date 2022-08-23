package config

import (
	odigosv1 "github.com/keyval-dev/odigos/api/v1alpha1"
	commonconf "github.com/keyval-dev/odigos/autoscaler/controllers/common"
	"github.com/keyval-dev/odigos/common"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	coralogixLogsEndpoint    = "CORALOGIX_LOGS_ENDPOINT"
	coralogixMetricsEndpoint = "CORALOGIX_METRICS_ENDPOINT"
	coralogixTracesEndpoint  = "CORALOGIX_TRACES_ENDPOINT"
	coralogixAppName         = "CORALOGIX_APPNAME"
	coralogixSubsystemName   = "CORALOGIX_SUBSYSTEMNAME"
)

type Coralogix struct{}

func (c *Coralogix) DestType() common.DestinationType {
	return common.CoralogixDestinationType
}

func (c *Coralogix) ModifyConfig(dest *odigosv1.Destination, currentConfig *commonconf.Config) {
	if isTracingEnabled(dest) {
		currentConfig.Exporters["coralogix"] = commonconf.GenericMap{
			"endpoint": coralogixTracesEndpoint,
		}
		currentConfig.Service.Pipelines["traces/coralogix"] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch", "memory_limiter", "k8sattributes"},
			Exporters:  []string{"coralogix"},
		}
	}
	if isMetricsEnabled(dest) {
		currentConfig.Exporters["coralogix"] = commonconf.GenericMap{
			"endpoint": coralogixMetricsEndpoint,
		}
		currentConfig.Service.Pipelines["metrics/coralogix"] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch", "memory_limiter", "k8sattributes"},
			Exporters:  []string{"coralogix"},
		}
	}
	if isLoggingEnabled(dest) {
		currentConfig.Exporters["coralogix"] = commonconf.GenericMap{
			"endpoint": coralogixTracesEndpoint,
		}
		currentConfig.Service.Pipelines["logs/coralogix"] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch", "memory_limiter", "k8sattributes"},
			Exporters:  []string{"coralogix"},
		}
	}
	an, exists := dest.Spec.Data[coralogixAppName]
	if !exists {
		ah = "odigos"
		log.Log.V(0).Info("Application name was not specified, using odigos as default value")
	}
	sn, exists := dest.Spec.Data[coralogixSubsystemName]
	if !exists {
		log.Log.V(0).Info("Subsystem name was not specified")
	}
}
