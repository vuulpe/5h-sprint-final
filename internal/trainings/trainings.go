package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vuulpe/5h-sprint-final/internal/personaldata"
	"github.com/vuulpe/5h-sprint-final/internal/spentenergy"
)

// structure Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// metod Parse()
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format")
	}

	// steps parsing
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %v", err)
	}
	t.Steps = steps

	// training type check up
	trainingType := strings.TrimSpace(parts[1])
	if trainingType != "Бег" && trainingType != "Ходьба" {
		return fmt.Errorf("unknown training type: %s", trainingType)
	}
	t.TrainingType = trainingType

	// duration parsing
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", fmt.Errorf("duration must be positive")
	}

	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция: %.2f км.\n Скорость: %.2f км/ч\n Сожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return info, nil
}
