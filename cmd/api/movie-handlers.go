package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil{
		app.logger.Print(errors.New("invalid id parametar"))
		app.ErrorJSON(w, err)
		return
	}

	app.logger.Println("id id", id)

	movie, err := app.models.DB.Get(id)
	if err != nil{
		app.logger.Print(err)
		app.ErrorJSON(w, err)
		return
	}

	err = app.WriteJSON(w, http.StatusOK, movie, "movie")
	if err != nil{
		app.logger.Print(err)
		app.ErrorJSON(w, err)
		return
	}
	
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request){
	
}