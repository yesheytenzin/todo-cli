package functions

import (
	"encoding/json"
	"fmt"
	"os"
)

const file = "tasks.json"

func Save(task []Task) error {
	data, err := json.MarshalIndent(task,"", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
