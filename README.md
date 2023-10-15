# order-pack

![Status](https://github.com/murat/order-pack/actions/workflows/test.yml/badge.svg)

[![Status](https://cloud.drone.io/api/badges/murat/order-pack/status.svg)](https://cloud.drone.io/murat/order-pack)

## Description

**order-pack** ...

- [ ] write instructions

## Installation & Usage

order-pack developed by using only standard library.

- [ ] write instructions

### Locally

```shell
git clone git@github.com:murat/order-pack.git && cd order-pack

go build -o ./bin/order-pack ./cmd/main.go # or make build

go run cmd/main.go

# or
./bin/order-pack
```

### Docker

```shell
git clone git@github.com:murat/order-pack.git && cd order-pack

docker build . -t order-pack
docker run -p 8080:8080 order-pack
```

order-pack is going to start a http server that listens **8080** port.

### Usage

```shell
# Create product packs
http --form POST :8080/products name=100 size=100 && \
http --form POST :8080/products name=250 size=250 && \
http --form POST :8080/products name=500 size=500 && \
http --form POST :8080/products name=1000 size=1000 && \
http --form POST :8080/products name=2000 size=2000 && \
http --form POST :8080/products name=5000 size=5000

# Create order
# Will not save order to db
echo -n '{"item_count":501}' | http POST :8080/orders
## Resp:
{
    "data": {
        "CreatedAt": "0001-01-01T00:00:00Z",
        "ID": 0,
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "item_count": 501,
        "items": [
            {
                "count": 1,
                "product": {
                    "CreatedAt": "2023-10-15T15:32:41.57313+03:00",
                    "ID": 3,
                    "UpdatedAt": "2023-10-15T15:32:41.57313+03:00",
                    "name": "500",
                    "package_size": 500
                }
            },
            {
                "count": 1,
                "product": {
                    "CreatedAt": "2023-10-15T15:32:40.837342+03:00",
                    "ID": 1,
                    "UpdatedAt": "2023-10-15T15:32:40.837342+03:00",
                    "name": "100",
                    "package_size": 100
                }
            }
        ]
    }
}

```

## Specs

- [ ] write specs

## Contribute

All PR’s and issues are welcome!