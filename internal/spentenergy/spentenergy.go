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

// Distance дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}

// MeanSpeed скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	hours := duration.Hours()

	if hours <= 0 {
		return 0
	}

	return distance / hours
}

// калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("неверный формат данных")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed <= 0 {
		return 0, errors.New("неверная скорость")
	}

	durationMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationMinutes) / minInH
	return calories, nil
}

// калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("неверный формат данных")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed <= 0 {
		return 0, errors.New("неверная скорость")
	}

	durationMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationMinutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}