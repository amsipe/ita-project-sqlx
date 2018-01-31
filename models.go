package main


type Product struct {
	ProductID int `db:"ProductID"`
	Name string `db:"Name"`
	ImgUrl string `db:"ImgURL"`
	CategoryID int `db:"CategoryID"`
	Price float32 `db:"Price"`
	Description string `db:"Description"`
}

type Order struct {
	OrderID int `db:"OrderID"`
	OrderDateTime string `db:"OrderDateTime"`
	OrderTotal float32 `db:"OrderTotal"`
	CustomerID int `db:"CustomerID"`
	Completed int `db:"Completed"`
}

type OrderDetail struct {
	OrderDetailID int `db:"OrderDetailID"`
	OrderID int `db:"OrderID"`
	ProductID int `db:"ProductID"`
	Quantity int `db:"Quantity"`
	Price float32 `db:"Price"`
	Name string `db:"Name"`
}