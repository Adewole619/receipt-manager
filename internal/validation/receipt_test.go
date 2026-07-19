package validation

import (
	"strings"
	"testing"
)

func TestValidateReceiptReturnsTrimmedValuesOnSuccess(t *testing.T) {
	storeName, receiptNumber, err := ValidateReceipt("  Shoprite  ", "  RCP-000001  ")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if storeName != "Shoprite" {
		t.Fatalf("expected trimmed store name 'Shoprite', got %q", storeName)
	}

	if receiptNumber != "RCP-000001" {
		t.Fatalf("expected trimmed receipt number 'RCP-000001', got %q", receiptNumber)
	}
}

func TestValidateReceiptReturnsCleanValuesWhenValidationFails(t *testing.T) {
	storeName, receiptNumber, err := ValidateReceipt("   ", "  RCP-000001  ")
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	if storeName != "" {
		t.Fatalf("expected empty store name, got %q", storeName)
	}

	if receiptNumber != "RCP-000001" {
		t.Fatalf("expected trimmed receipt number 'RCP-000001', got %q", receiptNumber)
	}

	if !strings.Contains(err.Error(), "Store Name is required") {
		t.Fatalf("expected error to mention missing store name, got %v", err)
	}
}
