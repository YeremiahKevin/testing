# Testing
Change database credentials in function `openDBConnection` to your own database

How To Run
1. `go mod vendor`
2. `go run .`
3. HTTP Server will serve at port 8090
4. Hit using this cURL

``curl --location --request POST 'localhost:8090/insert' \
--header 'Content-Type: application/json' \
--data-raw '{
"request_id": 12334355,
"data": [
{
"id": 12345,
"customer": "John Smith",
"quantity": 1,
"price": 10.00,
"timestamp": "2022-01-01 22:10:44"
}
]
}'``