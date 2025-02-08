package storage

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStorage_Open(t *testing.T) {
	storage, err := Open("users.db")
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

func TestStorage_Write(t *testing.T) {
	storage, err := Open("users.db")
	if err != nil {
		t.Fatal(err)
	}

	err = storage.Write([]byte(hex.EncodeToString([]byte("hello world"))))

	t.Run("Write file via storage", func(t *testing.T) {
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}
	})

	defer func(storage *Storage) {
		err := storage.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(storage)
}

func TestStorage_Read(t *testing.T) {
	storage, err := Open("users.db")
	if err != nil {
		t.Fatal(err)
	}

	res, err := hex.DecodeString(string(storage.Read()))

	t.Run("Read file via storage", func(t *testing.T) {
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !bytes.Equal(res, []byte("hello world")) {
			t.Fatalf("expected %s, got %s", []byte("hello world"), res)
		}
	})

}
