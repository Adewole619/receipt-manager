# Receipt Manager Wireframes

## Purpose

These wireframes define the basic layout and functionality of the Receipt Manager Android application. The focus is on user flow and screen structure, not colors, fonts, or visual design.

---

# Home Screen

## Purpose

The home screen is the main entry point where users can access receipt management features.

```text
+--------------------------------+
|        Receipt Manager         |
+--------------------------------+
|                                |
|  Total Receipts: 25            |
|  Total Amount: $450.00         |
|                                |
|                                |
|        [+ Add Receipt]         |
|                                |
|        View Receipts           |
|                                |
+--------------------------------+
```

### Functions

- Display receipt summary
- Navigate to add receipt screen
- Navigate to receipt list

---

# Add Receipt Screen

## Purpose

Allows users to create and save a new receipt.

```text
+--------------------------------+
|        Add Receipt             |
+--------------------------------+
| Merchant Name                  |
| [____________________]         |
|                                |
| Date                           |
| [____________________]         |
|                                |
| Amount                         |
| [____________________]         |
|                                |
| Category                       |
| [____________________]         |
|                                |
| Notes                          |
| [____________________]         |
|                                |
|       [Save Receipt]           |
+--------------------------------+
```

### Functions

- Enter merchant name
- Enter receipt date
- Enter receipt amount
- Select category
- Add notes
- Save receipt

---

# Receipt List Screen

## Purpose

Displays all saved receipts.

```text
+--------------------------------+
|        Receipts                |
+--------------------------------+
| Search: [_____________]        |
|                                |
|--------------------------------|
| Store A             $25.00     |
| 17/07/2026                     |
|                                |
|--------------------------------|
| Store B             $50.00     |
| 16/07/2026                     |
|                                |
|--------------------------------|
| Store C             $15.00     |
| 15/07/2026                     |
|                                |
+--------------------------------+
```

### Functions

- View all receipts
- Search receipts
- Select a receipt to view details

---

# Receipt Details Screen

## Purpose

Displays complete information about a selected receipt.

```text
+--------------------------------+
|      Receipt Details           |
+--------------------------------+
| Merchant: Store A              |
| Date: 17/07/2026               |
| Amount: $25.00                 |
| Category: Food                 |
| Notes: Lunch purchase          |
|                                |
|                                |
| [Edit]          [Delete]       |
|                                |
+--------------------------------+
```

### Functions

- View receipt information
- Edit receipt
- Delete receipt

---

# User Flow

```text
Home
 |
 +--> Add Receipt
 |       |
 |       +--> Save Receipt
 |
 +--> Receipt List
         |
         +--> Receipt Details
                 |
                 +--> Edit Receipt
                 |
                 +--> Delete Receipt
```