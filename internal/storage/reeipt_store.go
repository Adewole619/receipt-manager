package storage

import (
	"fmt"
	"receipt-manager/internal/models"
)

var Receipts []models.Receipt

func PrintReceipts(rcp []models.Receipt) {

	fmt.Println("===== RECEIPTS =====")
	fmt.Println()
	
	if len(rcp) == 0 {
		fmt.Println("No receipts found")
		return
	}

	for i, rec := range rcp {

		fmt.Printf("%d.\n Store: %s\n Receipt: %s\n", i+1, rec.StoreName, rec.ReceiptNumber)
		fmt.Println()
	}
}

func AddReceipt(rcpp models.Receipt) {
	Receipts = append(Receipts, rcpp)
}
