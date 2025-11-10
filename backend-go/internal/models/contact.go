package models

import (
	"database/sql"
	"time"
)

type Team string

const (
	TeamFilius Team = "The Filius Team"
	TeamCRC    Team = "CRC AN/TYQ-23A Program Team"
)

type Contact struct {
	ID             int64     `json:"id"`
	Title          *string   `json:"title,omitempty"`
	SequenceNumber *int      `json:"sequenceNumber,omitempty"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          *string   `json:"email,omitempty"`
	MobileNumber   *string   `json:"mobileNumber,omitempty"`
	WorkNumber     *string   `json:"workNumber,omitempty"`
	Organization   *string   `json:"organization,omitempty"`
	Notes          *string   `json:"notes,omitempty"`
	Team           Team      `json:"team"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type CreateContactRequest struct {
	Title          *string `json:"title,omitempty"`
	SequenceNumber *int    `json:"sequenceNumber,omitempty"`
	FirstName      string  `json:"firstName" binding:"required"`
	LastName       string  `json:"lastName" binding:"required"`
	Email          *string `json:"email,omitempty"`
	MobileNumber   *string `json:"mobileNumber,omitempty"`
	WorkNumber     *string `json:"workNumber,omitempty"`
	Organization   *string `json:"organization,omitempty"`
	Notes          *string `json:"notes,omitempty"`
	Team           Team    `json:"team" binding:"required"`
}

type UpdateContactRequest struct {
	ID             int64   `json:"id" binding:"required"`
	Title          *string `json:"title,omitempty"`
	SequenceNumber *int    `json:"sequenceNumber,omitempty"`
	FirstName      *string `json:"firstName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	Email          *string `json:"email,omitempty"`
	MobileNumber   *string `json:"mobileNumber,omitempty"`
	WorkNumber     *string `json:"workNumber,omitempty"`
	Organization   *string `json:"organization,omitempty"`
	Notes          *string `json:"notes,omitempty"`
	Team           *Team   `json:"team,omitempty"`
}

type ContactsListResponse struct {
	Contacts []Contact `json:"contacts"`
	Total    int       `json:"total"`
}

// Database row type for scanning from SQL
type ContactRow struct {
	ID             int64
	Title          sql.NullString
	SequenceNumber sql.NullInt64
	FirstName      string
	LastName       string
	Email          sql.NullString
	MobileNumber   sql.NullString
	WorkNumber     sql.NullString
	Organization   sql.NullString
	Notes          sql.NullString
	Team           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
