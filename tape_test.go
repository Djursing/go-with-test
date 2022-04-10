package poker

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	exp := "abc"

	assert.Equal(t, exp, got)
}
