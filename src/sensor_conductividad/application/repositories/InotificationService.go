package repositories

import "Integrador/src/sensor_conductividad/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.ConductivitySensor) error
}
