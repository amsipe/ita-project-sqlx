package main
import (
	"log"
	_ "github.com/go-sql-driver/mysql"

)

func fetchAllProducts() (products []Product, err error){
	err = dbx.Select(&products,"Select * FROM products")
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func fetchAllOrders() (orders []Order, err error){
	err = dbx.Select(&orders, "SELECT * FROM Orders")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return
}

func fetchOneByOrderID(id string) (details []OrderDetail, err error){
	err = dbx.Select(&details,"SELECT OD.*, P.Name FROM itafinalproject.orders_details OD left join itafinalproject.products P using (ProductID) WHERE OrderID = ?",id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func deleteOneByOrderID(id string) (rowCnt int64, err error) {
	stmt, err := dbx.Prepare("DELETE FROM `orders` WHERE `orders`.`OrderID` = ?;")
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

func addOneProduct(p *Product) (err error) {
	_, err = dbx.NamedExec("INSERT INTO `products` (`Name`, `ImgURL`, `CategoryID`, `Price`, `Description`) VALUES (:Name,:ImgURL,:CategoryID,:Price,:Description)", &p)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// func addOneOrder(od *OrderDetail) err error {
// 	tx, err := sqlx.Begin()
// 	res, err = tx.NamedExec("INSERT INTO `orders` (`OrderTotal`, `CustomerID`) VALUES (:OrderTotal, :CustomerID);", &o)
// 	err := tx.NamedExec()
// }