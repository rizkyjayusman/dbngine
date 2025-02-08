package storage

import (
	"log"
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

func (s *Storage) Write(val []byte) error {
	return os.WriteFile(s.file.Name(), val, 0666)
}

func (s *Storage) Read() []byte {
	val, err := os.ReadFile(s.file.Name())
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func (s *Storage) Close() error {
	return s.file.Close()
}
