package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/larssonoliver/lnkshrt/internal/config"
	"github.com/larssonoliver/lnkshrt/internal/models"
	"github.com/larssonoliver/lnkshrt/internal/util"
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

	url, err := url.Parse(link.Url)
	if err != nil {
		status = http.StatusBadRequest
	}

	if url.Scheme == "" {
		url.Scheme = "http"
	}
	link.Url = url.String()

	if status != http.StatusCreated {
		http.Error(w, http.StatusText(status), status)
		return
	}

	link.Id, _ = a.Database.GetId(link.Url)
	// If a db error occured, it will be caught when trying to generate a new id

	if link.Id == "" {
		for {
			link.Id = util.NewId()

			if l, _ := a.Database.Get(link.Id); l == "" {
				err = a.Database.Set(link.Id, link.Url)
				if err != nil {
					status = http.StatusInternalServerError
				}
				break
			}
		}
	}

	if status != http.StatusCreated {
		http.Error(w, http.StatusText(status), status)
		return
	}

	err = json.NewEncoder(w).Encode(link)
	if err != nil {
		status = http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
	}
}

func (a *App) IndexRedirect(w http.ResponseWriter, r *http.Request) {
	// If the url is defined as an env variable, redirect to it
	if config.IndexRedirect() == "" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, config.IndexRedirect(), http.StatusTemporaryRedirect)
}

func (a *App) ResolveLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	link, err := a.Database.Get(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}
