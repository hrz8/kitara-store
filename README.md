# Rahman Tennis (Try Go 1.15)

## Quick Start

### Requirements

```bash
Go >= 1.15
MySQL
```

### Usage

``` yaml
# Setting up Apps and DB configurations in `app.config.dev.yml` file
APP_PORT: 8080
DB_HOST: somehost
DB_PORT: 3306
DB_USER: someuser
DB_PASSWORD: somepassword
DB_NAME: some_name
```

``` bash
# Create some_name db
msql> CREATE DATABASE some_name;
```

``` bash
# Run apps with migration and default seed data
$ go run main.go migrate

# Run apps only
$ go run main.go
```

### Test Multiple Request at the same time

- NOTE: this test will PASS if you run `go run main.go migrate` and at the first attempt because the assert checking is based on fresh seed data.

``` bash
# test file in the `domains/order/service` path
# change host url and port constant if you modify app.config
$ go test
```

### Available Endpoints

- Apps run in: `HOST:APP_PORT`

- if you want to run the endpoint manually to create an order, here is the available routes and endpoint:


    1. `POST /api/v1/orders`: Create new player who absolutely not ready

        available productIds:

        - `35731de0-e646-4379-a0f1-b69e74742e0a`, qty `17`, price `13500`
        - `6da51a6f-10c0-404c-82f1-2dce60d720a4`, qty `11`, price `6000`
        - `0bde6df2-f505-401f-882c-808855c2871d`, qty `8` price `17000`

        ``` json
        # Body
        {
            "products": [
                {
                    "id": "0bde6df2-f505-401f-882c-808855c2871d",
                    "qty": 3
                },
                {
                    "id": "35731de0-e646-4379-a0f1-b69e74742e0a",
                    "qty": 3
                }
            ]
        }
        ```

## App Info

### Authors

Hirzi Nurfakhrian

### Version

1.0.0