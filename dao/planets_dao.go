package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peterhellberg/swapi"
	. "github.com/supaidaman/go-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PlanetsDAO struct
type PlanetsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	//COLLECTION of planets
	COLLECTION = "planets"
)

//Connect to database
func (m *PlanetsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//FindFilmsByPlanet Retrieves Films for a planet.
func FindFilmsByPlanet(name string) ([]swapi.FilmURL, error) {

	res, getErr := http.Get(fmt.Sprintf("https://swapi.co/api/planets/?search=%s", name))
	if getErr != nil {
		log.Fatalln("http.Get error: ", getErr)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalln("Read Error: ", readErr)
	}

	var data SearchFilmResult
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	if len(data.Results) > 0 {
		return data.Results[0].FilmURLs, nil
	}

	return nil, nil

}

//GetPlanetsByName get planets list with similar names
func (m *PlanetsDAO) GetPlanetsByName(name string) ([]MyPlanet, error) {
	var planets []MyPlanet

	err := db.C(COLLECTION).Find(bson.M{"name": name}).All(&planets)
	return planets, err

}

//GetAll -> get all planets
func (m *PlanetsDAO) GetAll() ([]MyPlanet, error) {
	var planets []MyPlanet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

//GetByID -> Get a single planet
func (m *PlanetsDAO) GetByID(id string) (MyPlanet, error) {
	var planet MyPlanet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

//Create Planet
func (m *PlanetsDAO) Create(planet MyPlanet) error {

	movieResult, movError := FindFilmsByPlanet(planet.Name)
	if movError != nil {
		print("Non existant planet, no movies avaiable")
	} else {
		planet.Movies = make([]string, len(movieResult))
		for i := 0; i < len(movieResult); i++ {
			planet.Movies[i] = string(movieResult[i])
		}

	}

	err := db.C(COLLECTION).Insert(&planet)

	return err
}

//Delete Planet
func (m *PlanetsDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

//Update Planet
func (m *PlanetsDAO) Update(id string, planet MyPlanet) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &planet)
	return err
}
