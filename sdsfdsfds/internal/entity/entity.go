package entity

import "time"

// Entity represents a domain entity (inner layer of Clean Architecture).
// Business rules at the enterprise level; independent of frameworks and DB.
// Replace with your entities (User, Order, etc.).
type Entity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
