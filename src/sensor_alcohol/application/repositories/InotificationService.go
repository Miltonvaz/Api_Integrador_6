package repositories

import "Integrador/src/sensor_alcohol/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.AlcoholSensor) error
}
