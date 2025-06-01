package scraper

import (
	"log"

	"github.com/vmamchur/vacancy-board/internal/repository"
)

type Scraper interface {
	Scrape() error
}

type ScraperService struct {
	vacancyRepository repository.VacancyRepository
	scrapers          []Scraper
}

func NewScraper(vacancyRepository repository.VacancyRepository, djEmail string, djPassword string) *ScraperService {
	return &ScraperService{
		vacancyRepository: vacancyRepository,
		scrapers: []Scraper{
			DjinniScraper{vacancyRepository: vacancyRepository, email: djEmail, password: djPassword},
		},
	}
}

func (s *ScraperService) Run() {
	for _, scr := range s.scrapers {
		err := scr.Scrape()
		if err != nil {
			log.Printf("Error scraping: %v\n", err)
			continue
		}
	}
}
