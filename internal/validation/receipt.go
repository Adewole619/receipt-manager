package validation

import (
	"errors"
	"strings"
)

// Validation: Store Name check and 🛠️ Receipt Number check
func ValidateReceipt(storeName, receiptNumber string) (string, string, error) {

	var errMessages []string

	cleanStoreName := strings.TrimSpace(storeName)
	cleanReceiptNumber := strings.TrimSpace(receiptNumber)

	if cleanStoreName == "" {

		errMessages = append(errMessages, "- Store Name is required.")
	}

	if cleanReceiptNumber == "" {

		errMessages = append(errMessages, "- Receipt Number is required.")
	}
	if len(errMessages) > 0 {
		combinedMessage := strings.Join(errMessages, "\n")
		return "", "", errors.New("validation Error:\n" + combinedMessage)
	}
	return cleanStoreName, cleanReceiptNumber, nil
}
