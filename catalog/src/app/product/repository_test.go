package product_test

import (
	"os"
	"testing"
	"log"

	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/database"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/product"
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

	conn.Exec("DROP TABLE IF EXISTS products;")
	conn.Exec(`
	CREATE TABLE products (
		id SERIAL PRIMARY KEY,
		title VARCHAR,
		description VARCHAR,
		price_in_cents INT
	);`)

    tx, _ := conn.Begin()
    stmt, _ := tx.Prepare("INSERT INTO products VALUES ($1, $2, $3, $4)")
    stmt.Exec(1, "Super NES Classic", "The Super NES Classic Edition", 9149)
    stmt.Exec(2, "Marvel’s Spider-Man", "PlayStation 4 game", 5499)
    stmt.Exec(3, "Nintendo Switch", "Neon Red and Neon Blue Joy-Con", 299)
    tx.Commit()

    repository = product.NewRepository(db)
}

func TestFetchAllProducts(t *testing.T) {
    setUpRepository()
	p, _ := repository.FetchAll()
	products := p.([]product.Product)
    
    if len(products) != 3 {
        t.Errorf("Got %d but I want %d", len(products), 3)
    }
}
