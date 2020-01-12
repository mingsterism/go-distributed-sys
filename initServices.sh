#!/bin/bash

go run ./paymentservice/main.go &> ./paymentservice/logs.txt &
go run ./orderservice/main.go &> ./orderservice/logs.txt &
go run ./restaurantservice/main.go  &> ./restaurantservice/logs.txt &
go run ./eventstore/main.go &> ./eventstore/logs.txt &
go run ./orderquery-store1/main.go &> ./orderquery-store1/logs.txt &
go run ./orderquery-store2/main.go &> ./orderquery-store2/logs.txt &
