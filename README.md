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

## Project Status

🚧 In Progress

Current Stage:
- Completed project planning
- Set up Go module and Git repository
- Built the initial Fyne application window
- Added the home screen with placeholder buttons

### Features Progress

#### Completed

- Application window
- Home screen
- Welcome message
- Add Receipt button
- View Receipts button
- Settings button
- Placeholder dialogs

#### Planned

- Add Receipt form
- Receipt list
- Receipt details
- Search receipts
- SQLite database
- PDF export
- Reports

## Getting Started

#### Clone the repository

```bash
git clone <repository-url>
```

#### Install dependencies

On Linux, Fyne/GLFW requires native development libraries for X11/Wayland and OpenGL.

Debian/Ubuntu:

```bash
sudo apt update
sudo apt install -y libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxkbcommon-dev libgl1-mesa-dev libwayland-dev libwayland-egl-backend-dev
```

Fedora/RHEL:

```bash
sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libxkbcommon-devel mesa-libGL-devel wayland-devel wayland-protocols-devel
```

Arch Linux:

```bash
sudo pacman -Syu mesa libx11 libxcursor libxrandr libxinerama libxkbcommon wayland
```

Then install Go module dependencies:

```bash
go mod tidy
```

#### Run the application

```bash
go run ./cmd/receipt-manager
```

#### Run with Docker

Build the image:

```bash
docker build -t receipt-manager:local .
```

Run the container on a Linux host with X11 access:

```bash
xhost +local:root

docker run --rm -it \
  -e DISPLAY=$DISPLAY \
  -v /tmp/.X11-unix:/tmp/.X11-unix:rw \
  --network host \
  receipt-manager:local
```

Or with Docker Compose:

```bash
docker compose up --build
```

#### Technologies

- Go
- Fyne
- SQLite (planned)
- Git

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

#### Roadmap Progress

- [x] Project planning
- [x] Project setup
- [x] Initial Fyne window
- [x] Home screen
- [ ] Add Receipt form
- [ ] SQLite database
- [ ] Search functionality
- [ ] Reports
- [ ] PDF export

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

