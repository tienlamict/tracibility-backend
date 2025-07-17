go get -u gorm.io/gorm

go get -u github.com/gin-gonic/gin

go get github.com/ethereum/go-ethereum

go get github.com/ethereum/go-ethereum/rpc@latest

go get github.com/ethereum/go-ethereum/accounts/keystore@latest

abigen --abi=internal/contracts/Traceability.abi.json --pkg=contracts --out=internal/contracts/contract.go

go get -u github.com/skip2/go-qrcode/...

go get github.com/spf13/viper@latest

go get gorm.io/driver/mysql


| Thành phần     | Vai trò đề xuất                                  |
| -------------- | ------------------------------------------------ |
| **Frontend**   | Vẫn dùng MetaMask để ký tx.                      |
| **Backend Go** | Thêm:                                            |
|                | - Kiểm tra txHash trạng thái (pending, fail...). |
|                | - Lưu lịch sử truy xuất.                         |
|                | - Tạo QR code từ product ID.                     |
|                | - Tự động crawl event từ contract và lưu DB.     |
|                | - Tạo dashboard thống kê.                        |
