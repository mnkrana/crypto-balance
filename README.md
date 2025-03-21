## Steps to get started

This works based on the `.env` file contents.
To switch the chain, change the network and RPC in the `.env`

### Install package

```go
go install github.com/mnkrana/crypto-balance@latest
```

### Add .env to your GOPATH

```sh
touch $GOPATH/bin/.env
```

### Environment variables

Check env-example file to add environment variables

### Usage example

- Get native ERC20 balance

```sh
crypto-balance balance <wallet or contract address>
```

- Get any ERC20 balance

```sh
crypto-balance token <wallet or contract address> <ECR20 contract address>
```

- Transfer crypto

```sh
crypto-balance transfer <private key of sender> <receiver address> <amount in wei>
```
