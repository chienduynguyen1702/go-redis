# Go client to redis
Simulate an inventory system that use for a sale campaign

## .env file
copy .env.example to .env and set your own values for the environment variables in it.
## Run Redis container
```
make redis
```
or
```
docker compose up --build -d
```
## Start program
### Install packages
```go
go mod tidy
```
### Run program
```
make start
```
or
```
go run cmd/server/main.go
```
## Try it now :v

```terminal
=====================================
1. Display all item in store
2. Regenerate items
3. Launch a sale
0. Exit
Enter option: ...
=====================================
```