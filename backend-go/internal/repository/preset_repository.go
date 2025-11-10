package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
)

type PresetRepository struct{}

func NewPresetRepository() *PresetRepository {
	return &PresetRepository{}
}

// Filter Presets
func (r *PresetRepository) CreateFilterPreset(ctx context.Context, req *models.CreatePresetRequest) (*models.FilterPreset, error) {
	filtersJSON, err := json.Marshal(req.Filters)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal filters: %w", err)
	}

	query := `
		INSERT INTO filter_presets (name, filters)
		VALUES ($1, $2)
		RETURNING id, name, filters, created_at, updated_at
	`

	var id int64
	var name string
	var filtersData []byte
	var createdAt, updatedAt time.Time

	err = DB.QueryRow(ctx, query, req.Name, filtersJSON).Scan(&id, &name, &filtersData, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create filter preset: %w", err)
	}

	var filters models.PresetFilters
	if err := json.Unmarshal(filtersData, &filters); err != nil {
		return nil, fmt.Errorf("failed to unmarshal filters: %w", err)
	}

	return &models.FilterPreset{
		ID:        id,
		Name:      name,
		Filters:   filters,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *PresetRepository) ListFilterPresets(ctx context.Context) (*models.ListPresetsResponse, error) {
	query := `SELECT id, name, filters, created_at, updated_at FROM filter_presets ORDER BY name`

	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list filter presets: %w", err)
	}
	defer rows.Close()

	var presets []models.FilterPreset
	for rows.Next() {
		var id int64
		var name string
		var filtersData []byte
		var createdAt, updatedAt time.Time

		if err := rows.Scan(&id, &name, &filtersData, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan filter preset: %w", err)
		}

		var filters models.PresetFilters
		if err := json.Unmarshal(filtersData, &filters); err != nil {
			return nil, fmt.Errorf("failed to unmarshal filters: %w", err)
		}

		presets = append(presets, models.FilterPreset{
			ID:        id,
			Name:      name,
			Filters:   filters,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return &models.ListPresetsResponse{Presets: presets}, nil
}

func (r *PresetRepository) DeleteFilterPreset(ctx context.Context, id int64) error {
	query := "DELETE FROM filter_presets WHERE id = $1"
	result, err := DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete filter preset: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("filter preset not found")
	}

	return nil
}

// Column Presets
func (r *PresetRepository) CreateColumnPreset(ctx context.Context, req *models.CreateColumnPresetRequest) (*models.ColumnPreset, error) {
	isDefault := false
	if req.IsDefault != nil {
		isDefault = *req.IsDefault
	}

	columnsJSON, err := json.Marshal(req.Columns)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal columns: %w", err)
	}

	query := `
		INSERT INTO column_presets (name, columns, is_default)
		VALUES ($1, $2, $3)
		RETURNING id, name, columns, is_default, created_at, updated_at
	`

	var id int64
	var name string
	var columnsData []byte
	var isDefaultResult bool
	var createdAt, updatedAt time.Time

	err = DB.QueryRow(ctx, query, req.Name, columnsJSON, isDefault).Scan(
		&id, &name, &columnsData, &isDefaultResult, &createdAt, &updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create column preset: %w", err)
	}

	var columns []string
	if err := json.Unmarshal(columnsData, &columns); err != nil {
		return nil, fmt.Errorf("failed to unmarshal columns: %w", err)
	}

	return &models.ColumnPreset{
		ID:        id,
		Name:      name,
		Columns:   columns,
		IsDefault: isDefaultResult,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *PresetRepository) ListColumnPresets(ctx context.Context) (*models.ListColumnPresetsResponse, error) {
	query := `SELECT id, name, columns, is_default, created_at, updated_at FROM column_presets ORDER BY name`

	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list column presets: %w", err)
	}
	defer rows.Close()

	var presets []models.ColumnPreset
	for rows.Next() {
		var id int64
		var name string
		var columnsData []byte
		var isDefault bool
		var createdAt, updatedAt time.Time

		if err := rows.Scan(&id, &name, &columnsData, &isDefault, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan column preset: %w", err)
		}

		var columns []string
		if err := json.Unmarshal(columnsData, &columns); err != nil {
			return nil, fmt.Errorf("failed to unmarshal columns: %w", err)
		}

		presets = append(presets, models.ColumnPreset{
			ID:        id,
			Name:      name,
			Columns:   columns,
			IsDefault: isDefault,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return &models.ListColumnPresetsResponse{Presets: presets}, nil
}

func (r *PresetRepository) DeleteColumnPreset(ctx context.Context, id int64) error {
	query := "DELETE FROM column_presets WHERE id = $1"
	result, err := DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete column preset: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("column preset not found")
	}

	return nil
}
