package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"receipt-manager/internal/models"
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

	// Our first line of code wi'll create a new application using the app.NewApp() function.
	// This function initializes a new Fyne application and returns an instance of the app.
	// App interface.
	myApp := app.NewWithID("receipt-manager")

	// The app variable now holds the instance of our Fyne application.
	// We can use this instance to create windows, set up the user interface,
	//  and manage the application's lifecycle.
	// You can think of the app as the main entry point for your Fyne application,
	//  and it provides various methods and properties to customize and control the
	//  behavior of your application.

	// Next, we create a new window for our application using the app.NewWindow() method.
	// This method takes a string parameter that represents the title of the window.
	// In this case, we set the title to "Receipt Manager".
	// The NewWindow() method returns an instance of the fyne.Window interface,
	//  which represents the main window of our application.

	// You can use any name you like for the window title, but "Receipt Manager" is a
	//  descriptive name that reflects the purpose of our application.
	myWindow := myApp.NewWindow("Receipt Manager")

	//resize the window to a specific size using the Resize() method.
	// The Resize() method takes a fyne.Size parameter that specifies the width
	//  and height of the window.
	// In this case, we set the size to 800 pixels wide and 600 pixels tall.
	// You can adjust these values to fit your application's needs.
	myWindow.Resize(fyne.NewSize(800, 600))

	// Next, we set the content of the window using the SetContent() method.
	// The SetContent() method takes a fyne.CanvasObject parameter that represents
	//  the content to be displayed in the window.
	// In this case, we create a new label using the widget.NewLabel() function,
	//  which displays the text "Welcome to my first Go desktop Application!" in the window.

	//w.SetContent(widget.NewLabel("Welcome to my first Go desktop Application!"))

	// Instead of directly setting the label as the content, we can create a vertical box container
	//  to hold the label and the buttons. This allows us to arrange the elements vertically
	//  and create a more organized layout for our application.
	title1 := widget.NewLabel("Receipt Manager v1.0")
	title2 := widget.NewLabel("Welcome to my first Go desktop Application!")

	welcome := widget.NewLabel("This application helps you manage your receipts efficiently.")

	// Task 1: Create two Entry widgets
	// --- Input Fields ---
	storeEntry := widget.NewEntry()
	// Set the default store name text
	storeEntry.SetText("Shoprite")

	receiptEntry := widget.NewEntry()
	// Set the default receipt number text
	receiptEntry.SetText("RCP-000001")

	// Task 2: Create two Labels
	storeLabel := widget.NewLabel("Store Name")
	receiptLabel := widget.NewLabel("Receipt Number")

	// Task 3: Create two Buttons
	saveBtn := widget.NewButton("Save", func() {
		// Callback for saving data

		// storeName := strings.TrimSpace(storeEntry.Text)
		// receiptNumber := strings.TrimSpace(receiptEntry.Text)
		receipt := models.Receipt{
			StoreName:     storeEntry.Text,
			ReceiptNumber: receiptEntry.Text,
		}
		storeName, receiptNumber, err := validation.ValidateReceipt(receipt)

		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		fmt.Printf("Saving receipt: Store Name: %s, Receipt Number: %s\n", storeName, receiptNumber)
		dialog.ShowInformation("Receipt Saved", fmt.Sprintf("Receipt saved:\nStore Name: %s\nReceipt Number: %s", storeName, receiptNumber), myWindow)
	})

	cancelBtn := widget.NewButton("Cancel", func() {
		myWindow.Close()
	})

	// Task 4: Put the buttons inside an HBox
	buttonBox := container.NewHBox(saveBtn, cancelBtn)

	// Next, we create two buttons using the widget.NewButton() function.
	// The first button is labeled "+ Add Receipt" and has a callback function that
	//  prints "Add Receipt button clicked!" to the console when the button is clicked.
	// The second button is labeled "View Receipts" and has a callback function that
	//  prints "View Receipts button clicked!" to the console when the button is clicked.
	addButton := widget.NewButton("+ Add Receipt", func() {
		fmt.Println("Add Receipt button clicked!")
		// we will add the code to open a new window for adding receipts here in the future.

		// For now, we will display a dialog box to inform the user that
		dialog.ShowInformation("Add Receipt Message", "Add Receipt window is under construction.", myWindow)
	})

	// The second button is labeled "View Receipts" and has a callback function that
	//  prints "View Receipts button clicked!" to the console when the button is clicked.
	viewButton := widget.NewButton("View Receipts", func() {
		fmt.Println("View Receipts button clicked!")
		// we will add the code to open a new window for viewing receipts here in the future.

		// For now, we will display a dialog box to inform the user that
		dialog.ShowInformation("View Receipts Message", "View Receipts window is under construction.", myWindow)
	})

	// The third button is labeled "Settings" and has a callback function that
	//  prints "Settings button clicked!" to the console when the button is clicked.
	settingsButton := widget.NewButton("Settings", func() {
		fmt.Println("Settings button clicked!")

		// we will add the code to open a new window for viewing receipts here in the future.

		// For now, we will display a dialog box to inform the user that
		// the settings window is under construction.
		dialog.ShowInformation("Settings Message", "Settings window is under construction.", myWindow)
	})

	// Next, we create a vertical box container using the container.NewVBox() function.
	// The NewVBox() function takes one or more fyne.CanvasObject parameters that represent
	//  the content to be displayed in the container.
	// In this case, we pass the label and the two buttons as parameters to the NewVBox() function.
	// The vertical box container arranges its child objects vertically, one below the other.
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

	// Finally, we set the content of the window to the
	// vertical box container using the SetContent() method.
	myWindow.SetContent(content)

	// Finally, we call the ShowAndRun() method on the window instance.
	// This method displays the window and starts the application's main event loop.
	// The main event loop listens for user interactions, such as button clicks or
	//  text input, and updates the user interface accordingly.
	// The ShowAndRun() method blocks the execution of the program until the window is closed,
	//  allowing the application to run continuously until the user decides to exit.
	myWindow.ShowAndRun()

}
