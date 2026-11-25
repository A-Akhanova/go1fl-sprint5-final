package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, record := range dataset {
		if err := dp.Parse(record); err != nil {
			log.Printf("не удалось разобрать запись %q: %v", record, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось сформировать информацию об активности %q: %v", record, err)
			continue
		}
		fmt.Println(strings.TrimSuffix(info, "\n"))
	}
}
