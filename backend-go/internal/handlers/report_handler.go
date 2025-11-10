package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Bu444y/fmr-tracking-system/internal/repository"
)

type ReportHandler struct{}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{}
}

type ExportResponse struct {
	Data        string `json:"data"`
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
}

func (h *ReportHandler) Export(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	query := `
		SELECT
			id,
			title,
			description,
			status,
			creation_date,
			resolution_date,
			created_at,
			updated_at
		FROM fmr_requests
	`

	args := []interface{}{}
	if startDate != "" && endDate != "" {
		query += " WHERE creation_date >= $1 AND creation_date <= $2"
		args = append(args, startDate, endDate)
	}

	query += " ORDER BY creation_date DESC"

	rows, err := repository.DB.Query(r.Context(), query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Build CSV
	headers := []string{"ID", "Title", "Description", "Status", "Creation Date", "Resolution Date", "Created At", "Updated At"}
	csvRows := []string{strings.Join(headers, ",")}

	for rows.Next() {
		var id int64
		var title, status string
		var description *string
		var creationDate, createdAt, updatedAt time.Time
		var resolutionDate *time.Time

		if err := rows.Scan(&id, &title, &description, &status, &creationDate, &resolutionDate, &createdAt, &updatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		desc := ""
		if description != nil {
			desc = strings.ReplaceAll(*description, "\"", "\"\"")
		}

		resDate := ""
		if resolutionDate != nil {
			resDate = resolutionDate.Format(time.RFC3339)
		}

		row := []string{
			fmt.Sprintf("%d", id),
			fmt.Sprintf("\"%s\"", strings.ReplaceAll(title, "\"", "\"\"")),
			fmt.Sprintf("\"%s\"", desc),
			status,
			creationDate.Format(time.RFC3339),
			resDate,
			createdAt.Format(time.RFC3339),
			updatedAt.Format(time.RFC3339),
		}

		csvRows = append(csvRows, strings.Join(row, ","))
	}

	response := ExportResponse{
		Data:        strings.Join(csvRows, "\n"),
		Filename:    fmt.Sprintf("fmr_export_%s.csv", time.Now().Format("2006-01-02")),
		ContentType: "text/csv",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
