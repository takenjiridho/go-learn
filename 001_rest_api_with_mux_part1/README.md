## 001. REST API with MUX Part 1

``` bash
# go to 001_rest_api_with_mux_part1
cd 001_rest_api_with_mux_part1/
# build and run 
go build -o rest_api && ./rest_api
```

## Endpoints

### Get All Cars
``` bash
GET api/cars
```
### Get Single Car
``` bash
GET api/cars/{id}
```

### Delete Book
``` bash
DELETE api/cars/{id}
```

### Create Car
``` bash
POST api/cars

# Request sample
{
    "manufacture": "Daihatsu",
    "product": "Terrios",
    "year": "2015",
    "engine": "15L",
    "model": {
        "code": "1TR",
        "type": "Advanture",
        "price": "230 jt",
        "color": "White"
    }
}
```

### Update Car
``` bash
PUT api/cars/{id}

# Request sample
{
        "id": "887",
        "manufacture": "Daihatsu",
        "product": "Terrios",
        "year": "2014",
        "engine": "15L",
        "model": {
            "code": "1TR",
            "type": "TX",
            "price": "130 jt",
            "color": "Black"
        }
    }

```
