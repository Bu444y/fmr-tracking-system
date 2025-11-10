package models

import (
	"encoding/json"
	"time"
)

// Filter Presets
type FilterPreset struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Filters   PresetFilters `json:"filters"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type PresetFilters struct {
	Search                *string `json:"search,omitempty"`
	StatusFilter          *string `json:"statusFilter,omitempty"`
	ProgramFilter         *string `json:"programFilter,omitempty"`
	UnitFilter            *string `json:"unitFilter,omitempty"`
	PointOfFailureFilter  *string `json:"pointOfFailureFilter,omitempty"`
	RepairActionFilter    *string `json:"repairActionFilter,omitempty"`
	TierLevelFilter       *string `json:"tierLevelFilter,omitempty"`
}

type CreatePresetRequest struct {
	Name    string        `json:"name" binding:"required"`
	Filters PresetFilters `json:"filters" binding:"required"`
}

type UpdatePresetRequest struct {
	ID      int64          `json:"id" binding:"required"`
	Name    *string        `json:"name,omitempty"`
	Filters *PresetFilters `json:"filters,omitempty"`
}

type ListPresetsResponse struct {
	Presets []FilterPreset `json:"presets"`
}

// Column Presets
type ColumnPreset struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Columns   []string  `json:"columns"`
	IsDefault bool      `json:"isDefault"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateColumnPresetRequest struct {
	Name      string   `json:"name" binding:"required"`
	Columns   []string `json:"columns" binding:"required"`
	IsDefault *bool    `json:"isDefault,omitempty"`
}

type UpdateColumnPresetRequest struct {
	ID        int64    `json:"id" binding:"required"`
	Name      *string  `json:"name,omitempty"`
	Columns   []string `json:"columns,omitempty"`
	IsDefault *bool    `json:"isDefault,omitempty"`
}

type ListColumnPresetsResponse struct {
	Presets []ColumnPreset `json:"presets"`
}

// Helper methods for JSON marshaling/unmarshaling
func (pf PresetFilters) MarshalJSON() ([]byte, error) {
	type Alias PresetFilters
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&pf),
	})
}
