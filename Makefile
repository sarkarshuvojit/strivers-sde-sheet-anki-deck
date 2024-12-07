default:
	@echo 'Hello'

download:
	@go run cmd/download-snapshot/main.go

create-deck:
	@go run cmd/create-deck/main.go --create-dir
