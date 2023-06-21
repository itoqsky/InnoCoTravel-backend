package service

import (
	"github.com/itoqsky/InnoCoTravel_backend/internal/core"
	"github.com/itoqsky/InnoCoTravel_backend/internal/repository"
)

type TripService struct {
	repo repository.Trip
}

func NewTripService(repo repository.Trip) *TripService {
	return &TripService{repo: repo}
}

func (s *TripService) Create(trip core.Trip) (int, error) {
	return s.repo.Create(trip)
}

func (s *TripService) GetById(userId, tripId int) (core.Trip, error) {
	return s.repo.GetById(userId, tripId)
}

func (s *TripService) Delete(userId, tripId int) error {
	return s.repo.Delete(userId, tripId)
}

func (s *TripService) Update(trip core.Trip) error {
	return s.repo.Update(trip)
}
