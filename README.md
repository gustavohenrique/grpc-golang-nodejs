It is a proof of concept based in a fictional e-commerce platform. There are two microservices:

1) Catalog: It is written in Go and exposes a REST API that returns a product list.
2) Discount: It is consumed by Catalog, written in NodeJS and returns the received product with a discount applied according of the customer ID sent by X-USER-ID header.

## Getting Started

```shell
git clone https://github.com/gustavohenrique/grpc-golang-nodejs.git
cd grpc-golang-nodejs
docker-compose up -d
```

The ports used are `11443` (Discount) and `11080` (Catalog). 
You can testing using `curl`:

```shell
curl -H 'X-USER-ID: 2' http://localhost:11080/products | python -m json.tool
```

The response looks like this:

```json
[
    {
        "id": 1,
        "title": "Super NES Classic",
        "description": "The Super NES Classic Edition",
        "price_in_cents": 9149,
        "discount_value": {
            "pct": 0.05,
            "value_in_cents": 8692
        }
    },
    {
        "id": 2,
        "title": "Marvel's Spider-Man",
        "description": "PlayStation 4 game",
        "price_in_cents": 5499,
        "discount_value": {
            "pct": 0.05,
            "value_in_cents": 5224
        }
    },
    {
        "id": 3,
        "title": "Nintendo Switch",
        "description": "Neon Red and Neon Blue Joy-Con",
        "price_in_cents": 299,
        "discount_value": {
            "pct": 0.05,
            "value_in_cents": 284
        }
    }
]
```

## Connecting to database

Sometime theres is necessary to connect to database and manipulates records for manual testing.  You can access `pgweb` GUI by http://localhost:8899

## Testing

First of all, you need to create the database container:

```shell
docker run -d --name mydb -p 5123:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=test postgres:alpine
```

### Catalog (Golang)

```shell
cd catalog
test_database_url="postgres://admin:password@127.0.0.1:5123/test?sslmode=disable" go test -v --cover ./src/app/...
```

### Discount (NodeJS)

```shell
npm i
test_database_url="postgres://admin:password@127.0.0.1:5123/test?sslmode=disable" npm test
```

