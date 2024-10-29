package datastore

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDB2 struct {
	bucket           string
	org              string
	token            string
	client           influxdb2.Client
	writeAPI         api.WriteAPI
	blockingWriteAPI api.WriteAPIBlocking
	queryAPI         api.QueryAPI
}

func loadAPIToken() string {
	file, err := os.Open("/var/lib/shared_config/influxdb_config.txt")
	if err != nil {
		fmt.Println("File not found for loading API token")
		return ""
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error encountered while reading API token file, err: %s", err.Error())
		return ""
	}
	return strings.TrimSpace(string(data))
}

func NewInfluxDB2() InfluxDataStore {
	bucket := os.Getenv("DOCKER_INFLUXDB_INIT_BUCKET")
	org := os.Getenv("DOCKER_INFLUXDB_INIT_ORG")
	token := loadAPIToken()
	const RETRY_COUNT = 20
	for retry := 1; retry <= RETRY_COUNT; retry++ {
		if token != "" {
			break
		}
		fmt.Printf("Token not found for influxDB, waiting... for %d seconds\n", 2*retry)
		time.Sleep(time.Duration(2*retry) * time.Second)
		token = loadAPIToken()
	}
	url := "http://influxdb2:8086"
	// Batch size would increase in real application, should change to higher numbers
	client := influxdb2.NewClientWithOptions(url, token, influxdb2.DefaultOptions().SetBatchSize(1))
	writeAPI := client.WriteAPI(org, bucket)
	blockingWriteAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)
	return &InfluxDB2{
		bucket:           bucket,
		org:              org,
		token:            token,
		client:           client,
		writeAPI:         writeAPI,
		blockingWriteAPI: blockingWriteAPI,
		queryAPI:         queryAPI,
	}
}

func (db *InfluxDB2) GetLastValue(ctx context.Context, measurement string, tags map[string]string) (interface{}, error) {
	tagQuery := ""
	for key, val := range tags {
		if key != "" && val != "" {
			tagQuery += fmt.Sprintf(` and r.%s == "%s"`, key, val)
		}
	}
	result, err := db.queryAPI.Query(ctx,
		fmt.Sprintf(`from(bucket:"%s")
		|> range(start:-1d)
		|> filter(fn: (r) => r._measurement == "%s"%s)
		|> last()
		`, db.bucket, measurement, tagQuery),
	)
	if err != nil {
		return result, err
	}
	var value interface{}
	for result.Next() {
		value = result.Record().Value()
	}
	if err := result.Err(); err != nil {
		return value, err
	}
	return value, nil
}

func (db *InfluxDB2) Get(ctx context.Context, measurement string) ([]interface{}, error) {
	result, err := db.queryAPI.Query(ctx,
		fmt.Sprintf(`from(bucket:"%s")|>range(start:-1h) |> filter(fn: (r) => r._measurement == "%s")`, db.bucket, measurement))
	if err != nil {
		return []interface{}{}, err
	}
	results := []interface{}{}
	for result.Next() {
		results = append(results, result.Record().Value())
	}
	if err := result.Err(); err != nil {
		return []interface{}{}, err
	}
	return results, nil
}

func (db *InfluxDB2) WriteSync(ctx context.Context, key string, tags map[string]string, fields map[string]interface{}) error {
	point := influxdb2.NewPoint(key, tags, fields, time.Now())
	if err := db.blockingWriteAPI.WritePoint(ctx, point); err != nil {
		return err
	}
	return nil
}

// Writes asynchronously to the database
func (db *InfluxDB2) WriteAsync(ctx context.Context, key string, tags map[string]string, fields map[string]interface{}) {
	point := influxdb2.NewPoint(key, tags, fields, time.Now())
	db.writeAPI.WritePoint(point)
}
