package utils

import "log"

func MustInit[T any](initializer func() (T, error), name string) T {
	instance, err := initializer()
	if err != nil {
		log.Fatalf("Error initializing %s: %v", name, err)
	}
	return instance
}
