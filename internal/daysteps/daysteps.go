package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	parsed := strings.Split(data, ",")
	if len(parsed) != 2 {
		return 0, 0, errors.New("Некорректный ввод данных")
	}

	steps, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, errors.New("Некорректный ввод данных")
	}

	walkTime, err := time.ParseDuration(parsed[1])
	if err != nil {
		return 0, 0, err
	}

	if walkTime <= 0 {
		return 0, 0, errors.New("Некорректный ввод данных")
	}
	return steps, walkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkTime, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	walkDistance := float64(steps) * stepLength / mInKm

	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, walkDistance, spentCalories,
	)
}
