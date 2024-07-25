package main

import (
	"log/slog"

	"snippetbox.elmm.net/internal/models"
)

type Application struct {
	Logger   *slog.Logger
	snippets *models.SnippetModel
}
