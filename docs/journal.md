# 📓 Development Journal: Building with Fyne

A daily log of concepts learned, challenges tackled, and milestones reached while building desktop applications in Go using the Fyne toolkit.

---

## 🗓️ Daily Entries

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