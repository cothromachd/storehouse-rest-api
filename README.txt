This REST API implements a system of interaction with your storehouse.
It can register a new item, change its state, and delete it.

Little guide of usage:
---------------------------------------------------------
POST:

request:
  URI: "/card/create"
  HTTP Body: *json of new record*
  example:
    {"name": "Coke", "price": 99, "amount":15}
    
response:
  *your json of record*
  example:
    {"id": 1, "name": "Coke", "price":99, "amount":15}
---------------------------------------------------------
GET

request:
  URI: "/card?id=*your integer value*
  example: "card?id=1"
  
response:
  *your json of record*
  example:
    {"id": 1, "name": "Coke", "price": 99, "amount": 15}
---------------------------------------------------------
PUT

request:
  URI:
    "/card/edit?id=*your integer value*"
  example:
    "/card/edit?id=1"
  HTTP Body: *json of new record*
  example:
    {"name": "Coke", "price": 79, "amount": 10}
    
response:
  *your json of edited record*
  example:
    {"name": "Coke", "price": 79, "amount": 10}
---------------------------------------------------------
DELETE

request:
  URI
    "/card/delete?id=*your integer value*"
  example:
    "/card/delete?id=1
    
response:
"record succesfully deleted" if delete is success, error if delete is failed.

Wish you enjoy using it!
