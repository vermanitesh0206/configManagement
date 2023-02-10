# My-App

My-App is a golang application for dealing with configuration management.

## Pre-requisites

Golang [steps to install](https://go.dev/doc/install)

Mongo DB Atlas [steps for free account](https://www.mongodb.com/atlas/database)

Postman [steps to install](https://www.postman.com/downloads/)

## Routes

```python
login
# sample request
POST http://localhost:3030/login/
BODY
{
    "Username": "<enter username>",
    "Password": "<enter password>"
}
```
```python
add/update
# sample request
POST http://localhost:3030/add/
PUT http://localhost:3030/update/
TOKEN <your-jwt-token>
BODY
{
    "Key": "<enter config name>",
    "Value": "<enter config value>"
}
```
```python
search/delete
# sample request
GET http://localhost:3030/search/<configname>
DELETE http://localhost:3030/delete/<configname>
TOKEN <your-jwt-token>
```

## Steps to Run

```bash
cd cmd/main
go run main.go
```
Once the server is up and running, use postman to test the Apis.