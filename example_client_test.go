package redis_timeseries_go_test

import (
	"fmt"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/gomodule/redigo/redis"
	"log"
)

// exemplifies the NewClientFromPool function
//nolint:errcheck
func ExampleNewClientFromPool() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")
	client.Add("ts", 1, 5)
	datapoints, _ := client.RangeWithOptions("ts", 0, 1000, redistimeseries.DefaultRangeOptions)
	fmt.Println(datapoints[0])
	// Output: {1 5}
}

// Exemplifies the usage of RangeWithOptions function
// nolint:errcheck
func ExampleClient_RangeWithOptions() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")
	for ts := 1; ts < 10; ts++ {
		client.Add("ts", int64(ts), float64(ts))
	}

	datapoints, _ := client.RangeWithOptions("ts", 0, 1000, redistimeseries.DefaultRangeOptions)
	fmt.Printf("Datapoints: %v\n", datapoints)
	// Output:
	// Datapoints: [{1 1} {2 2} {3 3} {4 4} {5 5} {6 6} {7 7} {8 8} {9 9}]
}

// nolint
// Exemplifies the usage of ReverseRangeWithOptions function
func ExampleClient_ReverseRangeWithOptions() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")
	for ts := 1; ts < 10; ts++ {
		client.Add("ts", int64(ts), float64(ts))
	}

	datapoints, _ := client.ReverseRangeWithOptions("ts", 0, 1000, redistimeseries.DefaultRangeOptions)
	fmt.Printf("Datapoints: %v\n", datapoints)
	// Output:
	// Datapoints: [{9 9} {8 8} {7 7} {6 6} {5 5} {4 4} {3 3} {2 2} {1 1}]
}

// nolint
// Exemplifies the usage of MultiRangeWithOptions function.
func ExampleClient_MultiRangeWithOptions() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")

	labels1 := map[string]string{
		"machine": "machine-1",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-1", 2, 1.0, redistimeseries.CreateOptions{Labels: labels1})
	client.Add("time-serie-1", 4, 2.0)

	labels2 := map[string]string{
		"machine": "machine-2",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-2", 1, 5.0, redistimeseries.CreateOptions{Labels: labels2})
	client.Add("time-serie-2", 4, 10.0)

	ranges, _ := client.MultiRangeWithOptions(1, 10, redistimeseries.DefaultMultiRangeOptions, "az=us-east-1")

	fmt.Printf("Ranges: %v\n", ranges)
	// Output:
	// Ranges: [{time-serie-1 map[] [{2 1} {4 2}]} {time-serie-2 map[] [{1 5} {4 10}]}]
}

// Exemplifies the usage of MultiReverseRangeWithOptions function.
// nolint:errcheck
func ExampleClient_MultiReverseRangeWithOptions() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")

	labels1 := map[string]string{
		"machine": "machine-1",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-1", 2, 1.0, redistimeseries.CreateOptions{Labels: labels1})
	client.Add("time-serie-1", 4, 2.0)

	labels2 := map[string]string{
		"machine": "machine-2",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-2", 1, 5.0, redistimeseries.CreateOptions{Labels: labels2})
	client.Add("time-serie-2", 4, 10.0)

	ranges, _ := client.MultiReverseRangeWithOptions(1, 10, redistimeseries.DefaultMultiRangeOptions, "az=us-east-1")

	fmt.Printf("Ranges: %v\n", ranges)
	// Output:
	// Ranges: [{time-serie-1 map[] [{4 2} {2 1}]} {time-serie-2 map[] [{4 10} {1 5}]}]
}

//nolint:errcheck
// Exemplifies the usage of MultiGetWithOptions function while using the default MultiGetOptions and while using user defined MultiGetOptions.
func ExampleClient_MultiGetWithOptions() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")

	labels1 := map[string]string{
		"machine": "machine-1",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-1", 2, 1.0, redistimeseries.CreateOptions{Labels: labels1})
	client.Add("time-serie-1", 4, 2.0)

	labels2 := map[string]string{
		"machine": "machine-2",
		"az":      "us-east-1",
	}
	client.AddWithOptions("time-serie-2", 1, 5.0, redistimeseries.CreateOptions{Labels: labels2})
	client.Add("time-serie-2", 4, 10.0)

	ranges, _ := client.MultiGetWithOptions(redistimeseries.DefaultMultiGetOptions, "az=us-east-1")

	rangesWithLabels, _ := client.MultiGetWithOptions(*redistimeseries.NewMultiGetOptions().SetWithLabels(true), "az=us-east-1")

	fmt.Printf("Ranges: %v\n", ranges)
	fmt.Printf("Ranges with labels: %v\n", rangesWithLabels)

	// Output:
	// Ranges: [{time-serie-1 map[] [{4 2}]} {time-serie-2 map[] [{4 10}]}]
	// Ranges with labels: [{time-serie-1 map[az:us-east-1 machine:machine-1] [{4 2}]} {time-serie-2 map[az:us-east-1 machine:machine-2] [{4 10}]}]
}

// Exemplifies the usage of MultiAdd.
func ExampleClient_MultiAdd() {
	host := "localhost:6379"
	password := ""
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}
	client := redistimeseries.NewClientFromPool(pool, "ts-client-1")

	labels1 := map[string]string{
		"machine": "machine-1",
		"az":      "us-east-1",
	}
	labels2 := map[string]string{
		"machine": "machine-2",
		"az":      "us-east-1",
	}

	err := client.CreateKeyWithOptions("timeserie-1", redistimeseries.CreateOptions{Labels: labels1})
	if err != nil {
		log.Fatal(err)
	}
	err = client.CreateKeyWithOptions("timeserie-2", redistimeseries.CreateOptions{Labels: labels2})
	if err != nil {
		log.Fatal(err)
	}

	// Adding multiple datapoints to multiple series
	datapoints := []redistimeseries.Sample{
		{"timeserie-1", redistimeseries.DataPoint{1, 10.5}},
		{"timeserie-1", redistimeseries.DataPoint{2, 40.5}},
		{"timeserie-2", redistimeseries.DataPoint{1, 60.5}},
	}
	timestamps, _ := client.MultiAdd(datapoints...)

	fmt.Printf("Example adding multiple datapoints to multiple series. Added timestamps: %v\n", timestamps)

	// Adding multiple datapoints to the same serie
	datapointsSameSerie := []redistimeseries.Sample{
		{"timeserie-1", redistimeseries.DataPoint{3, 10.5}},
		{"timeserie-1", redistimeseries.DataPoint{4, 40.5}},
		{"timeserie-1", redistimeseries.DataPoint{5, 60.5}},
	}
	timestampsSameSerie, _ := client.MultiAdd(datapointsSameSerie...)

	fmt.Printf("Example of adding multiple datapoints to the same serie. Added timestamps: %v\n", timestampsSameSerie)

	// Output:
	// Example adding multiple datapoints to multiple series. Added timestamps: [1 2 1]
	// Example of adding multiple datapoints to the same serie. Added timestamps: [3 4 5]
}