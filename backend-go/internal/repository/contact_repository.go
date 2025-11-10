package repository

import (
	"context"
	"fmt"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
)

type ContactRepository struct{}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func (r *ContactRepository) Create(ctx context.Context, req *models.CreateContactRequest) (*models.Contact, error) {
	query := `
		INSERT INTO contacts (title, sequence_number, first_name, last_name, email, mobile_number, work_number, organization, notes, team)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, title, sequence_number, first_name, last_name, email, mobile_number, work_number, organization, notes, team, created_at, updated_at
	`

	var row models.ContactRow
	err := DB.QueryRow(ctx, query,
		req.Title, req.SequenceNumber, req.FirstName, req.LastName, req.Email,
		req.MobileNumber, req.WorkNumber, req.Organization, req.Notes, req.Team,
	).Scan(
		&row.ID, &row.Title, &row.SequenceNumber, &row.FirstName, &row.LastName,
		&row.Email, &row.MobileNumber, &row.WorkNumber, &row.Organization,
		&row.Notes, &row.Team, &row.CreatedAt, &row.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create contact: %w", err)
	}

	return rowToContact(&row), nil
}

func (r *ContactRepository) List(ctx context.Context) (*models.ContactsListResponse, error) {
	query := `
		SELECT id, title, sequence_number, first_name, last_name, email, mobile_number, work_number, organization, notes, team, created_at, updated_at
		FROM contacts
		ORDER BY last_name, first_name
	`

	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list contacts: %w", err)
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var row models.ContactRow
		err := rows.Scan(
			&row.ID, &row.Title, &row.SequenceNumber, &row.FirstName, &row.LastName,
			&row.Email, &row.MobileNumber, &row.WorkNumber, &row.Organization,
			&row.Notes, &row.Team, &row.CreatedAt, &row.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan contact: %w", err)
		}
		contacts = append(contacts, *rowToContact(&row))
	}

	var total int
	if err := DB.QueryRow(ctx, "SELECT COUNT(*) FROM contacts").Scan(&total); err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	return &models.ContactsListResponse{
		Contacts: contacts,
		Total:    total,
	}, nil
}

func (r *ContactRepository) Update(ctx context.Context, req *models.UpdateContactRequest) (*models.Contact, error) {
	query := `
		UPDATE contacts
		SET title = COALESCE($1, title),
		    sequence_number = COALESCE($2, sequence_number),
		    first_name = COALESCE($3, first_name),
		    last_name = COALESCE($4, last_name),
		    email = COALESCE($5, email),
		    mobile_number = COALESCE($6, mobile_number),
		    work_number = COALESCE($7, work_number),
		    organization = COALESCE($8, organization),
		    notes = COALESCE($9, notes),
		    team = COALESCE($10, team),
		    updated_at = NOW()
		WHERE id = $11
		RETURNING id, title, sequence_number, first_name, last_name, email, mobile_number, work_number, organization, notes, team, created_at, updated_at
	`

	var row models.ContactRow
	err := DB.QueryRow(ctx, query,
		req.Title, req.SequenceNumber, req.FirstName, req.LastName, req.Email,
		req.MobileNumber, req.WorkNumber, req.Organization, req.Notes, req.Team, req.ID,
	).Scan(
		&row.ID, &row.Title, &row.SequenceNumber, &row.FirstName, &row.LastName,
		&row.Email, &row.MobileNumber, &row.WorkNumber, &row.Organization,
		&row.Notes, &row.Team, &row.CreatedAt, &row.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update contact: %w", err)
	}

	return rowToContact(&row), nil
}

func (r *ContactRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM contacts WHERE id = $1"
	result, err := DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete contact: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("contact not found")
	}

	return nil
}

func rowToContact(row *models.ContactRow) *models.Contact {
	contact := &models.Contact{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Team:      models.Team(row.Team),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	if row.Title.Valid {
		contact.Title = &row.Title.String
	}
	if row.SequenceNumber.Valid {
		num := int(row.SequenceNumber.Int64)
		contact.SequenceNumber = &num
	}
	if row.Email.Valid {
		contact.Email = &row.Email.String
	}
	if row.MobileNumber.Valid {
		contact.MobileNumber = &row.MobileNumber.String
	}
	if row.WorkNumber.Valid {
		contact.WorkNumber = &row.WorkNumber.String
	}
	if row.Organization.Valid {
		contact.Organization = &row.Organization.String
	}
	if row.Notes.Valid {
		contact.Notes = &row.Notes.String
	}

	return contact
}
