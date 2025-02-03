package storage

import (
	"testing"
)

func TestStorage_Open(t *testing.T) {
	storage, err := Open("users.bin")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Open file via storage", func(t *testing.T) {
		if storage == nil {
			t.Fatal("storage is nil")
		}
	})

	defer func(storage *Storage) {
		err := storage.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(storage)
}
