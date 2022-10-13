# Readme
## Packages
``` bash
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

go get github.com/joho/godotenv

go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

go get github.com/google/wire
go install github.com/google/wire/cmd/wire@latest

go get github.com/gin-contrib/cors
go get github.com/rs/cors

go get github.com/stretchr/testify
```

## Dependency injection
``` bash
wire gen dailyreport
```


## Run updated code automatically
``` bash
CompileDaemon -command="./pcharter"
```


## Unit Test
``` bash
go test -v ./...
go test -v
go test -v -run TestFunctionName
go test -v -run TestFunctionName/NamaSubTest
go test -run /NamaSubTest
```

## Docker 
``` bash
docker ps  # liat proses yang jalan
docker compose up --build
docker compose exec project-charter cat .env
docker compose exec project-charter /bin/sh
docker compose exec project-charter go run main.go
```

## error
Network change 172.26.0.0/16 to 172.28.0.0/16
```bash


```