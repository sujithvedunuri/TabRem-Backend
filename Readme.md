# configure your dev enviroment
1. run docker-compose-shell script in terminal from infra/migration
2. Run go mod init Project-name/module-name
3. Run go mod tidy
4. Add .env File and SET env variables a. DB_USER b. DB_PASSWORD c. DB_NAME d. DB_HOST



# Run or Debug
go run main.go

# build or develop
go build