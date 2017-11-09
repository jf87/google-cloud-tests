// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command simpleapp queries the Sentinel dataset in Google BigQuery.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"

	"context"
)

func main() {
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	rows, err := query(proj)
	if err != nil {
		log.Fatal(err)
	}
	if err := printResults(os.Stdout, rows); err != nil {
		log.Fatal(err)
	}
}

// query returns a slice of the results of a query.
func query(proj string) (*bigquery.RowIterator, error) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, proj)
	if err != nil {
		return nil, err
	}

	query := client.Query(
		`SELECT * FROM ` +
			"`bigquery-public-data.cloud_storage_geo_index.sentinel_2_index`" +
			` WHERE west_lon < 32.0470531263 AND west_lon > 32.0465531263` +
			` LIMIT 10;`)
	// Use standard SQL syntax for queries.
	// See: https://cloud.google.com/bigquery/sql-reference/
	query.QueryConfig.UseStandardSQL = true
	return query.Read(ctx)
}

func printResults(w io.Writer, iter *bigquery.RowIterator) error {
	for {
		var row []bigquery.Value
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(row)
	}
}
