## Running the Clinic

To run this exmaple, from the root of this project:

```sh
go run ./*.go
```

## URLS samples
curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 50.52, "kind": "custo", "day": 1, "month": 2, "year": 2017, "clinic": "Três Pontas"}' http://localhost:3000/closures
curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 100.52, "kind": "custo", "day": 2, "month": 2, "year": 2017, "clinic": "Três Pontas"}' http://localhost:3000/closures
curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 100.52, "kind": "custo", "day": 3, "month": 2, "year": 2017, "clinic": "Três Pontas"}' http://localhost:3000/closures

curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 50.52, "kind": "custo", "day": 1, "month": 2, "year": 2017, "clinic": "Lapa"}' http://localhost:3000/closures
curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 100.52, "kind": "custo", "day": 2, "month": 2, "year": 2017, "clinic": "Lapa"}' http://localhost:3000/closures
curl -XPOST  -H 'Content-Type: application/json'  -d '{"name": "Salário Luana", "value": 100.52, "kind": "custo", "day": 3, "month": 2, "year": 2017, "clinic": "Lapa"}' http://localhost:3000/closures


curl -XPUT  -d '{"id": "5890c8a7b3abdc64a639e58e", "name": "Salário Thiago", "value2, "year": 2017}' http://localhost:3000/closures/5890c8a7b3abdc64a639e58e
curl -XDELETE http://localhost:3000/closures/5497246c380a967ff1000003

