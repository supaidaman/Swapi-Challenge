package dao

import (
	"testing"

	"github.com/supaidaman/go-restapi/models"
)

func TestFindPlanet(t *testing.T) {

	testFindFilms, err := FindFilmsByPlanet("Tatooine")

	if err != nil {
		t.Errorf("FindFilmsByPlanet failed with error")
	}

	t.Logf("Completed with Success, found %d movies", len(testFindFilms))

}

func TestGetAll(t *testing.T) {

	var dao = PlanetsDAO{}
	dao.Connect()
	result, err := dao.GetAll()

	if err != nil {
		t.Errorf("GetAll failed with error")
	}

	t.Logf("Completed with Success, found %d Planets", len(result))

}

func (p *PlanetsDAO) CreateMock(planet models.MyPlanet) error {
	return nil
}

func (p *PlanetsDAO) GetByIDMock(id string) (models.MyPlanet, error) {
	return models.MyPlanet{Name: "Coruscant", Climate: "Dry", Terrain: "Rocky"}, nil
}

func (p *PlanetsDAO) DeleteMock(id string) error {
	return nil
}

func TestInsertAndDelete(t *testing.T) {

	var dao = PlanetsDAO{}
	dao.Connect()

	var myPlanet = models.MyPlanet{Name: "Coruscant", Climate: "Dry", Terrain: "Rocky"}

	err := dao.CreateMock(myPlanet)
	if err != nil {
		t.Errorf("Create failed with error")
	}

	queryPlanet, queryErr := dao.GetByIDMock("somerandomIDReturned")

	if queryErr != nil {
		t.Errorf("Query failed with error")
	}

	t.Logf(queryPlanet.Name)
	deleteErr := dao.DeleteMock("somerandomIDReturned")

	if deleteErr != nil {
		t.Errorf("Delete failed with error")
	}

	t.Logf("Completed with Success, mocked")
}
