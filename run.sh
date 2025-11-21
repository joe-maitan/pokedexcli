#!/bin/bash

go clean
go build .
./pokedexcli
go clean