package repositories

import "Integrador/src/sensor_densidad_o/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.DensitySensor) error
}
