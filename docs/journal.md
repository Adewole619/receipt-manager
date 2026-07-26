# 📓 Development Journal: Building with Fyne

A daily log of concepts learned, challenges tackled, and milestones reached while building desktop applications in Go using the Fyne toolkit.

---

## 🗓️ Daily Entries

# Day 7

## Date
26 July 2026

## Objective

Understand why databases are needed, learn the role of SQLite in my receipt application, and begin integrating SQLite into my Go project.

---

## What I Learned

### Why I Need a Database

Before today, my application stored receipts in a slice.

Although this allowed me to save multiple receipts while the application was running, I discovered that all the data was lost whenever the application closed because slices are stored in memory (RAM).

A database solves this problem by storing data permanently on disk.

---

### What is SQLite?

SQLite is a lightweight relational database management system that stores an entire database inside a single file.

Unlike databases such as MySQL or PostgreSQL, SQLite does not require a separate database server to run.

This makes SQLite an excellent choice for my receipt manager because:

- It is lightweight.
- It is easy to distribute with my application.
- It stores all data in a single file.
- It works well for desktop and Android applications.

---

### Understanding "Serverless"

I learned that "serverless" does not mean SQLite cannot be used with the internet.

Instead, it means that SQLite does not require a separate database server process.

My Go application communicates directly with the SQLite database file.

```
Go Application

↓

SQLite Driver

↓

receipt.db
```

---

### Database vs Memory

Memory (Slices)

- Fast
- Temporary
- Data disappears when the application closes

SQLite Database

- Permanent
- Stored on disk
- Data remains after restarting the application

---

### Project Architecture

I learned that SQLite should not communicate directly with the user interface.

Instead, every layer has its own responsibility.

```
User

↓

Fyne UI

↓

Validation

↓

Storage

↓

SQLite Database
```

Responsibilities:

- UI → Interacts with the user.
- Validation → Checks that receipt data is valid.
- Storage → Coordinates saving and retrieving receipts.
- SQLite → Permanently stores data and enforces database rules.

---

### Why the Storage Package Exists

I now understand why I created the storage package before introducing SQLite.

Previously:

```
Storage

↓

Slice
```

Later:

```
Storage

↓

SQLite
```

The user interface will not need to change because it only communicates with the storage package.

---

### Database Constraints

I learned that the database should enforce important rules.

For my project, the receipt number should be unique.

If someone attempts to save another receipt with the same receipt number:

- SQLite should reject the insert.
- The application should receive the error.
- The UI should decide how to inform the user.

This keeps responsibilities separated.

---

### Database Connection

I created a new package:

```
internal/
    database/
```

Inside it I began building the database connection.

I learned about:

- `database/sql`
- SQLite drivers
- Blank imports (`_`)
- `sql.Open()`
- Shared database connections

---

## What I Built

- Created the `database` package.
- Installed the SQLite driver.
- Created the initial database connection.
- Learned how Go communicates with SQLite.
- Created (or prepared to create) the `receipt.db` database file.

---

## Challenges

- Understanding why I needed another package (`database`) when I already had a `storage` package.
- Understanding what "serverless" actually means.
- Understanding how SQLite communicates with Go.

---

## Solutions

- Learned that the database package manages the database connection.
- Learned that the storage package manages receipt operations.
- Understood that SQLite stores data permanently while slices only exist during program execution.

---

## Key Takeaways

- Databases provide persistent storage.
- SQLite stores data in one file.
- SQLite does not require a separate database server.
- The UI should never communicate directly with the database.
- The storage package should act as the bridge between the application and SQLite.
- The database should enforce important rules such as unique receipt numbers.

---

## Design Decisions

Today I decided that:

- Receipt numbers must be unique.
- SQLite should enforce uniqueness.
- The UI should handle duplicate receipt errors.
- The storage package should remain responsible for communicating with the database.
- The database package should only manage database connections.

---

## Questions I Asked Today

- Why do I need SQLite?
- What does serverless mean?
- Why shouldn't the UI communicate directly with SQLite?
- Which layer should check for duplicate receipt numbers?
- Why do we need a database package?

---

## Reflection

Today helped me understand that a database is much more than a place to store information. It is responsible for protecting data, enforcing rules, and ensuring information remains available even after the application closes. I also learned the importance of separating responsibilities between the UI, storage layer, and database layer. This architecture will make my application easier to maintain and extend in the future.

---

## Architecture Progress

Current architecture:

```
User

↓

Fyne UI

↓

Validation

↓

Storage

↓

SQLite Database
```

---

## Tomorrow's Goal (Day 8)

- Create the `receipts` table.
- Learn SQL `CREATE TABLE`.
- Save the first receipt into SQLite.
- Replace in-memory storage with database storage.
- Retrieve saved receipts from the database.

# Day 6

## Date
22 July 2026

## Objective

Learn how to store multiple receipts in memory using slices, organize storage logic into its own package, and understand how data flows through the application.

---

## What I Learned

### Slices

- A slice is a dynamic collection that can store multiple values of the same type.
- Instead of storing only one `Receipt`, I can store many `Receipt` objects inside a slice.
- Slices grow dynamically using the `append()` function.

### In-Memory Storage

- Data stored in a slice exists only while the application is running.
- When the application closes, all receipts stored in memory are lost.
- This is called **volatile storage**.
- A database will later provide **persistent storage**, allowing receipts to remain after the application is closed.

### Storage Package

I created a new package:

```
internal/
    storage/
```

This package is responsible for managing receipts stored in memory.

Current responsibilities include:

- Holding the receipt slice.
- Adding new receipts.
- Printing stored receipts.

### Data Flow

I now understand the complete flow when the Save button is clicked:

```
User Input

↓

Create Receipt

↓

Validate Receipt

↓

If validation succeeds

↓

Store Receipt

↓

Print Stored Receipts

↓

Show Success Dialog
```

This helped me understand how information moves through different parts of the application.

---

## What I Built

- Created the `storage` package.
- Added a slice to store multiple receipts.
- Created an `AddReceipt()` function.
- Created a `PrintReceipts()` function.
- Connected the Save button to the storage package.
- Successfully stored multiple receipts during program execution.
- Displayed all stored receipts in the terminal.

---

## Challenges

- Understanding where the receipts slice should be declared.
- Deciding which package should own the slice.
- Understanding why data disappears after the application closes.
- Thinking about who should be responsible for preventing duplicate receipts.

---

## Solutions

- Learned that the slice must exist outside the Save button callback so it survives multiple button clicks.
- Moved storage responsibilities into a dedicated package.
- Learned that memory is temporary and a database will later solve this problem.
- Decided that duplicate checking currently belongs in the storage layer, while the database will eventually enforce uniqueness.

---

## Key Takeaways

- A slice stores multiple values of the same type.
- `append()` adds new elements to a slice.
- Packages should own their responsibilities.
- The UI should not know how receipts are stored.
- Storage should manage receipt collections.
- Databases provide permanent storage, while slices provide temporary storage.

---

## Design Decisions

Today I decided that:

- The storage package should own receipt storage.
- The Save button should only ask the storage package to save a receipt.
- Storage should expose functions like:
  - `AddReceipt()`
  - `PrintReceipts()`
- The UI should not directly manage the receipt slice.

---

## Questions I Asked Today

- Where should the receipts slice be declared?
- Should the receipts slice be global?
- Why should the storage package own the slice?
- Should duplicate receipts be allowed?
- Which layer should be responsible for checking duplicates?

---

## Reflection

Today I learned that storing data is more than simply appending values to a slice. I learned that deciding where data lives and which package owns it is an important software design decision. I also realized that keeping responsibilities separated makes the application easier to maintain and prepares it for future improvements like SQLite.

---

## Architecture Progress

Current project structure:

```
receipt-manager/
│
├── cmd/
│
├── internal/
│   ├── models/
│   ├── validation/
│   └── storage/
│
├── docs/
│
├── README.md
│
└── main.go
```

Current application flow:

```
User

↓

Fyne UI

↓

Receipt Model

↓

Validation

↓

Storage (Slice)

↓

Terminal Output
```

---

## Tomorrow's Goal (Day 7)

- Learn about persistent storage.
- Understand why databases are needed.
- Introduce SQLite.
- Save receipts permanently instead of only in memory.
- Learn how the storage layer can switch from using slices to using a database without changing the UI.

# Day 5

## Date
21 July 2026

## Objective

Learn how to use Go structs to model real-world data, create a Receipt model, and understand why structs make applications easier to maintain as they grow.

---

## What I Learned

### Structs

- A struct is a custom data type that groups related information together.
- Structs allow me to represent real-world objects in code.
- Instead of passing many individual variables, I can pass one struct that contains all the related data.

### Receipt Model

- A Receipt is a real-world object.
- Every receipt has related information such as:
  - Store Name
  - Receipt Number
  - Date
  - Category
  - Total Amount
  - Notes
- These fields belong together, making a struct the best way to represent a receipt.

### Models

- A model is a blueprint for creating objects.
- Every receipt in my application will follow the same structure.
- Different receipts contain different values but use the same Receipt model.

### Project Organization

I created (or planned to create) a new package:

```
internal/
    models/
        receipt.go
```

This package is responsible for defining application data structures.

### Separation of Concerns

I learned that every package should have one responsibility.

- `ui` → User interface
- `validation` → Validate receipt data
- `models` → Define data structures
- `database` → Save and retrieve data (later)

Keeping these responsibilities separate makes the project easier to understand and maintain.

---

## What I Built

- Planned the Receipt model.
- Created the `models` package.
- Created (or prepared to create) the `Receipt` struct.
- Planned to update the validation function so it accepts a Receipt instead of multiple strings.
- Continued improving the project structure.

---

## Challenges

- Understanding when to use a struct instead of multiple variables.
- Deciding what fields belong inside the Receipt model.
- Thinking about how introducing a Receipt struct affects the rest of the application.

---

## Solutions

- Learned that a struct groups related data together.
- Realized that using a Receipt struct keeps function signatures simple.
- Understood that adding new fields later becomes much easier because the Receipt struct can grow without changing every function call.

---

## Key Takeaways

- A struct represents one complete object.
- Related information should be grouped together.
- Models describe the shape of application data.
- Structs improve readability and maintainability.
- Large applications rely heavily on models to organize data.

---

## Design Decisions

Today I decided that:

- A Receipt should be represented by a struct.
- Validation should eventually accept a Receipt instead of individual strings.
- Future features like saving to a database, exporting receipts, and searching receipts should all work with the Receipt model.

---

## Questions I Asked Today

- Should I continue passing individual strings?
- Would a Receipt struct be a better design?
- Why do professional Go applications use models?

---

## Reflection

Today I realized that software design is about organizing information as much as it is about writing code. Instead of thinking about separate variables, I started thinking about a receipt as a single object with related pieces of information. I also understood that using a Receipt struct now will make it much easier to add new features later without rewriting many function calls.

---

## Tomorrow's Goal (Day 6)

- Finish integrating the Receipt model into the application.
- Refactor the validation function to use the Receipt struct.
- Prepare the project for storing receipts.
- Continue improving the application's architecture.

# Day 4

## Date
20 July 2026

## Objective

Learn how to validate user input, organize code into packages, and write cleaner, reusable Go code following good software design principles.

---

## What I Learned

### Form Validation

- Validation is the process of checking user input before using or saving it.
- Validation prevents invalid or incomplete data from entering the application.
- Validation should happen when the user clicks the **Save** button, not when the Entry widget is created.

### Validation Function

- Validation logic should be moved into its own function instead of writing everything inside the button callback.
- A validation function should receive only the data it needs, not UI widgets.
- Passing strings into the validation function makes it reusable for desktop, web, CLI, or mobile applications.

### Error Handling

- Go commonly uses the `error` type to indicate success or failure.
- Instead of returning `(bool, error)`, it is more idiomatic to return only `error`.
- If the returned error is `nil`, validation succeeded.
- If the error is not `nil`, validation failed.

### Multiple Validation Errors

- Instead of stopping at the first error, validation can collect multiple errors.
- Using a slice of strings and `strings.Join()` makes it easy to return one combined error message.

### Separation of Concerns

I learned that every package should have a single responsibility..

Examples:

- `ui` → Build windows and interact with the user.
- `validation` → Validate receipt data.
- `database` → Save and retrieve receipts.

The validation package should not know anything about Fyne dialogs.

### Project Organization

I learned why large Go projects are divided into packages instead of placing everything inside `main.go`.

As projects grow, separating responsibilities makes the code easier to maintain and extend.

---

## What I Built

- Created a new package:

```
internal/
    validation/
        receipt.go
```

- Created a `ValidateReceipt()` function.
- Passed the Store Name and Receipt Number as strings.
- Checked for empty input.
- Collected multiple validation errors.
- Combined validation errors into one error message using `strings.Join()`.
- Returned the combined error.

---

## Challenges

- Deciding what parameters the validation function should receive.
- Understanding whether validation should return a boolean, an error, or both.
- Understanding where `strings.TrimSpace()` should be used.
- Learning how to separate UI code from validation logic.

---

## Solutions

- Decided to pass strings instead of Entry widgets.
- Learned that Go usually returns only an `error` for validation functions.
- Learned that validation functions should protect themselves by trimming input before checking it.
- Learned that dialogs belong in the UI layer, not inside the validation package.

---

## Key Takeaways

- Functions should have one responsibility.
- Validation should be independent of the user interface.
- Packages make applications easier to organize.
- Returning an `error` is the Go convention for reporting validation failures.
- Collecting multiple validation errors provides a better user experience than stopping at the first error.

---

## Design Decisions

Today I made several design decisions before writing code:

- Validation belongs in its own package.
- The validation function should receive strings instead of Fyne widgets.
- Validation should eventually return only an `error`.
- The Save button should only coordinate the process:
  - Read user input.
  - Call validation.
  - Show an error dialog if validation fails.
  - Show a success dialog if validation succeeds.

---

## Questions I Asked Today

- Should validation be placed inside a function?
- Should I pass Entry widgets or strings to the validation function?
- Should validation return a boolean, an error, or both?
- Should validation return the first error or all validation errors?
- Should the validation function eventually receive a Receipt struct instead of multiple parameters?

---

## Reflection

Today I realized that writing software is not just about making the program work. It is also about deciding where each piece of code belongs. I learned that functions and packages should have clear responsibilities so that the application is easier to understand, test, and maintain. I also started thinking more like a software designer by considering how today's code can be reused in future versions of the application.

---

## Tomorrow's Goal (Day 5)

- Refactor the Save button to use the new validation function.
- Move more UI logic into the `ui` package.
- Improve the application's structure.
- Continue building the Receipt Manager with cleaner, more maintainable code.

### Day 3

#### Date
19 July 2026

#### Objective
Learn how to organize widgets using containers and layouts, collect user input using Entry widgets, and build the first functional receipt entry screen.

---

#### What I Learned

##### Fyne Containers
- A container holds one or more widgets.
- A container can also contain another container (nested containers).
- Containers use layouts to arrange widgets.

##### Layouts
- VBox arranges widgets vertically.
- HBox arranges widgets horizontally.
- Different sections of a window can use different layouts.

##### Entry Widget
- An Entry widget allows users to type information.
- The Entry widget stores the current text entered by the user.
- Placeholder text is only a guide and is not actual user input.

##### Buttons and Callbacks
- Each button has its own callback function.
- Clicking a button executes only its callback function.
- The Save button can read the text currently stored in the Entry widgets.

##### Dialogs
- Dialogs display information to the user.
- A dialog must belong to a specific window.
- The window (`w`) tells Fyne where to display the dialog.

---

#### What I Built

- Created a receipt entry form.
- Added:
  - Application title
  - Store Name label
  - Store Name Entry
  - Receipt Number label
  - Receipt Number Entry
  - Save button
  - Cancel button
- Arranged the interface using:
  - VBox
  - HBox
- Created a dialog that displays the Store Name and Receipt Number entered by the user.

---

#### Challenges

- Understanding the difference between a Container and a Layout.
- Understanding why nested containers are necessary.
- Understanding where the Entry widget stores the user's input.
- Understanding why dialogs require the application window.

---

#### Solutions

- Learned that a container is responsible for holding widgets.
- Learned that layouts determine how widgets are arranged.
- Learned that an Entry widget stores its current text internally and updates it whenever the user types.
- Learned that dialogs are attached to a specific window.

---

#### Key Takeaways

- Containers can be nested.
- Different layouts can be combined to create complex interfaces.
- Entry widgets maintain their own internal state.
- Callback functions execute only when an event occurs.
- Building a GUI is about combining small components together.

---

#### Questions I Asked Today

- Why does `dialog.ShowInformation()` require the window?
- Where is the text entered by the user stored?
- Why do we use an `HBox` inside a `VBox` instead of putting everything inside one `VBox`?
- Why does an Entry widget become an empty string (`""`) instead of `nil` when its contents are deleted?

---

#### What I Found Interesting

I discovered that GUI programming is event-driven. The application spends most of its time waiting for the user to perform an action, such as clicking a button or typing into an Entry widget. I also learned that widgets have state, and that state changes as the user interacts with them.

---

#### Tomorrow's Goal (Day 4)

- Learn form validation.
- Prevent users from saving empty receipts.
- Organize the project into separate packages.
- Continue improving the Receipt Manager application.

### Day 2

#### Date
18 July 2026

#### Objective
Learn the basics of Fyne.

#### What I Learned
- What Fyne is.
- How to create an application.
- How to create a window.
- What widgets are.
- What callback functions are.
- Why dialogs require a window.

#### Challenges
- Understanding callback functions.
- Understanding why `dialog.ShowInformation()` needs `w`.

#### Solutions
- Learned that the callback only runs when the button is clicked.
- Learned that dialogs belong to a specific window.

#### Tomorrow
- Learn layouts.
- Learn containers.
- Learn Entry widgets.

---

### Day 1
* **What I Learned:** Set up the Go development environment and installed the Fyne toolkit dependencies.
* **Key Insight:** Explored the basic structure of a Fyne app (`app.New()` and `window.ShowAndRun()`).
* **Milestone:** Initialized the project repository and rendered a blank window.

---

## 🚀 Core Concepts Mastered

* **Fyne Event Loop:** Understanding how the application waits for user input without locking up the UI thread.
* **Window Context:** Managing application windows and modals within the lifecycle of the `fyne.App`.