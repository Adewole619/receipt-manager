package validation

import (
	"errors"
	"strings"

	"receipt-manager/internal/models"
)

// Validation: Store Name check and 🛠️ Receipt Number check
func ValidateReceipt(receipt models.Receipt) (models.Receipt, error) {
	var errMessages []string

	receipt.StoreName = strings.TrimSpace(receipt.StoreName)
	receipt.ReceiptNumber = strings.TrimSpace(receipt.ReceiptNumber)

	if receipt.StoreName == "" {
		errMessages = append(errMessages, "- Store Name is required.")
	}

	if receipt.ReceiptNumber == "" {
		errMessages = append(errMessages, "- Receipt Number is required.")
	}

	if len(errMessages) > 0 {
		combinedMessage := strings.Join(errMessages, "\n")
		return receipt, errors.New("validation error:\n" + combinedMessage)
	}

	return receipt, nil
}
