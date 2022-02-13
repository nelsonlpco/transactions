clean:
	rm -rf ./*.db
	go clean --testcache
	rm -rf coverage

test: clean
	go test -race -timeout 30s -covermode=atomic -cover ./shared/commonerrors ./domain/entity ./domain/usecases ./domain/valueobjects ./infrastructure/datasource ./infrastructure/inframodel ./infrastructure/repository  ./application/services

coverage: clean
	mkdir coverage
	go test -race -timeout 30s -covermode=atomic -coverprofile coverage/c.out ./shared/commonerrors ./domain/entity ./domain/usecases ./domain/valueobjects ./infrastructure/datasource ./infrastructure/inframodel ./infrastructure/repository  ./application/services

	go tool cover -html=coverage/c.out -o coverage/cover.html
	open coverage/cover.html

updatemocks:
	mockgen -destination=./domain/repository/mock/account_repository.go -source=./domain/repository/account_repository.go
	mockgen -destination=./domain/repository/mock/operatontype_repository.go -source=./domain/repository/operationtype_repository.go
	mockgen -destination=./domain/repository/mock/transaction_repository.go -source=./domain/repository/transaction_repository.go
	mockgen -destination=./infrastructure/datasource/mock/account_datasource.go -source=./infrastructure/datasource/account_datasource.go
	mockgen -destination=./infrastructure/datasource/mock/operationtype_datasource.go -source=./infrastructure/datasource/operationtype_datasource.go
	mockgen -destination=./infrastructure/datasource/mock/transaction_datasource.go -source=./infrastructure/datasource/transaction_datasource.go

build:  
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(GOPATH)/bin/transactions main.go

buildimage: 
	docker build -f ./docker/Dockerfile -t nelsonlpco/transactions:prod .

rundev: clean
	ENVFILE=.env-dev go run ./main.go

runtest: clean
	ENVFILE=.env-test go run ./main.go -d -s