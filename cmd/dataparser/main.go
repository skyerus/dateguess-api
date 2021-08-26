package main

import (
	"bufio"
	"dateguess-api/internal/app"
	"dateguess-api/internal/model"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := app.Init()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	historicalEvents, birthEvents, deathEvents, holidayEvents, err := parseFile(a.Env.RawDataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, historicalEvent := range *historicalEvents {
		err = a.EventRepository.SaveHistoricalEvent(&historicalEvent)
		if err != nil {
			log.Fatalf("failed to save event: %v", err)
		}
	}
	for _, birthEvent := range *birthEvents {
		err = a.EventRepository.SaveBirthEvent(&birthEvent)
		if err != nil {
			log.Fatalf("failed to save event: %v", err)
		}
	}
	for _, deathEvent := range *deathEvents {
		err = a.EventRepository.SaveDeathEvent(&deathEvent)
		if err != nil {
			log.Fatalf("failed to save event: %v", err)
		}
	}
	for _, holidayEvent := range *holidayEvents {
		err = a.EventRepository.SaveHolidayEvent(&holidayEvent)
		if err != nil {
			log.Fatalf("failed to save event: %v", err)
		}
	}

	fmt.Println("Finished importing data.")
}

func parseFile(path string) (*[]model.HistoricalEvent, *[]model.BirthEvent, *[]model.DeathEvent, *[]model.HolidayEvent, error) {
	var historicalEvents []model.HistoricalEvent
	var birthEvents []model.BirthEvent
	var deathEvents []model.DeathEvent
	var holidayEvents []model.HolidayEvent
	var date *time.Time
	var err error
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		splitData := strings.Split(scanner.Text(), "	")
		if len(splitData) != 5 {
			return nil, nil, nil, nil, errors.New("data not formatted in 5 sections")
		}
		date, err = parseDate(splitData[1], splitData[2], splitData[3])
		if err != nil {
			return nil, nil, nil, nil, err
		}
		fact := splitData[4]
		switch splitData[0] {
		case "Deaths":
			deathEvents = append(deathEvents, model.DeathEvent{Date: *date, Fact: fact})
		case "Births":
			birthEvents = append(birthEvents, model.BirthEvent{Date: *date, Fact: fact})
		case "Events":
			historicalEvents = append(historicalEvents, model.HistoricalEvent{Date: *date, Fact: fact})
		case "Holidays":
			holidayEvents = append(holidayEvents, model.HolidayEvent{Date: *date, Fact: fact})
		default:
			return nil, nil, nil, nil, errors.New("Unknown category found: " + splitData[0])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, nil, err
	}

	return &historicalEvents, &birthEvents, &deathEvents, &holidayEvents, nil
}

func parseDate(year string, month string, day string) (*time.Time, error) {
	var y int
	var err error
	if year == "" {
		y = 2000
	} else {
		y, err = strconv.Atoi(year)
		if err != nil {
			return nil, err
		}
	}
	m, err := strconv.Atoi(month)
	if err != nil {
		return nil, err
	}
	d, err := strconv.Atoi(day)
	if err != nil {
		return nil, err
	}
	t := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
	return &t, nil
}
