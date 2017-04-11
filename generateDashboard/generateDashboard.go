/*
 * Originally by Komang
 * Found on https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
 */

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "text/template"
    "sort"
)

type Items struct {
    General GeneralSettings
    Graphs []Dataholder
    Tables []Dataholder
    Gauges []Gauge
}

type Items2 struct {
    General   GeneralSettings
    Items     []Item
    ItemCount int
}

type Item struct {
    Graph Dataholder
    Table Dataholder
    Gauge Gauge
    Selected int
}

type GeneralSettings struct {
    DatabaseName  string
    DashboardName  string
    TimeRange     string
    Refresh     string
}

type Dataholder struct {
    Title     string
    Width     int
    Row       int
    Column    int
    Variables []Variable
}

type Gauge struct {
    Title       string
    Width       int
    Row         int
    Column      int
    MinValue    int
    MaxValue    int
    Threshold1  int
    Threshold2  int
    Type        string
    FillColor   string
    LineColor   string
    Colors      []string
    Variables   []Variable
}

type Variable struct {
    Identifier   string
    Name         string
    ValueName    string
    Measurement  string
    Format       string
}

type ByRow []Item

func (a ByRow) Len() int           { return len(a) }
func (a ByRow) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRow) Less(i, j int) bool {
  if a[i].Selected == 0 && a[j].Selected == 0 {
    return a[i].Graph.Row < a[j].Graph.Row || (a[i].Graph.Row == a[j].Graph.Row && a[i].Graph.Column < a[j].Graph.Column)
  } else if a[i].Selected == 0 && a[j].Selected == 1 {
    return a[i].Graph.Row < a[j].Table.Row || (a[i].Graph.Row == a[j].Table.Row && a[i].Graph.Column < a[j].Table.Column)
  } else if a[i].Selected == 0 && a[j].Selected == 2 {
    return a[i].Graph.Row < a[j].Gauge.Row || (a[i].Graph.Row == a[j].Gauge.Row && a[i].Graph.Column < a[j].Gauge.Column)
  } else if a[i].Selected == 1 && a[j].Selected == 0 {
    return a[i].Table.Row < a[j].Graph.Row || (a[i].Table.Row == a[j].Graph.Row && a[i].Table.Column < a[j].Graph.Column)
  } else if a[i].Selected == 1 && a[j].Selected == 1 {
    return a[i].Table.Row < a[j].Table.Row || (a[i].Table.Row == a[j].Table.Row && a[i].Table.Column < a[j].Table.Column)
  } else if a[i].Selected == 1 && a[j].Selected == 2 {
    return a[i].Table.Row < a[j].Gauge.Row || (a[i].Table.Row == a[j].Gauge.Row && a[i].Table.Column < a[j].Gauge.Column)
  } else if a[i].Selected == 2 && a[j].Selected == 0 {
    return a[i].Gauge.Row < a[j].Graph.Row || (a[i].Gauge.Row == a[j].Graph.Row && a[i].Gauge.Column < a[j].Graph.Column)
  } else if a[i].Selected == 2 && a[j].Selected == 1 {
    return a[i].Gauge.Row < a[j].Table.Row || (a[i].Gauge.Row == a[j].Table.Row && a[i].Gauge.Column < a[j].Table.Column)
  } else if a[i].Selected == 2 && a[j].Selected == 2 {
    return a[i].Gauge.Row < a[j].Gauge.Row || (a[i].Gauge.Row == a[j].Gauge.Row && a[i].Gauge.Column < a[j].Gauge.Column)
  } else {
    return false
  }
}


func readDatabase(items Items) {
  fmt.Println("Database",items.General.DatabaseName)
}

func readGraphs(items []Item) {
  for key, item := range items { // key = index starting from 0
    fmt.Println("Key:", key)
    fmt.Println("Item:", item.Graph.Title)
    fmt.Println("Item:", item.Table.Title)
    fmt.Println("Item:", item.Gauge.Title)
    /*
    for _, item3 := range item2.Variables { // key = index starting from 0
      fmt.Println(item3.Name, item3.ValueName, item3.Measurement)
    }
    */
  }
}

func main() {
  items := getItemsFromFile()
  if (items.General.DatabaseName == "") {
    fmt.Println("No database name found")
    os.Exit(1)
  }
  //readDatabase(items)
  //readGraphs(items)
  parseTemplate(reOrganize(items))
}

func getItemsFromFile() Items {
    raw, err := ioutil.ReadFile("definitions.json")
    checkError(err)
    var items Items
    json.Unmarshal(raw, &items)
    return items
}

func reOrganize(items Items) Items2 {
  var items2 Items2
  items2.General = items.General
  a := len(items.Graphs)
  b := len(items.Tables)
  c := len(items.Gauges)
  items2.ItemCount =  a+b+c
  index := 0
  var arr = make([]Item, items2.ItemCount)
  for _, value := range items.Graphs {
    var item Item
    item.Graph = value
    item.Selected = 0
    arr[index] = item
    index += 1
  }
  for _, value := range items.Tables {
    var item Item
    item.Table = value
    item.Selected = 1
    arr[index] = item
    index += 1
  }
  for _, value := range items.Gauges {
    var item Item
    item.Gauge = value
    item.Selected = 2
    arr[index] = item
    index += 1
  }
  sort.Sort(ByRow(arr))
  items2.Items = arr
  return items2
}

func parseTemplate(items Items2) {
  template2 := readTemplateFile()
  t := template.New("template")
	t, err := t.Parse(template2)
  checkError(err)

  f, err := os.Create(items.General.DashboardName + ".json")
  checkError(err)
  err = t.Execute(f, items)
  checkError(err)
}

func readTemplateFile() string {
  dat, err := ioutil.ReadFile("template.tmpl")
  checkError(err)
  return string(dat)
}

func checkError(err error) {
  if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
