# Receipt Manager Database Design

## Purpose

This document describes the data that Receipt Manager will store. The design focuses on the essential information needed for Version 1 of the application.

---

# Entity: Receipt

A receipt represents a user's purchase record.

## Fields

| Field | Type | Description |
|---|---|---|
| ID | Integer | Unique identifier for each receipt |
| Store Name | String | Name of the store or merchant |
| Date | Date | Date the purchase was made |
| Total Amount | Decimal | Total cost of the receipt |
| Category | String | Receipt category (Food, Shopping, etc.) |
| Notes | String | Additional information about the receipt |
| Created At | DateTime | Date and time the receipt was saved |

---

# Entity: Receipt Item

A receipt item represents individual products or services listed on a receipt.

## Fields

| Field | Type | Description |
|---|---|---|
| ID | Integer | Unique identifier for each item |
| Receipt ID | Integer | Links the item to a receipt |
| Item Name | String | Name of the purchased item |
| Quantity | Integer | Number of items purchased |
| Price | Decimal | Price of the item |

---

# Entity Relationships

```text
Receipt
   |
   | 1-to-many
   |
   +---- Receipt Item
```

A single receipt can contain multiple receipt items.

Example:

```text
Receipt
----------------
ID: 1
Store Name: SuperMart
Date: 17/07/2026
Total Amount: $50.00

Receipt Items
----------------
ID: 1
Receipt ID: 1
Item Name: Milk
Quantity: 2
Price: $5.00

ID: 2
Receipt ID: 1
Item Name: Bread
Quantity: 1
Price: $3.00
```

---

# Version 1 Storage Requirements

The application must be able to:

- Save receipt information
- Retrieve saved receipts
- Update receipt details
- Delete receipts
- Search receipts
- Display receipt details

---

# Future Database Improvements

The following data may be added in future versions:

- User accounts
- Cloud synchronization data
- Receipt image storage
- OCR extracted text
- Expense analytics
- Payment method information