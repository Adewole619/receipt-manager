package validation

import (
	"errors"
	"strings"

	"receipt-manager/internal/models"
)

// Validation: Store Name check and 🛠️ Receipt Number check
func ValidateReceipt(receipt models.Receipt) (string, string, error) {
	var errMessages []string

	cleanStoreName := strings.TrimSpace(receipt.StoreName)
	cleanReceiptNumber := strings.TrimSpace(receipt.ReceiptNumber)

	if cleanStoreName == "" {
		errMessages = append(errMessages, "- Store Name is required.")
	}

	if cleanReceiptNumber == "" {
		errMessages = append(errMessages, "- Receipt Number is required.")
	}

	if len(errMessages) > 0 {
		combinedMessage := strings.Join(errMessages, "\n")
		return cleanStoreName, cleanReceiptNumber, errors.New("validation Error:\n" + combinedMessage)
	}

	return cleanStoreName, cleanReceiptNumber, nil
}
