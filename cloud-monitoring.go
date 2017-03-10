package main

import (
  "github.com/rcrowley/go-metrics"
  "github.com/vrischmann/go-metrics-influxdb"
  "time"
  "math/rand"
)

func main() {
    r := metrics.NewRegistry()

    go influxdb.InfluxDB(
        r, // metrics registry
        time.Second,        // interval
        "http://192.168.99.100:30955", // the InfluxDB url ** Insert the url here **
        "data",                  // your InfluxDB database
        "root",                // your InfluxDB user
        "root",            // your InfluxDB password
    )

    c := metrics.NewCounter()
	  r.Register("measurement", c)
	  for {
        count := c.Count()
        random := int64(rand.Intn(50))
        if (count - 25) < random {
          c.Inc(1)
        } else {
          c.Dec(1)
        }
		    time.Sleep(100e6)
    }
}
