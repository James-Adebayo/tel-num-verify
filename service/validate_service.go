package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ValidateNumber(number string) (map[string]any, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Failed to load ENV")
	}

	if number == "" {
		return nil, errors.New("Phone number empty")
	}

	id, err := strconv.Atoi(number)
	if err != nil {
		return nil, errors.New("Phone number is not a digit")
	}

	url := fmt.Sprintf(
		"%s/validate?access_key=%s&number=%d",
		os.Getenv("NUM_BASE_URL"),
		os.Getenv("NUM_ACCESS_KEY"),
		id,
	)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Failed to validate number")
	}

	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, errors.New("Failed to read response")
	}

	return result, nil
}
