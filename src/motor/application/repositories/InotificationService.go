package repositories

import "Integrador/src/motor/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.MotorFermentador) error
}
