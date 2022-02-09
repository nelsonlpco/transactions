clean:
	go clean --testcache

updatemocks:
	mockgen -destination=./domain/repository/mock/account_repository.go -source=./domain/repository/account_repository.go
	mockgen -destination=./domain/repository/mock/operatontype_repository.go -source=./domain/repository/operationtype_repository.go
	mockgen -destination=./domain/repository/mock/transaction_repository.go -source=./domain/repository/transaction_repository.go

test: clean
	go test -race -cover -timeout 30s ./...