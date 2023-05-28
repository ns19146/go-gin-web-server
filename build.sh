#!/usr/bin/env bash
go build -o ./app main.go
go build -o ./app/controller route.go
