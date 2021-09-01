How to run?
- import database `dbpokemon.sql`
- do `export DATA_POKEMON="{name}:{password}@tcp({host}:{port})/dbpokemon"` on terminal
- do `go run server.go` on terminal
- access http://localhost:5000/  ( you can set the port from `server.go` or set `PORT` env variable )
- call the schema
