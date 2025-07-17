go get -u gorm.io/gorm

go get -u github.com/gin-gonic/gin

go get github.com/ethereum/go-ethereum

go get github.com/ethereum/go-ethereum/rpc@latest

go get github.com/ethereum/go-ethereum/accounts/keystore@latest

abigen --abi=internal/contracts/Traceability.abi.json --pkg=contracts --out=internal/contracts/contract.go

go get -u github.com/skip2/go-qrcode/...

go get github.com/spf13/viper@latest

go get gorm.io/driver/mysql
