package repository

import (
	"context"
	"fmt"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
)

type NoteRepository struct{}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{}
}

func (r *NoteRepository) Create(ctx context.Context, req *models.CreateNoteRequest) (*models.Note, error) {
	noteType := models.NoteTypeGeneral
	if req.NoteType != nil {
		noteType = *req.NoteType
	}

	query := `
		INSERT INTO notes (fmr_id, content, author, note_type)
		VALUES ($1, $2, $3, $4)
		RETURNING id, fmr_id, content, author, note_type, created_at
	`

	var row models.NoteRow
	err := DB.QueryRow(ctx, query, req.FmrID, req.Content, req.Author, noteType).Scan(
		&row.ID, &row.FmrID, &row.Content, &row.Author, &row.NoteType, &row.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create note: %w", err)
	}

	return rowToNote(&row), nil
}

func (r *NoteRepository) ListByFMRID(ctx context.Context, fmrID int64) (*models.NotesListResponse, error) {
	query := `
		SELECT id, fmr_id, content, author, note_type, created_at
		FROM notes
		WHERE fmr_id = $1
		ORDER BY created_at DESC
	`

	rows, err := DB.Query(ctx, query, fmrID)
	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %w", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var row models.NoteRow
		err := rows.Scan(&row.ID, &row.FmrID, &row.Content, &row.Author, &row.NoteType, &row.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		notes = append(notes, *rowToNote(&row))
	}

	return &models.NotesListResponse{Notes: notes}, nil
}

func rowToNote(row *models.NoteRow) *models.Note {
	note := &models.Note{
		ID:        row.ID,
		FmrID:     row.FmrID,
		Content:   row.Content,
		NoteType:  models.NoteType(row.NoteType),
		CreatedAt: row.CreatedAt,
	}

	if row.Author.Valid {
		note.Author = &row.Author.String
	}

	return note
}
