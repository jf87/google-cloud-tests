package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {

	var s = "gcp-public-data-sentinel-2"
	list(s)
}

func list(bucket string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	query := &storage.Query{Prefix: "tiles/32/D/PG/S2A_MSIL1C_20171014T073921_N0205_R063_T32DPG_20171014T073917.SAFE/GRANULE/L1C_T32DPG_A012072_20171014T073917/IMG_DATA"}
	it := client.Bucket(bucket).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(attrs.Name)
	}
}
