# Receipt Manager

## Description

Receipt Manager is an Android application written in Go that helps users organize and manage their receipts in one place. The app is designed to make it easy to store, search, and categorize receipts, reducing the need to keep paper copies.

## Goal

The goal of this project is to build a simple, fast, and reliable receipt management application that demonstrates modern Go development practices while providing a practical solution for tracking purchases and expenses.

## Features (Version 1)

- Add new receipts
- View all saved receipts
- Edit receipt information
- Delete receipts
- Search receipts by merchant or description
- Organize receipts by category
- Store purchase date and total amount
- Save receipt images
- Local data storage
- Clean and user-friendly Android interface

## Technologies

- **Language:** Go
- **Platform:** Android
- **Database:** SQLite
- **Version Control:** Git
- **Repository Hosting:** GitHub
- **Build Tool:** Go Modules

## Project Structure

```text
receipt-manager/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── database/
│   ├── models/
│   ├── repository/
│   ├── services/
│   └── handlers/
├── assets/
├── docs/
├── go.mod
├── go.sum
└── README.md
```

## Roadmap

### Version 1.0

- [x] Initialize Go project
- [ ] Create Android application
- [ ] Design user interface
- [ ] Implement local SQLite database
- [ ] Add receipt CRUD operations
- [ ] Add receipt search
- [ ] Add receipt categories
- [ ] Store receipt images
- [ ] Testing and bug fixes

### Future Versions

- Cloud backup and synchronization
- User authentication
- OCR for automatic receipt scanning
- Expense reports and analytics
- Export receipts to PDF or CSV
- Budget tracking
- Multi-device synchronization
- Dark mode
- Barcode and QR code support

## License

This project is licensed under the HYMERS License.