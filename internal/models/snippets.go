package models

import (
	"database/sql"
	"time"
)

type Snippets struct {
	ID      int
	Title   string
	Content string
	Create  time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}
