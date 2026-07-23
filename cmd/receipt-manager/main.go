package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"receipt-manager/internal/models"
	"receipt-manager/internal/storage"
	"receipt-manager/internal/validation"
)

func main() {
	// Welcome message
	fmt.Println("Hello, Welcome to the Receipt Manager Application!")
	fmt.Println("This application helps you manage your receipts efficiently.")
	fmt.Println("You can add, view, and delete receipts as needed.")
	fmt.Println("Let's get started!")

	fmt.Println()

	fmt.Println("Testing app.NewApp()")

	myApp := app.NewWithID("receipt-manager")

	myWindow := myApp.NewWindow("Receipt Manager")

	myWindow.Resize(fyne.NewSize(800, 600))

	title1 := widget.NewLabel("Receipt Manager v1.0")
	title2 := widget.NewLabel("Welcome to my first Go desktop Application!")

	welcome := widget.NewLabel("This application helps you manage your receipts efficiently.")

	storeEntry := widget.NewEntry()

	storeEntry.SetText("Shoprite")

	receiptEntry := widget.NewEntry()

	receiptEntry.SetText("RCP-000001")

	storeLabel := widget.NewLabel("Store Name")
	receiptLabel := widget.NewLabel("Receipt Number")

	saveBtn := widget.NewButton("Save", func() {

		receipt := models.Receipt{
			StoreName:     storeEntry.Text,
			ReceiptNumber: receiptEntry.Text,
		}
		receipt , err := validation.ValidateReceipt(receipt)

		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		storage.AddReceipt(receipt)
		storage.PrintReceipts(storage.Receipts)

		fmt.Printf("Saving receipt: Store Name: %s, Receipt Number: %s\n", receipt.StoreName, receipt.ReceiptNumber)
		dialog.ShowInformation("Receipt Saved", fmt.Sprintf("Receipt saved:\nStore Name: %s\nReceipt Number: %s", receipt.StoreName, receipt.ReceiptNumber), myWindow)
	})

	cancelBtn := widget.NewButton("Cancel", func() {
		myWindow.Close()
	})

	buttonBox := container.NewHBox(saveBtn, cancelBtn)

	addButton := widget.NewButton("+ Add Receipt", func() {
		fmt.Println("Add Receipt button clicked!")

		dialog.ShowInformation("Add Receipt Message", "Add Receipt window is under construction.", myWindow)
	})

	viewButton := widget.NewButton("View Receipts", func() {
		fmt.Println("View Receipts button clicked!")

		dialog.ShowInformation("View Receipts Message", "View Receipts window is under construction.", myWindow)
	})

	settingsButton := widget.NewButton("Settings", func() {
		fmt.Println("Settings button clicked!")

		dialog.ShowInformation("Settings Message", "Settings window is under construction.", myWindow)
	})

	content := container.NewVBox(
		title1,
		title2,
		welcome,
		storeLabel,
		storeEntry,
		receiptLabel,
		receiptEntry,
		buttonBox,
		addButton,
		viewButton,
		settingsButton,
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()

}
