package personaldata

import "fmt"

// Pursonal structure
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// metod Print()
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.1f\n", p.Weight)
	fmt.Printf("Рост: %.1f\n", p.Height)
}
