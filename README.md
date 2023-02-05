# IKEA Family KVITTERA (Purchases and Returns)

This program is designed to allow IKEA Family customers to sign up and sign in, access their purchase receipts, and manage their items. It includes the following features:

- Sign up and sign in for IKEA Family customers
- Access a registered purchase receipt
- Register a purchase receipt
- List all items in the purchase
- All items have the product number, item image, and price
- Each item that is eligible for return will have a return option
- Each item has a like icon
- Each item can be scored between 1-5
- Each item has a comment section

## Data Structures

The program uses the following data structures to store customer information, purchase receipts, and items:

- `Customer`: This structure stores the name, email, and password of a customer.
- `Receipt`: This structure stores a customer and a list of items associated with the customer's purchase.
- `Item`: This structure stores the product number, item image, price, returnability, likes, rating, and comments for each item in the purchase.

## Functions

The program includes the following functions to manage customer information, purchase receipts, and items:

- `signUp(name string, email string, password string)`: This function allows a new customer to sign up by storing their information in the `customers` slice.
- `signIn(email string, password string) Customer`: This function allows a customer to sign in by checking their email and password against the information stored in the `customers` slice.
- `registerReceipt(customer Customer, items []Item)`: This function allows a customer to register a purchase receipt by storing the customer and items in a new `Receipt` structure and adding it to the `receipts` slice.
- `getReceipt(customer Customer) Receipt`: This function retrieves a purchase receipt for a customer by searching the `receipts` slice for the receipt associated with the customer's email.
- `listItems(receipt Receipt) []Item`: This function lists all items in a purchase receipt by returning the `Items` slice stored in the `Receipt` structure.
- `likeItem(item *Item)`: This function updates the like count for an item by incrementing the `Likes` field in the `Item` structure.
- `rateItem(item *Item, rating int)`: This function adds a rating for an item by updating the `Rating` field in the `Item` structure.
- `commentItem(item *Item, comment string)`: This function adds a comment for an item by adding the comment to the `Comments` slice stored in the `Item` structure.

## Getting Started

To use this program, you will need to have Go installed on your computer. You can then clone or download the program and run it using a Go compiler or an IDE. You can also modify the code to suit your needs and add additional functionality as needed.
