package repositories

import "Integrador/src/fermentation/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.Fermentation) error
}
