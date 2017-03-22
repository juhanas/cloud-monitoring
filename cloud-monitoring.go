package main

import (
  "github.com/rcrowley/go-metrics"
  "github.com/vrischmann/go-metrics-influxdb"
  "time"
  "math/rand"
  "os"
)

func main() {
  if (len(os.Args) < 5) {
    return;
  }
    r := metrics.NewRegistry()

    go influxdb.InfluxDB(
        r, // metrics registry
        time.Second,  // interval
        os.Args[1],   // InfluxDB url from first command line argument
        os.Args[2],   // InfluxDB database
        os.Args[3],   // InfluxDB user
        os.Args[4],   // InfluxDB password
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
