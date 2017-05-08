package main

import (
  "github.com/rcrowley/go-metrics"
  "github.com/vrischmann/go-metrics-influxdb"
  "time"
  "math/rand"
  "os"
  "fmt"
)

func counterFunction(counter metrics.Counter, timeDelay time.Duration) {
  for {
    counter.Inc(1)
    random2 := time.Duration(rand.Intn(10)+1)
    random := time.Duration(rand.Float64())
    time.Sleep(random2 * ((timeDelay/2) + random*timeDelay))
  }
}

func main() {
  if (len(os.Args) < 5) {
    fmt.Println("Required arguments: influxdb url, databasename, dbUsername, dbPassword")
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

    c1 := metrics.NewCounter()
	  r.Register("song1", c1)
    c2 := metrics.NewCounter()
	  r.Register("song2", c2)
    c3 := metrics.NewCounter()
	  r.Register("song3", c3)
    c4 := metrics.NewCounter()
	  r.Register("song4", c4)
    c5 := metrics.NewCounter()
	  r.Register("song5", c5)

    g1 := metrics.NewGauge()
    r.Register("song11", g1)

    t1 := metrics.NewTimer()
    r.Register("song12", t1)

    go counterFunction(c1,500e7)
    go counterFunction(c2,640e7)
    go counterFunction(c3,800e7)
    go counterFunction(c4,1300e7)
    go counterFunction(c5,950e7)

    for {
        time1 := time.Now();
        count := g1.Value()
        random := int64(rand.Intn(50))
        if (count - 25) < random {
          g1.Update(count+1)
        } else {
          g1.Update(count-1)
        }
        time2 := time.Since(time1)
		    time.Sleep(500e6)
        t1.Update(time2)
    }
}
