package resourceid

import "github.com/google/uuid"

type Generator interface {
	Next() (string, error)
}

func Next() string {
	if id, err := SonyFlakeGenerator().Next(); err == nil {
		return id
	}

	return uuid.New().String()
}
