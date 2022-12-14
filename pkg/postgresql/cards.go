package postgresql

import (
	"context"
	"log"
	//"errors"
	"github.com/cothromachd/rest-api/pkg/models"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type CardModel struct {
	DBpool *pgxpool.Pool
}

func (m *CardModel) Insert(card *models.Card) (int, error) {
	stmt := `INSERT INTO cards (name, price, amount)
	VALUES ($1, $2, $3) RETURNING id;`
	resultRow := m.DBpool.QueryRow(context.Background(), stmt, &card.Name, &card.Price, &card.Amount)
	log.Printf("query is done")
	var id int
	err := resultRow.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *CardModel) Get(id int) (*models.Card, error) {

	stmt := `SELECT id, name, price, amount FROM cards WHERE id = $1;`

	sqlRow := m.DBpool.QueryRow(context.Background(), stmt, id)
	
	s := &models.Card{}

	err := sqlRow.Scan(&s.Id, &s.Name, &s.Price, &s.Amount)
	if err != nil {
		return nil, err
	}

	return s, nil

}

func (m *CardModel) GetAll() ([]*models.Card, error) {
	
	stmt := `SELECT * FROM cards;`

	rows, err := m.DBpool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	log.Print("get all query is done")
	defer rows.Close()

	var cards []*models.Card

	for rows.Next() {
		
		s := &models.Card{}

		err = rows.Scan(&s.Id, &s.Name, &s.Price, &s.Amount)
		if err != nil {
			return nil, err
		}

		cards = append(cards, s)

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return cards, nil
}

func (m *CardModel) Update(card *models.Card, id int) (int, error) {
	
	if card.Id != 0 {

		stmt := `UPDATE cards SET id = $1, name = $2, price = $3, amount = $4 WHERE id = $5 RETURNING id`

		row := m.DBpool.QueryRow(context.Background(), stmt, &card.Id, &card.Name, &card.Price, &card.Amount, id)
	
		log.Print("update is done")

		err := row.Scan(&id)
		if err != nil {
			return -1, err
		}
		return id, nil

	} else {

		stmt := `UPDATE cards SET name = $1, price = $2, amount = $3 WHERE id = $4 RETURNING id`

		row := m.DBpool.QueryRow(context.Background(), stmt, &card.Name, &card.Price, &card.Amount, id)

		log.Print("update is done")

		err := row.Scan(&id)
		if err != nil {
			return -1, err
		}

		return id, nil

	}

}

func (m *CardModel) Delete(id int) (error) {

	stmt := `DELETE FROM cards WHERE id=$1;`
	
	result, err := m.DBpool.Exec(context.Background(), stmt, id)
	if err != nil {
		return err
	}

	log.Printf("delete query is %v", result.Delete())
	return nil

}