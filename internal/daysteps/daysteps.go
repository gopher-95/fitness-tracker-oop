package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	//преобразовали строку в слайс строк
	slice := strings.Split(datastring, ",")

	//проверка на длину слайса
	if len(slice) != 2 {
		return errors.New("длина слайса < 2")
	}

	//преобразуем строку в тип int
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("произошла ошибка преобразования %w", err)
	}
	if steps <= 0 {
		return errors.New("число шагов <= 0")
	}

	//записываем в структуру шаги
	ds.Steps = steps

	//преобразовываем строку в тип time.Duration
	time, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("произошла ошибка преобразования %w", err)
	}
	if time <= 0 {
		return errors.New("число шагов <= 0")
	}

	//записываем в структуру время
	ds.Duration = time

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	//вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	//вычисляем калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
