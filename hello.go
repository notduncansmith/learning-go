package main

import (
	"github.com/go-martini/martini"
)

type Database map[string]string

func main() {
	m := martini.Classic()
	db := Database{}

	m.Map(db)

	routes(m)

	m.Run()
}

func routes(m *martini.ClassicMartini) {

	m.Get("/", func() string {
		return "Hello World!"
	})

	m.Get("/store/:key/:value", func(p martini.Params, db Database) string {
		key := p["key"]
		value := p["value"]
		db[key] = value

		return "Stored " + value + " in " + key
	})

	m.Get("/get/:key", func(p martini.Params, db Database) string {
		key := p["key"]
		value := db[key]
		return value
	})

	m.Get("/hello/:name", func(p martini.Params) string {
		return "Hello " + p["name"]
	})

}
