package database

import (
	"database/sql"
	"fmt"
	"hub/src/models"
	"log"
)

// DBHandler struct to hold the database connection
type DBHandler struct {
	db *sql.DB
}

// NewDBHandler creates a new DBHandler
func NewDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{db: db}
}

// createTable creates the subscribers table if it doesn't exist
func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS subscribers (
            id VARCHAR(255) PRIMARY KEY,
            callback_url VARCHAR(2048),
            topic VARCHAR(255),
            secret VARCHAR(255)
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create subscribers table: %w", err)
	}

	fmt.Println("Subscribers table created successfully")
	return nil
}

// AddSubscriber adds a new subscriber to the table
func (handler *DBHandler) AddSubscriber(id, callbackURL, topic, secret string) error {
	_, err := handler.db.Exec(`
        INSERT INTO subscribers (id, callback_url, topic, secret) 
        VALUES (?, ?, ?, ?)
    `, id, callbackURL, topic, secret)

	if err != nil {
		return fmt.Errorf("failed to add subscriber: %w", err)
	}
	return nil
}

// RemoveSubscriber removes a subscriber from the table based on the id
func (handler *DBHandler) RemoveSubscriber(id string) error {
	_, err := handler.db.Exec(`
        DELETE FROM subscribers WHERE id = ?
    `, id)

	if err != nil {
		return fmt.Errorf("failed to remove subscriber: %w", err)
	}
	return nil
}

// Get all subscribers with the given topic
func (handler *DBHandler) GetTopicSubscribers(topic string) ([]models.Subscriber, error) {
	rows, err := handler.db.Query(`
		SELECT * FROM subscribers WHERE topic = ?
	`, topic)
	if err != nil {
		log.Println("Error in GetTopicSubscribers:", err)
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	var subscribers []models.Subscriber
	for rows.Next() {
		var id string
		var callbackURL string
		var secret string
		var topic string
		var subscriber models.Subscriber
		if err := rows.Scan(&id, &callbackURL, &topic, &secret); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan subscribers: %w", err)
		}
		subscriber.ID = id
		subscriber.CallbackURL = callbackURL
		subscriber.Topic = topic
		subscriber.Secret = secret
		subscribers = append(subscribers, subscriber)
	}

	log.Println("Subscribers retrieved:", subscribers)
	return subscribers, nil
}
