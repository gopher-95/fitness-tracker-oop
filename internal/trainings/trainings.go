package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	//делим строку на слайс строк
	slice := strings.Split(datastring, ",")

	//проверка на длину слайса
	if len(slice) != 3 {
		return errors.New("длина слайса не равна 3")
	}

	//преобразовываем первый элемент в тип int
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("не удалось сделать преобразование %w", err)
	}

	//проверка на количество шагов
	if steps <= 0 {
		return errors.New("количество шагов <= 0")
	}

	//поместили в структуру преобразованные шаги
	t.Steps = steps

	//добавили в структуру тип тренировки
	t.TrainingType = slice[1]

	//преобразовываем третий элемент слайса в тип time.Duration
	time, err := time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("не удалось сделать преобразование %w", err)
	}

	//проверка на время
	if time <= 0 {
		return errors.New("продолжительность <= 0")
	}

	t.Duration = time
	return nil

}

func (t Training) ActionInfo() (string, error) {
	//создали ошибку для неизвестной тренировки
	var err error = errors.New("неизвестный тип тренировки")

	//вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	//вычисляем среднню скорость
	speedWalking := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	if t.TrainingType == "Ходьба" {
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speedWalking, calories), nil

	}

	if t.TrainingType == "Бег" {
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speedWalking, calories), nil
	}

	return "", err
}
