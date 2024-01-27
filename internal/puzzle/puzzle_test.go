package puzzle_test

import (
	"os"
	"testing"

	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/stretchr/testify/assert"
)

func TestRead_EmptyFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "TestRead_EmptyFile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	want := ""
	result, err := puzzle.Read(tempFile.Name())

	assert.Nil(t, err)
	assert.Equal(t, want, result)
}

func TestRead_Puzzle(t *testing.T) {
	tempFile, err := os.CreateTemp("", "TestRead_EmptyFile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte("1, 2, 3, 4\n"))
	if err != nil {
		t.Fatal(err)
	}

	want := "1, 2, 3, 4\n"
	result, err := puzzle.Read(tempFile.Name())

	assert.Nil(t, err)
	assert.Equal(t, want, result)
}

func TestRead_NonExistent(t *testing.T) {
	_, err := puzzle.Read("donotexists")

	assert.Error(t, err)
	assert.ErrorContains(t, err, "no such file or directory")
	assert.ErrorContains(t, err, "Read:")
}

func TestRead_NoReadPermission(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Change the file permissions to remove the read permission
	err = os.Chmod(tempFile.Name(), 0222)
	if err != nil {
		t.Fatal(err)
	}

	_, err = puzzle.Read(tempFile.Name())
	assert.Error(t, err)
	assert.ErrorContains(t, err, "permission denied")
	assert.ErrorContains(t, err, "Read:")
}
