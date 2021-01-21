package go_server

import (
	"fmt"
	"net/http"
	"strings"
)
/*
This is like a database
*/
type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayServer struct {
	store PlayerStore
}

func (p *PlayServer) ServeHTTP(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w,r)
	}
}

func (p *PlayServer) showScore(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path,"/players/")
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w,score)
}

func (p *PlayServer) processWin(w http.ResponseWriter){
	w.WriteHeader(http.StatusAccepted)
}




func getPlayerScore(name string) string {
	if name == "Neymar"{
		return "20"
	}
	if name == "Messi"{
		return "30"
	}
	return ""
}