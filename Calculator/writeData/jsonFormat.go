package writeData

import (
	"encoding/json"
	"errors"
	"os"
)

func WriteToJson(path string, data interface{}) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data.")
	}

	return nil
}
