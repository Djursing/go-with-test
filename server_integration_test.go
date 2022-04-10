package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, cleanDb := createTempFile(t, "[]")
	defer cleanDb()

	store, err := NewFileSystemPlayerStore(db)
	assert.NoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		rs := httptest.NewRecorder()
		server.ServeHTTP(rs, newGetScoreRequest(player))

		assert.Equal(t, http.StatusOK, rs.Code)
		assert.Equal(t, "3", rs.Body.String())

	})

	t.Run("get league", func(t *testing.T) {
		rs := httptest.NewRecorder()
		server.ServeHTTP(rs, newLeagueRequest())
		assert.Equal(t, http.StatusOK, rs.Code)

		got := getLeagueFromResponse(t, rs.Body)
		exp := []Player{{"Pepper", 3}}

		assert.Equal(t, exp, got)

	})
}
