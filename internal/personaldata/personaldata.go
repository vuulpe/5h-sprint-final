package personaldata

import "fmt"

// Pursonal structure
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// metod Print()
func (p Personal) Print() string {
	return fmt.Sprintf("Имя: %s\nВес: %.1f\nРост: %.1f", p.Name, p.Weight, p.Height)
}
