package product

import (
	"errors"
	"log"
	"os"
	"time"
	"context"

	"google.golang.org/grpc"
	
	pb "github.com/gustavohenrique/grpc-golang-nodejs/catalog/ecommerce"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/customer"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
)

type ProductService struct {
	customerRepository interfaces.Repository
	productRepository interfaces.Repository
}

func NewService(repositories map[string]interfaces.Repository) *ProductService {
	return &ProductService{
		customerRepository: repositories["customer"],
		productRepository: repositories["product"],
	}
}

func (this *ProductService) FetchAllProducts() ([]Product, error) {
	p, err := this.productRepository.FetchAll()
	return p.([]Product), err
}

func (this *ProductService) FindCustomerByID(id string) (customer.Customer, error) {
	if id == "" {
		return customer.Customer{}, errors.New("Empty ID.")
	}

	c, err := this.customerRepository.FindByID(id)
	customer := c.(customer.Customer)
	return customer, err
}

func (this *ProductService) ApplyDiscount(customer customer.Customer, products []Product) (interface{}, error) {
	host := os.Getenv("DISCOUNT_SERVICE_HOST")
	if len(host) == 0 {
		host = "localhost:11443"
	}
	conn, err := getDiscountConnection(host)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return products, err
	}
	defer conn.Close()

	c := pb.NewDiscountClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	productsWithDiscountApplied := []*pb.Product{}
	for _, product := range products {
		pbProduct := pb.Product{
			Id: product.ID,
			Title: product.Title,
			Description: product.Description,
			PriceInCents: product.PriceInCents,
		}
		pbCustomer := pb.Customer{
			Id: customer.ID,
			FirstName: customer.FirstName,
			LastName: customer.LastName,
			BirthDate: customer.BirthDate.Format("2006-01-02"),
		}
		r, err := c.ApplyDiscount(ctx, &pb.DiscountRequest{Customer: &pbCustomer, Product: &pbProduct})
		productWithDiscount := r.GetProduct()
		if err == nil {
			productsWithDiscountApplied = append(productsWithDiscountApplied, productWithDiscount)
		} else {
			log.Printf("No discount found to %s.", product.Title)
		}
	}

	if len(productsWithDiscountApplied) > 0 {
		return productsWithDiscountApplied, nil
	}
	return products, nil
}

func getDiscountConnection(host string) (*grpc.ClientConn, error) {
	// wd, _ := os.Getwd()
	// parentDir := filepath.Dir(wd)
	// certFile := filepath.Join(parentDir, "keys", "cert.pem")
	// creds, _ := credentials.NewClientTLSFromFile(certFile, "")
	// return grpc.Dial(host, grpc.WithTransportCredentials(creds))
	return grpc.Dial(host, grpc.WithInsecure())
}
