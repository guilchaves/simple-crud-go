## Simple GO REST API - GIN

A simple CRUD application written in Golang and Gin framework using clean architecture.

#### How to run
```docker
docker build -t go-api-gin .
```
```bash
docker compose up -d
```

The application will serve on ```127.0.0.1:8080```
#### Endpoints:

_GetProducts_
```bash
GET /products
```
_GetById_
```bash
GET /product/:id
```
_CreateProduct_
```bash
POST  /product
```

#### TODO

- Add slog
- Add swagger
- Add delete and update
