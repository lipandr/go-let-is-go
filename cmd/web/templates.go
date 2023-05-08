package main

import "github.com/lipandr/go-let-is-go/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
