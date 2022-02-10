clean:
	go clean --testcache
	rm -rf coverage

test: clean
	go test -race -timeout 30s -covermode=atomic -cover ./...

coverage: clean
	mkdir coverage
	go test -race -timeout 30s -covermode=atomic -coverprofile coverage/c.out ./...
	go tool cover -html=coverage/c.out -o coverage/cover.html
	open coverage/cover.html

updatemocks:
	mockgen -destination=./domain/repository/mock/account_repository.go -source=./domain/repository/account_repository.go
	mockgen -destination=./domain/repository/mock/operatontype_repository.go -source=./domain/repository/operationtype_repository.go
	mockgen -destination=./domain/repository/mock/transaction_repository.go -source=./domain/repository/transaction_repository.go
	mockgen -destination=./infrastructure/datasource/mock/account.go -source=./infrastructure/datasource/account.go
	mockgen -destination=./infrastructure/datasource/mock/operationtype.go -source=./infrastructure/datasource/operationtype.go
	mockgen -destination=./infrastructure/datasource/mock/transaction.go -source=./infrastructure/datasource/transaction.go

build: clean test 
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(GOPATH)/bin/transactions main.go

buildimage: 
	docker build -f ./docker/Dockerfile -t nelsonlpco/transactions .