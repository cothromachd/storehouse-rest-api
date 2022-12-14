package models

type Card struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

/*
CREATE TABLE cards (
id BIGSERIAL NOT NULL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
price INTEGER NOT NULL,
amount INTEGER,
);


INSERT INTO cards (name, price, amount)
VALUES ('Zhumaysinba', 100, 15);

INSERT INTO snippets (title, content, created, expires) (
VALUES ('Не имей сто рублей',
'Не имей сто рублей,\nа имей сто друзей.',
NOW() AT TIME ZONE ('UTC'),
NOW() AT TIME ZONE ('UTC') + INTERVAL '365' DAY
);
*/