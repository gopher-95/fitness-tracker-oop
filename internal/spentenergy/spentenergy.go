package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("данные значения не могут быть меньше нуля")
	}

	//рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	//переводим в минуты тип duration
	durationInMin := duration.Minutes()

	//промежуточный расчет
	random := weight * speed * durationInMin

	//число калорий
	calories := random / minInH

	//поправка числа калорий
	walkingCalories := calories * walkingCaloriesCoefficient

	return walkingCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//проверка на значения, которые могут быть меньше нуля
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("данные значения не могут быть меньше нуля")
	}

	//рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	//переводим в минуты тип duration
	durationInMin := duration.Minutes()

	//промежуточный расчет
	random := weight * speed * durationInMin

	//число калорий
	calories := random / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	//проверка на duration
	if duration <= 0 {
		return 0
	}

	//вычисление дистанции с помощью функции Distance
	distance := Distance(steps, height)

	//средняя скорость
	speed := distance / duration.Hours()

	return speed
}

func Distance(steps int, height float64) float64 {
	//длина шага
	length := height * stepLengthCoefficient

	//пройденное расстояние в метрах
	distance := float64(steps) * length

	//пройденное расстояние в километрах
	distanceInKm := distance / mInKm

	return distanceInKm
}
