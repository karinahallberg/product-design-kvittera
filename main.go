package main

import (
	"fmt"
)

// Customer structure stores the name, email, and password of a customer
type Customer struct {
	Name     string
	Email    string
	Password string
}

// Item structure stores the product number, item image, price, returnability, likes, rating, and comments for each item in the purchase
type Item struct {
	ProductNumber string
	Image         string
	Price         float64
	Returnable    bool
	Likes         int
	Rating        int
	Comments      []string
}

// Receipt structure stores a customer and a list of items associated with the customer's purchase
type Receipt struct {
	Customer Customer
	Items    []Item
}

var (
	customers []Customer
	receipts  []Receipt
)

// signUp allows a new customer to sign up by storing their information in the customers slice
func signUp(name string, email string, password string) {
	customers = append(customers, Customer{
		Name:     name,
		Email:    email,
		Password: password,
	})
}

// signIn allows a customer to sign in by checking their email and password against the information stored in the customers slice
func signIn(email string, password string) Customer {
	for _, customer := range customers {
		if customer.Email == email && customer.Password == password {
			return customer
		}
	}
	return Customer{}
}

// registerReceipt allows a customer to register a purchase receipt by storing the customer and items in a new Receipt structure and adding it to the receipts slice
func registerReceipt(customer Customer, items []Item) {
	receipts = append(receipts, Receipt{
		Customer: customer,
		Items:    items,
	})
}

// getReceipt retrieves a purchase receipt for a customer by searching the receipts slice for the receipt associated with the customer's email
func getReceipt(customer Customer) Receipt {
	for _, receipt := range receipts {
		if receipt.Customer.Email == customer.Email {
			return receipt
		}
	}
	return Receipt{}
}

// listItems lists all items in a purchase receipt by returning the Items slice stored in the Receipt structure
func listItems(receipt Receipt) []Item {
	return receipt.Items
}

// likeItem updates the like count for an item by incrementing the Likes field in the Item structure
func likeItem(item *Item) {
	item.Likes++
}

// rateItem adds a rating for an item by updating the Rating field in the Item structure
func rateItem(item *Item, rating int) {
	item.Rating = rating
}

// commentItem adds a comment for an item by adding the comment to the Comments slice stored in the Item structure
func commentItem(item *Item, comment string) {
	item.Comments = append(item.Comments, comment)
}

func main() {
	/* Example usage of the functions
	signUp("John Doe", "johndoe@email.com", "password123")
	customer := signIn("johndoe@email.com", "password123")*/
	
}