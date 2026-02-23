package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"maxweise/sankey-script/internal"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	queries, err := internal.GetQueries(ctx, ":memory:")
	if err != nil {
		return err
	}

	entries, err := queries.GetAllEntries(ctx)
	if err != nil {
		return err
	}

	fmt.Println(entries)

	return nil
}
