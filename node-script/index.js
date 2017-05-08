"use strict";

var InfluxMetrics = require('metrics-influxdb');

if (process.argv.length < 7) {
  console.log("Required arguments: influxdb host url, influxdb host port, databasename, dbUsername, dbPassword");
  return;
}

var reporter = new InfluxMetrics.Reporter({
  host: process.argv[2],
  port: process.argv[3],
  protocol: 'http',
  database: process.argv[4],
  username: process.argv[5],
  password: process.argv[6],
  tags: {
    'server': 'one',
  },
  skipIdleMetrics: false,
  precision: 'ms'
});

function counterFunction() {
  var multiplier = 100000;
  var rand = Math.floor((Math.random() * 10) + 1);
  for (var i = 1; i >= 0; i++) {
    if (i%((50+(5*rand))*multiplier) == 0) {
      c1.inc();
    }
    if (i%((80+(8*rand))*multiplier) == 0) {
      c2.inc();
    }
    if (i%((120+(12*rand))*multiplier) == 0) {
      c3.inc();
    }
    if (i%((150+(15*rand))*multiplier) == 0) {
      c4.inc();
    }
    if (i%((170+(17*rand))*multiplier) == 0) {
      c5.inc();
    }
    if (i%(500*multiplier) == 0) {
      break;
    }
  }
}

function gaugeFunction() {
  var multiplier = 100000;
  for (var i = 1; i > 0; i++) {
    if (i%(50*multiplier) == 0) {
      var count = g.latest();
      var rand = Math.floor((Math.random() * 50) + 1);
      if (count-25 < rand) {
        var newCount = count+1;
        g.set(newCount);
      } else {
        var newCount = count-1;
        g.set(newCount);
      }
    }
    if (i%(500*multiplier) == 0) {
      break;
    }
  }
}

function histogramFunction() {
  var multiplier = 100000;
  for (var i = 1; i > 0; i++) {
    if (i%(50*multiplier) == 0) {
      var rand = Math.floor((Math.random() * 100) + 1);
      h.update(rand);
    }
    if (i%(500*multiplier) == 0) {
      break;
    }
  }
}

function meterFunction() {
  var multiplier = 100000;
  var j = 1;
  var rand = Math.floor((Math.random() * 200) + 1);
  for (var i = 1; i > 0; i++) {
    if (i%(rand*j*multiplier) == 0) {
      m.mark();
      j++;
      rand = Math.floor((Math.random() * 200) + 1);
    }
    if (i%(500*multiplier) == 0) {
      break;
    }
  }
}

function timerFunction() {
  var multiplier = 100000;
  var j = 1;
  var rand = Math.floor((Math.random() * 200) + 1);
  for (var i = 1; i > 0; i++) {
    if (i%(rand*j*multiplier) == 0) {
      t.update(rand/2);
      j++;
      rand = Math.floor((Math.random() * 200) + 1);
    }
    if (i%(500*multiplier) == 0) {
      break;
    }
  }
}

var c1 = new InfluxMetrics.Counter();
reporter.addMetric('song1.count', c1);
var c2 = new InfluxMetrics.Counter();
reporter.addMetric('song2.count', c2);
var c3 = new InfluxMetrics.Counter();
reporter.addMetric('song3.count', c3);
var c4 = new InfluxMetrics.Counter();
reporter.addMetric('song4.count', c4);
var c5 = new InfluxMetrics.Counter();
reporter.addMetric('song5.count', c5);

var g = new InfluxMetrics.Gauge();
reporter.addMetric('song11.gauge', g);

var h = new InfluxMetrics.Histogram();
reporter.addMetric('song12.histogram', h);

var m = new InfluxMetrics.Meter();
reporter.addMetric('song13.meter', m);

var t = new InfluxMetrics.Timer();
reporter.addMetric('song14.timer', t);

counterFunction();
gaugeFunction();
histogramFunction();
meterFunction();
timerFunction();
setInterval(reporter.report.bind(reporter, true), 4900);
setInterval(counterFunction, 5000);
setInterval(gaugeFunction, 5000);
setInterval(histogramFunction, 5000);
setInterval(meterFunction, 5000);
setInterval(timerFunction, 5000);
