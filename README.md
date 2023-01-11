# Architecture of a GoLang Rest API with JWT Authentication

## TECH STACK
- Go
- Posgresql
- Redis
- Minio/S3


## START PROJECT
### run docker compose
```console
sudo docker-compose up -d
```

### enable go modules
```console
export GO111MODULE="on"
```

### build wire
```console
go generate goblog.com/goblog/cmd/goblog
```

### run server
```console
go run cmd/goblog/wire_gen.go cmd/goblog/main.go
```

Visit the [Swagger docs](http://localhost:8080/api/docs/#/)
