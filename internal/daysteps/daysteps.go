package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vuulpe/5h-sprint-final/internal/personaldata"
	"github.com/vuulpe/5h-sprint-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// structure DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// metod Parse()
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных, ожидается 'шаги,длительность'")
	}

	// steps parsing
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	// duration parsing
	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("неверный формат длительности: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

// metod ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duretion mast be positive")
	}

	distance := spentenergy.Distance(ds.Steps)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error calories counting: %v", err)
	}

	info := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	info += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	info += fmt.Sprintf("Вы сожгли %.2f ккал.", calories)

	return info, nil
}
