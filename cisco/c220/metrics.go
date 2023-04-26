package c220

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics map[string]*prometheus.GaugeVec

func newServerMetric(metricName string, docString string, constLabels prometheus.Labels, labelNames []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        metricName,
			Help:        docString,
			ConstLabels: constLabels,
		},
		labelNames,
	)
}

func NewDeviceMetrics() *map[string]*metrics {
	var (
		ThermalMetrics = &metrics{
			"fanSpeed":          newServerMetric("c220_thermal_fan_speed", "Current fan speed in the unit of RPM", nil, []string{"name", "chassisSerialNumber"}),
			"fanStatus":         newServerMetric("c220_thermal_fan_status", "Current fan status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber"}),
			"sensorTemperature": newServerMetric("c220_thermal_sensor_temperature", "Current sensor temperature reading in Celsius", nil, []string{"name", "chassisSerialNumber"}),
			"sensorStatus":      newServerMetric("c220_thermal_sensor_status", "Current sensor status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber"}),
			"thermalSummary":    newServerMetric("c220_thermal_summary_status", "Current sensor status 1 = OK, 0 = BAD", nil, []string{"url", "chassisSerialNumber"}),
		}

		PowerMetrics = &metrics{
			"voltageOutput":       newServerMetric("c220_power_voltage_output", "Power voltage output in watts", nil, []string{"name", "chassisSerialNumber"}),
			"voltageStatus":       newServerMetric("c220_power_voltage_status", "Current power voltage status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber"}),
			"supplyOutput":        newServerMetric("c220_power_supply_output", "Power supply output in watts", nil, []string{"name", "chassisSerialNumber", "manufacturer", "partNumber", "serialNumber", "powerSupplyType", "model"}),
			"supplyStatus":        newServerMetric("c220_power_supply_status", "Current power supply status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber", "manufacturer", "partNumber", "serialNumber", "powerSupplyType", "model"}),
			"supplyTotalConsumed": newServerMetric("c220_power_supply_total_consumed", "Total output of all power supplies in watts", nil, []string{"url", "chassisSerialNumber"}),
		}

		MemoryMetrics = &metrics{
			"memoryStatus": newServerMetric("c220_memory_dimm_status", "Current dimm status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber", "capacityMiB", "manufacturer", "partNumber", "serialNumber"}),
		}

		ProcessorMetrics = &metrics{
			"processorStatus": newServerMetric("c220_cpu_status", "Current cpu status 1 = OK, 0 = BAD", nil, []string{"name", "chassisSerialNumber", "description", "totalThreads"}),
		}

		DriveMetrics = &metrics{
			"storageControllerStatus": newServerMetric("c220_storage_controller_status", "Current storage controller status 1 = OK, 0 = BAD, -1 = DISABLED", nil, []string{"name", "chassisSerialNumber", "firmwareVersion", "memberId", "model"}),
			"driveStatus":             newServerMetric("c220_drive_status", "Current drive status 1 = OK, 0 = BAD, -1 = DISABLED", nil, []string{"name", "chassisSerialNumber", "capacityBytes", "id", "model"}),
		}

		DeviceMetrics = &metrics{
			"deviceInfo": newServerMetric("device_info", "Current snapshot of device firmware information", nil, []string{"name", "chassisSerialNumber", "firmwareVersion", "biosVersion", "model"}),
		}

		Metrics = &map[string]*metrics{
			"thermalMetrics":   ThermalMetrics,
			"powerMetrics":     PowerMetrics,
			"memoryMetrics":    MemoryMetrics,
			"processorMetrics": ProcessorMetrics,
			"driveMetrics":     DriveMetrics,
			"deviceInfo":       DeviceMetrics,
		}
	)

	return Metrics
}
