package go_server

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {}

func(i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main(){
	server := &PlayServer{&InMemoryPlayerStore{}}


	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v",err)
	}
}