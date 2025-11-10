package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/Bu444y/fmr-tracking-system/internal/handlers"
	"github.com/Bu444y/fmr-tracking-system/internal/middleware"
	"github.com/Bu444y/fmr-tracking-system/internal/repository"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Check if FIPS mode is enabled
	log.Println("Starting FMR Tracking System Backend (FIPS-Enabled)")
	log.Printf("Go Version: %s", runtime.Version())
	log.Printf("GODEBUG: %s", os.Getenv("GODEBUG"))

	// Initialize database connection
	if err := repository.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer repository.CloseDB()

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.CORS)
	r.Use(middleware.Logger)

	// Initialize handlers
	fmrHandler := handlers.NewFMRHandler()
	contactHandler := handlers.NewContactHandler()
	noteHandler := handlers.NewNoteHandler()
	presetHandler := handlers.NewPresetHandler()
	reportHandler := handlers.NewReportHandler()

	// FMR routes (matching Encore API paths)
	r.Post("/fmr", fmrHandler.Create)
	r.Get("/fmr", fmrHandler.List)
	r.Get("/fmr/{id}", fmrHandler.Get)
	r.Put("/fmr/{id}", fmrHandler.Update)

	// Contact routes
	r.Post("/contacts", contactHandler.Create)
	r.Get("/contacts", contactHandler.List)
	r.Put("/contacts/{id}", contactHandler.Update)
	r.Delete("/contacts/{id}", contactHandler.Delete)

	// Note routes
	r.Post("/notes", noteHandler.Create)
	r.Get("/notes/fmr/{fmrId}", noteHandler.List)

	// Filter Preset routes
	r.Post("/presets", presetHandler.CreateFilterPreset)
	r.Get("/presets", presetHandler.ListFilterPresets)
	r.Delete("/presets/{id}", presetHandler.DeleteFilterPreset)

	// Column Preset routes
	r.Post("/column-presets", presetHandler.CreateColumnPreset)
	r.Get("/column-presets", presetHandler.ListColumnPresets)
	r.Delete("/column-presets/{id}", presetHandler.DeleteColumnPreset)

	// Report routes
	r.Get("/reports/export", reportHandler.Export)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting on %s", addr)
	log.Printf("FIPS Mode: Check runtime logs for crypto/internal/fips140 initialization")

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
