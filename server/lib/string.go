package lib

import "github.com/google/uuid"

func NewStringID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
