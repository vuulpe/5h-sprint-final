package actioninfo

import (
	"fmt"
)

// interface DataParser
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// function Info()
func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {

		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("error of name %d: %v\n", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("error of forming information %d: %v\n", i+1, err)
			continue
		}

		fmt.Println(info)
	}
}
