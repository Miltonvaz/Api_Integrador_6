package repositories

import "Integrador/src/sensor_turbuidez/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.TurbiditySensor) error
}
