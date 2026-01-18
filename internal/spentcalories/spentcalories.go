package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parsed := strings.Split(data, ",")
	if len(parsed) != 3 {
		return 0, "", 0, errors.New("Некорректный ввод данных")
	}

	steps, err := strconv.Atoi(parsed[0])

	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("Некорректный ввод данных")
	}

	trainingType := parsed[1]

	trainingTime, err := time.ParseDuration(parsed[2])
	if err != nil {
		return 0, "", 0, err
	}

	if trainingTime <= 0 {
		return 0, "", 0, errors.New("Некорректный ввод данных")
	}
	return steps, trainingType, trainingTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepLength := stepLengthCoefficient * height
	distance := stepLength * float64(steps) / mInKm

	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	hours := duration.Hours()
	if hours <= 0 {
		return 0.0
	}
	dist := distance(steps, height)

	return dist / hours

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	steps, trainType, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var spendCalories float64

	switch trainType {
	case "Бег":
		spendCalories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

	case "Ходьба":
		spendCalories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	distance := distance(steps, height)

	speed := meanSpeed(steps, height, duration)

	//Были ли способы сделать это более читабельным?
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		trainType,
		duration.Hours(),
		distance,
		speed,
		spendCalories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	time := duration.Minutes()

	if time <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("Некорректный ввод данных")
	}

	speed := meanSpeed(steps, height, duration)

	return (weight * speed * time) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	time := duration.Minutes()

	if time <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("Некорректный ввод данных")
	}

	speed := meanSpeed(steps, height, duration)

	return (weight * speed * time) / minInH * walkingCaloriesCoefficient, nil
}
