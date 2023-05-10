build:
	@go build -o cronjob

run: build
	@./cronjob