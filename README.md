This REST API implements a system of interaction with your storehouse.
It can register a new item, change its state, and delete it.

Little guide of usage:
=================================================================================================
|| POST                                      || GET                                            ||
||                                           ||                                                ||
|| request:                                  || request:                                       ||
||    URI: "/card/create"                    ||    URI:     "/card?id=*your integer value*"    ||
||    HTTP Body: *json of new record*        ||    example: "/card?id=1"                       ||
||    example:                               || response:                                      ||
||        {"name": "Coke", "price": 99,      ||     *your json of record*                      ||
||         "amount": 15}                     ||     example: {"id": 1, "name": "Coke",         ||
|| response:                                 ||               "price": 99, "amount": 15}       ||
||    *your json of record*                  ||                                                ||
||    example:                               ||                                                ||
||        {"id": 1, "name": "Coke",          ||                                                ||
||         "price": 99, "amount": 15}        ||                                                ||
=================================================================================================
|| PUT                                       ||
||                                           ||           
|| request:                                  ||
||    URI:                                   ||
||      "/card/edit?id=*your integer value*" ||
||    example:                               ||
||      "/card/edit?id=1"                    ||
||    HTTP Body *json of new record*         ||            Wish you enjoy using it! ;)
||    example:                               ||
||          {"name": "Coke",                 ||
||           "price": 79, "amount: 10}       ||
|| response:                                 ||
||    *your json of record*                  ||
||    example:                               ||
||          {"name": "Coke",                 ||
||           "price": 79, "amount: 10}       ||
===============================================