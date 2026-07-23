package validation

import (
	"strings"
	"testing"

	"receipt-manager/internal/models"
)

func TestValidateReceiptReturnsTrimmedValuesOnSuccess(t *testing.T) {
	receipt, err := ValidateReceipt(models.Receipt{
		StoreName:     "  Shoprite  ",
		ReceiptNumber: "  RCP-000001  ",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if receipt.StoreName != "Shoprite" {
		t.Fatalf("expected trimmed store name 'Shoprite', got %q", receipt.StoreName)
	}

	if receipt.ReceiptNumber != "RCP-000001" {
		t.Fatalf("expected trimmed receipt number 'RCP-000001', got %q", receipt.ReceiptNumber)
	}
}

func TestValidateReceiptReturnsCleanValuesWhenValidationFails(t *testing.T) {
	receipt, err := ValidateReceipt(models.Receipt{
		StoreName:     "   ",
		ReceiptNumber: "  RCP-000001  ",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	if receipt.StoreName != "" {
		t.Fatalf("expected empty store name, got %q", receipt.StoreName)
	}

	if receipt.ReceiptNumber != "RCP-000001" {
		t.Fatalf("expected trimmed receipt number 'RCP-000001', got %q", receipt.ReceiptNumber)
	}

	if !strings.Contains(err.Error(), "Store Name is required") {
		t.Fatalf("expected error to mention missing store name, got %v", err)
	}
}
