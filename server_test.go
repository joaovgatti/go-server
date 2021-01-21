package go_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
/*
	Next, we want to check that when we do our POST /players/{name} that our PlayerStore is told to record the win.
*/
func (s *StubPlayerStore) RecordWin(name string) int{
	score := s.scores[name]
	return score
}

func TestGetPlayers(t *testing.T){

	store := StubPlayerStore{
		map[string]int{
			"Neymar":20,
			"Messi":30,
		},
	}
	server := &PlayServer{&store}

	t.Run("return Neymar's score", func (t *testing.T){
		request := newGetScoreRequest("Neymar")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return Messi score", func (t *testing.T){
		request := newGetScoreRequest("Messi")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "30")

	})
	t.Run("returns 404 on missing players", func (t *testing.T){
		request := newGetScoreRequest("Kaka")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T){
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayServer{&store}

	t.Run("it returns accepted on POST", func (t *testing.T){
		request, _ := http.NewRequest(http.MethodPost, "/players/Neymar",nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	} )
}

func newGetScoreRequest(name string) *http.Request{
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name),nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got,want)
	}
}

func assertStatus(t testing.TB, got, want int){
	t.Helper()
	if got != want {
		t.Errorf("status error, got %q want %q",got, want)
	}
}