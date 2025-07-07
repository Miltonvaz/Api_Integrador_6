package adapters

import (
	"Integrador/src/sensor_turbuidez/application/repositories"
	"Integrador/src/sensor_turbuidez/domain/entities"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repositories.NotificationPort = (*RabbitMQAdapter)(nil)

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is not defined in the .env file")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	if err := declareQueue(ch); err != nil {
		return nil, err
	}

	if err := ch.Confirm(false); err != nil {
		return nil, fmt.Errorf("failed to enable message confirmations: %w", err)
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"sesnsor.turbidity",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	return nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.TurbiditySensor) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal event data: %w", err)
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",
		"sesnsor.turbidity",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	go handlePublishConfirmation(ack, nack)

	log.Printf("Event published: %s", eventType)
	return nil
}

func handlePublishConfirmation(ack, nack <-chan uint64) {
	select {
	case <-ack:
		log.Println("Message confirmed by RabbitMQ")
	case <-nack:
		log.Println("Message not confirmed by RabbitMQ")
	}
}

func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	}
}
