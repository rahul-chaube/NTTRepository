package service

import (
	"NTTHomeTestDemo/model"
	"fmt"
	"math"

	"github.com/google/uuid"
)

type ExoplanetInterface interface {
	CreateExoplanet(model.Exoplanet) error
	ListExoplanet() ([]model.Exoplanet, error)
	GetExoplanet(string) (model.Exoplanet, error)
	UpdateExoplanet(string, model.Exoplanet) (model.Exoplanet, error)
	DeleteExoplanet(string) error
	EstimateFuel(string) (float64, error)
}

type ExoplanetService struct {
	Exoplanets map[string]model.Exoplanet
}

func InitExpoServiceInit() *ExoplanetService {
	svc := ExoplanetService{
		Exoplanets: map[string]model.Exoplanet{},
	}
	return &svc
}

func (exo *ExoplanetService) CreateExoplanet(data model.Exoplanet) (model.Exoplanet, error) {

	_, err := data.IsValid()
	if err != nil {
		return data, err
	}

	data.ID = uuid.NewString()
	exo.Exoplanets[data.ID] = data
	return data, nil
}

func (exo *ExoplanetService) ListExoplanet() []model.Exoplanet {

	planets := []model.Exoplanet{}

	for _, planet := range exo.Exoplanets {
		planets = append(planets, planet)
	}
	return planets

}

func (exo *ExoplanetService) GetExoplanet(id string) (model.Exoplanet, error) {

	planet, isExists := exo.Exoplanets[id]
	if !isExists {
		return model.Exoplanet{}, fmt.Errorf("%s planet is not found", id)
	}
	return planet, nil

}

func (exo *ExoplanetService) UpdateExoplanet(id string, planet model.Exoplanet) (model.Exoplanet, error) {

	_, isExists := exo.Exoplanets[id]
	if !isExists {
		return model.Exoplanet{}, fmt.Errorf("%s planet is not found", id)
	}
	planet.ID = id
	exo.Exoplanets[id] = planet
	return planet, nil

}

func (exo *ExoplanetService) DeleteExoplanet(id string) (model.Exoplanet, error) {

	planet, isExists := exo.Exoplanets[id]
	if !isExists {
		return model.Exoplanet{}, fmt.Errorf("%s planet is not found", id)
	}

	delete(exo.Exoplanets, id)
	return planet, nil

}

func (exo *ExoplanetService) EstimateFuel(id string, noOfCrew int) (float64, error) {

	planet, isExists := exo.Exoplanets[id]

	if !isExists {
		return 0, fmt.Errorf("exoplanet not found")
	}

	var gravity float64

	if planet.Type == "GasGiant" {
		gravity = 0.5 / math.Pow(planet.Radius, 2)
	} else {
		gravity = planet.Mass / math.Pow(planet.Radius, 2)
	}
	fuelEstimation := float64(planet.Distance) / math.Pow(gravity, 2) * float64(noOfCrew)

	return fuelEstimation, nil

}
