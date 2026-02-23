package internal

import (
	"context"
	"database/sql"
	_ "embed"
	"maxweise/sankey-script/data_acess"
)

// A service object that bundles data necessary for the execution of queries.
type CoreService struct {
	repository *data_acess.Queries
	ctx        context.Context
}

// Check the entry according to the requirements and create a new db entry if possible
func (s CoreService) CreateEntry(source string, target string, amount float32, description string) (data_acess.Entry, error) {
	var desc sql.NullString
	if description == "" {
		desc = sql.NullString{String: "", Valid: false}
	} else {
		desc = sql.NullString{String: description, Valid: true}
	}
	arg := data_acess.CreateEntryParams{
		Source:      source,
		Target:      target,
		Amount:      amount,
		Description: desc,
	}
	o, err := s.repository.CreateEntry(s.ctx, arg)

	if err != nil {
		return data_acess.Entry{}, err
	}

	return o, nil
}

// Read all entries from the database
func (s CoreService) ReadAllEntries() ([]data_acess.Entry, error) {
	return s.repository.GetAllEntries(s.ctx)
}

//go:embed schema.sql
var schema string

// Create a query object which can execute database manipulations.
func GetQueries(ctx context.Context, connection string) (*data_acess.Queries, error) {
	// inintialize connections
	db, err := sql.Open("sqlite", connection)
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	return data_acess.New(db), err

}
