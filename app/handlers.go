package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"larssonoliver.com/lnkshrt/app/helpers"
	"larssonoliver.com/lnkshrt/app/models"
)

func (a *App) CreateLink(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	status := http.StatusCreated

	if err != nil {
		status = http.StatusBadRequest
	}

	var link models.Link

	if json.Unmarshal(body, &link) != nil || link.Url == "" {
		status = http.StatusBadRequest
	}

	if status != http.StatusCreated {
		http.Error(w, http.StatusText(status), status)
		return
	}

	for {
		link.Id = helpers.NewId()
		if _, found := a.Database.Get(link.Id); !found {
			break
		}
	}

	a.Database.Set(link.Id, link.Url)
	json.NewEncoder(w).Encode(link)
}

func (a *App) ResolveLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	link, found := a.Database.Get(id)

	if !found {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}
