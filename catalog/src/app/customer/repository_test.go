package customer_test

import (
	"os"
	"testing"
	"log"

	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/database"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/customer"
)

var (
    db      database.Database
    repository interfaces.Repository
)

func setUpRepository() {
	databaseURL := os.Getenv("test_database_url")
    db := database.NewDB(databaseURL)
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec("DROP TABLE IF EXISTS customers;")
	conn.Exec("CREATE TABLE customers (id SERIAL PRIMARY KEY, first_name VARCHAR, last_name VARCHAR, birth_date DATE);")
    tx, _ := conn.Begin()
    stmt, _ := tx.Prepare("INSERT INTO customers (id, first_name, last_name, birth_date) VALUES ($1, $2, $3, $4)")
    stmt.Exec(1, "Myrcella", "Baratheon", "1999-10-01")
    stmt.Exec(2, "Daenerys", "Targaryen", "1986-10-23")
    stmt.Exec(3, "Margaery", "Tyrell", "1982-02-11")
    stmt.Exec(4, "Ellaria", "Sand", "1973-05-14")
    tx.Commit()

    repository = customer.NewRepository(db)
}

func TestFindCustomerById(t *testing.T) {
    setUpRepository()
	c, _ := repository.FindByID("3")
	customer := c.(customer.Customer)
    
    if customer.ID != 3 {
        t.Errorf("Got %d but I want %d", customer.ID, 3)
    }
}

func TestNotFoundCustomerById(t *testing.T) {
    setUpRepository()
	_, err := repository.FindByID("0")
    
    if err == nil {
        t.Errorf("No customer should be found. %s", err)
    }
}