package repositories

import entities "Integrador/src/sensor_ph/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.PhSensor) error
}
