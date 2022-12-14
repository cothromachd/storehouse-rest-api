package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/cothromachd/rest-api/pkg/models"
)



func (app *application) createCard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Print("unexpected http method: expected POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	reqBody, _ := io.ReadAll(r.Body)

	var card models.Card

	err := json.Unmarshal(reqBody, &card)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	id, err := app.cards.Insert(&card)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r,  fmt.Sprintf("/card?id=%d", id), http.StatusSeeOther)

}

func (app *application) showCard(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	s, err := app.cards.Get(id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sJson, err := json.Marshal(*s)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "%s", sJson)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}


func (app *application) showAllCards(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Print("unexpected http method: expected GET")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data, err := app.cards.GetAll()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	allCardsJson, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	_, err = fmt.Fprintf(w, "%s", allCardsJson)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	
}



func (app *application) editCard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {

		log.Print("unexpected http method: expected PUT")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	id, err:= strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var card models.Card

	json.NewDecoder(r.Body).Decode(&card)

	id, err = app.cards.Update(&card, id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r,  fmt.Sprintf("/card?id=%d", id), http.StatusSeeOther)
}


func (app *application) deleteCard(w http.ResponseWriter, r *http.Request) {

	// получаем id записи, которую нужно удалить
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	// удаляем запись из бд
	err = app.cards.Delete(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	_, err = fmt.Fprint(w, "record succesfully deleted")
	if err != nil {
		log.Print(err)
	}

}
