package poker

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Peppers's score", func(t *testing.T) {
		rq := newGetScoreRequest("Pepper")
		rs := httptest.NewRecorder()

		server.ServeHTTP(rs, rq)

		assert.Equal(t, http.StatusOK, rs.Code)
		assert.Equal(t, "20", rs.Body.String())
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		rq := newGetScoreRequest("Apollo")
		rs := httptest.NewRecorder()

		server.ServeHTTP(rs, rq)

		assert.Equal(t, http.StatusNotFound, rs.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records a win on POST", func(t *testing.T) {
		player := "Pepper"
		rq := newPostWinRequest(player)
		rs := httptest.NewRecorder()

		server.ServeHTTP(rs, rq)

		assert.Equal(t, http.StatusAccepted, rs.Code)
		assert.Equal(t, 1, len(store.winCalls))
		assert.Equal(t, store.winCalls[0], player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		rq := newLeagueRequest()
		rs := httptest.NewRecorder()

		server.ServeHTTP(rs, rq)

		got := getLeagueFromResponse(t, rs.Body)

		assert.Equal(t, http.StatusOK, rs.Code)
		assert.Equal(t, wantedLeague, got)
		assert.Equal(t, jsonContentType, rs.Result().Header.Get("content-type"), "response did not have content-type of application/json")
	})
}

func newGetScoreRequest(player string) *http.Request {
	rq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return rq
}

func newPostWinRequest(player string) *http.Request {
	rq, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return rq
}

func newLeagueRequest() *http.Request {
	rq, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return rq
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	assert.NoError(t, err)

	return
}
