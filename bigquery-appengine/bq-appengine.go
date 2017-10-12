// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command simpleapp queries the Shakespeare sample dataset in Google BigQuery.
package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
	//"context"
)

func init() {

	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	rows, err := query(proj, c)
	if err != nil {
		log.Fatal(err)
	}
	if err := printResults(&w, rows); err != nil {
		log.Fatal(err)
	}
}

// query returns a slice of the results of a query.
func query(proj string, ctx context.Context) (*bigquery.RowIterator, error) {

	client, err := bigquery.NewClient(ctx, proj)
	if err != nil {
		return nil, err
	}

	query := client.Query(
		`SELECT * FROM ` +
			"`bigquery-public-data.cloud_storage_geo_index.sentinel_2_index`" +
			` WHERE west_lon = 32.0470531263` +
			` LIMIT 10;`)
	// Use standard SQL syntax for queries.
	// See: https://cloud.google.com/bigquery/sql-reference/
	query.QueryConfig.UseStandardSQL = true
	return query.Read(ctx)
}

// printResults prints results from a query to the Shakespeare dataset.
func printResults(w_p *http.ResponseWriter, iter *bigquery.RowIterator) error {
	for {
		var row []bigquery.Value
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(*w_p, row)
		fmt.Println(row)
	}
}
