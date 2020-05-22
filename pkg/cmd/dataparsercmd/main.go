//
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/skyerus/history-api/pkg/api"
	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/env"
	"github.com/skyerus/history-api/pkg/event/eventrepo"
	"github.com/skyerus/history-api/pkg/models"
)

func main() {
	env.SetEnv()
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Provide the path of file as argument")
		return
	}
	path := args[0]
	historicalEvents, birthEvents, deathEvents, holidayEvents, err := parseFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := api.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	eventRepo := eventrepo.NewEventRepo(db)
	var customErr customerror.Error
	for _, historicalEvent := range *historicalEvents {
		customErr = eventRepo.SaveHistoricalEvent(&historicalEvent)
		if customErr != nil {
			fmt.Println(customErr.OriginalError())
			return
		}
	}
	for _, birthEvent := range *birthEvents {
		customErr = eventRepo.SaveBirthEvent(&birthEvent)
		if customErr != nil {
			fmt.Println(customErr.OriginalError())
			return
		}
	}
	for _, deathEvent := range *deathEvents {
		customErr = eventRepo.SaveDeathEvent(&deathEvent)
		if customErr != nil {
			fmt.Println(customErr.OriginalError())
			return
		}
	}
	for _, holidayEvent := range *holidayEvents {
		customErr = eventRepo.SaveHolidayEvent(&holidayEvent)
		if customErr != nil {
			fmt.Println(customErr.OriginalError())
			return
		}
	}
}

func parseFile(path string) (*[]models.HistoricalEvent, *[]models.BirthEvent, *[]models.DeathEvent, *[]models.HolidayEvent, error) {
	var historicalEvents []models.HistoricalEvent
	var birthEvents []models.BirthEvent
	var deathEvents []models.DeathEvent
	var holidayEvents []models.HolidayEvent
	var date *time.Time
	var err error
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		splitData := strings.Split(scanner.Text(), "	")
		if len(splitData) != 5 {
			return nil, nil, nil, nil, errors.New("Data not formatted in 5 sections")
		}
		date, err = parseDate(splitData[1], splitData[2], splitData[3])
		if err != nil {
			return nil, nil, nil, nil, err
		}
		fact := splitData[4]
		switch splitData[0] {
		case "Deaths":
			deathEvents = append(deathEvents, models.DeathEvent{Date: *date, Fact: fact})
		case "Births":
			birthEvents = append(birthEvents, models.BirthEvent{Date: *date, Fact: fact})
		case "Events":
			historicalEvents = append(historicalEvents, models.HistoricalEvent{Date: *date, Fact: fact})
		case "Holidays":
			holidayEvents = append(holidayEvents, models.HolidayEvent{Date: *date, Fact: fact})
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
