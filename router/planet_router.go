package planetrouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/supaidaman/go-restapi/dao"
	. "github.com/supaidaman/go-restapi/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = PlanetsDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	planets, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, planets)
}

func GetPlanetsByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planets, err := dao.GetPlanetsByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, planets)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Planet ID")
		return
	}
	respondWithJson(w, http.StatusOK, planet)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet MyPlanet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	planet.ID = bson.NewObjectId()
	if err := dao.Create(planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, planet)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet MyPlanet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": planet.Name + " updated!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
