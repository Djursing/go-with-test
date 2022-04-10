package poker

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSystemStore(t *testing.T) {
	player := "Chris"

	t.Run("league from a reader", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assert.NoError(t, err)

		got := store.GetLeague()
		exp := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.EqualValues(t, exp, got)
		// Run test again to make sure we read from top of file each time
		got = store.GetLeague()
		assert.EqualValues(t, exp, got)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assert.NoError(t, err)

		got := store.GetPlayerScore(player)
		exp := 33

		assert.Equal(t, exp, got)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assert.NoError(t, err)

		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		exp := 34

		assert.Equal(t, exp, got)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
			{"Name": "Cleo", "Wiies": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assert.NoError(t, err)

		player := "Pepper"
		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		exp := 1

		assert.Equal(t, exp, got)
	})

	t.Run("Works with an empty file", func(t *testing.T) {
		db, cleanDb := createTempFile(t, "")
		defer cleanDb()

		_, err := NewFileSystemPlayerStore(db)
		assert.NoError(t, err)
	})
}

func createTempFile(t testing.TB, initData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	assert.NoError(t, err)

	tmpfile.Write([]byte(initData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
