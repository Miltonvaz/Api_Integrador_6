package repositories

import (
	"Integrador/src/motor/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(appointment entities.MotorFermentador) error {
	log.Println("Notificando la creación")

	err := sn.notificationPort.PublishEvent("creado", appointment)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}
