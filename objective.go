// Data structures to store customer information, purchase receipts, and items
type Customer struct {
  Name string
  Email string
  Password string
}

type Receipt struct {
  Customer Customer
  Items []Item
}

type Item struct {
  ProductNumber string
  Image string
  Price float64
  IsReturnable bool
  Likes int
  Rating int
  Comments []string
}

// Variables to store customer information and purchase receipts
var customers []Customer
var receipts []Receipt

// Function to sign up a new customer
func signUp(name string, email string, password string) {
  customer := Customer{
    Name: name,
    Email: email,
    Password: password,
  }
  customers = append(customers, customer)
}

// Function to sign in a customer
func signIn(email string, password string) Customer {
  for _, customer := range customers {
    if customer.Email == email && customer.Password == password {
      return customer
    }
  }
  return Customer{}
}

// Function to register a purchase receipt
func registerReceipt(customer Customer, items []Item) {
  receipt := Receipt{
    Customer: customer,
    Items: items,
  }
  receipts = append(receipts, receipt)
}

// Function to retrieve a purchase receipt
func getReceipt(customer Customer) Receipt {
  for _, receipt := range receipts {
    if receipt.Customer.Email == customer.Email {
      return receipt
    }
  }
  return Receipt{}
}

// Function to list all items in a purchase receipt
func listItems(receipt Receipt) []Item {
  return receipt.Items
}

// Function to update the like count for an item
func likeItem(item *Item) {
  item.Likes++
}

// Function to add a rating for an item
func rateItem(item *Item, rating int) {
  item.Rating = rating
}

// Function to add a comment for an item
func commentItem(item *Item, comment string) {
  item.Comments = append(item.Comments, comment)
}
