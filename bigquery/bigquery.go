// Sample bigquery-quickstart creates a Google BigQuery dataset.
package app

import (
	"fmt"
	"io"
	"log"
	"net/http"

	// Imports the Google Cloud BigQuery client package.
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func init() {

	registerHandlers()

}

func registerHandlers() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := testBigQuery()
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, res)
	}

}

func testBigQuery() (io.Writer, error) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "eastern-concord-176510"

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	q := client.Query(`
    SELECT year, SUM(number) as num
    FROM [bigquery-public-data:usa_names.usa_1910_2013]
    WHERE name = "William"
    GROUP BY year
    ORDER BY year
`)
	it, err := q.Read(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	var w io.Writer
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Fprint(w, values)
		fmt.Println(values)
	}

	return w, nil
}
