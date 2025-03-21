build:
	GOOS=linux GOARCH=arm64 go build -o balance .

install:
	go install

bal:
	crypto-balance balance 0x25e3E139F9b6f52b91023566c981b007E2446518

token:
	crypto-balance token 0x047B2B563ec910F82536b6411bB9b96B9496A48c 0xAc720702Df63fa92416B3dEB24Dc4a1854f73330