package service

import (
	"context"

	"github.com/vmamchur/vacancy-board/internal/model"
	"github.com/vmamchur/vacancy-board/internal/repository"
)

type VacancyService struct {
	vacancyRepository repository.VacancyRepository
}

func NewVacancyService(
	vacancyRepository repository.VacancyRepository,
) *VacancyService {
	return &VacancyService{vacancyRepository: vacancyRepository}
}

func (s *VacancyService) GetAll(ctx context.Context, filter model.GetVacanciesFilter) ([]model.Vacancy, error) {
	vacancies, err := s.vacancyRepository.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}
