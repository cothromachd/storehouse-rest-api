package main

import (
	"context"
	"flag"

	//"fmt"
	"log"
	"net/http"

	"github.com/cothromachd/rest-api/pkg/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	cards *postgresql.CardModel
}

func main() {

	// флагирую
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	dsn := flag.String("dsn", "postgres://khalidmagnificent:190204@localhost:5432/card_storage", "PostgreSQL name of data source")
	
	flag.Parse()
	// нафлагировал 






	// создаю пул подключений к бд
	db, err := pgxpool.New(context.Background(), *dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// жестко создал пул подкл к бд


	// внедряю зависимости 
	app := &application{
		cards: &postgresql.CardModel{DBpool: db},
	}


	// маршрутизирую
	mux := http.NewServeMux()


	mux.HandleFunc("/card/create", app.createCard)

	mux.HandleFunc("/card", app.showCard)

	mux.HandleFunc("/card/all", app.showAllCards)

	mux.HandleFunc("/card/edit", app.editCard)

	mux.HandleFunc("/card/delete", app.deleteCard)

	log.Printf("Запуск сервера на %s", *addr)

	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
	// жестко намаршрутизировал + запутил сервак
}