package logger

import (
	"context"
	"log"

	"github.com/makxtr/go-common/pkg/db"
	"github.com/makxtr/go-common/pkg/logger/model"

	sq "github.com/Masterminds/squirrel"
)

const (
	actionColumn   = "action"
	entityIDColumn = "entity_id"
)

// Repository defines the interface for logging operations
type Repository interface {
	Log(ctx context.Context, logEntry *model.Log) error
}

type repository struct {
	db        db.Client
	tableName string
}

// NewRepository creates a new log repository with the specified table name
func NewRepository(dbClient db.Client, tableName string) Repository {
	return &repository{
		db:        dbClient,
		tableName: tableName,
	}
}

func (r *repository) Log(ctx context.Context, logEntry *model.Log) error {
	builder := sq.Insert(r.tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(actionColumn, entityIDColumn).
		Values(logEntry.Action, logEntry.EntityID)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	_, err = r.db.DB().ExecContext(ctx, db.Query{Name: "log_repository.Log", QueryRaw: query}, args...)
	if err != nil {
		log.Printf("failed to insert log: %v", err)
		return err
	}

	return nil
}
