# 📓 Development Journal: Building with Fyne

A daily log of concepts learned, challenges tackled, and milestones reached while building desktop applications in Go using the Fyne toolkit.

---

## 🗓️ Daily Entries

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