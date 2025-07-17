package qrcode

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(productID string) ([]byte, error) {
	url := fmt.Sprintf("https://localhost:8090/products/%s", productID)
	return qrcode.Encode(url, qrcode.Medium, 256)
}
