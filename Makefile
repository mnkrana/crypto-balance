build:
	GOOS=linux GOARCH=arm64 go build -o balance ./crypto-balance

install:
	go install ./crypto-balance

bal1:
	crypto-balance 0x25e3E139F9b6f52b91023566c981b007E2446518