package storage

import (
	"os"
)

type Storage struct {
	file *os.File
}

func Open(filename string) (*Storage, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &Storage{file: f}, nil
}

func (s *Storage) Close() error {
	return s.file.Close()
}
