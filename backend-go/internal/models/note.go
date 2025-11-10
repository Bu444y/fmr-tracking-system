package models

import (
	"database/sql"
	"time"
)

type NoteType string

const (
	NoteTypeGeneral    NoteType = "general"
	NoteTypeTechnician NoteType = "technician"
	NoteTypeCustomer   NoteType = "customer"
)

type Note struct {
	ID        int64     `json:"id"`
	FmrID     int64     `json:"fmrId"`
	Content   string    `json:"content"`
	Author    *string   `json:"author,omitempty"`
	NoteType  NoteType  `json:"noteType"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateNoteRequest struct {
	FmrID    int64     `json:"fmrId" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Author   *string   `json:"author,omitempty"`
	NoteType *NoteType `json:"noteType,omitempty"`
}

type NotesListResponse struct {
	Notes []Note `json:"notes"`
}

// Database row type for scanning from SQL
type NoteRow struct {
	ID        int64
	FmrID     int64
	Content   string
	Author    sql.NullString
	NoteType  string
	CreatedAt time.Time
}
