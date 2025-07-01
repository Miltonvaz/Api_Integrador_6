package repositories

import "Integrador/src/sensor_temperatura/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.TemperatureSensor) error
}
