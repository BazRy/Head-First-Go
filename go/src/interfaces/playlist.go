package main

import (
	"Head-First-Go/go/src/interfaces/gadget"
	"fmt"
	"log"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(player Player, songs []string) error {
	fmt.Println("-------------------------------------------")
	numSongs := len(songs)
	fmt.Println(NumSongs(numSongs))
	if numSongs > 3 {
		return TooManySongsError(numSongs)
	}
	for _, song := range songs {
		player.Play(song)
	}
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
	return nil
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}

	recorder := gadget.TapeRecorder{}
	err := playList(recorder, mixtape)
	if err != nil {
		log.Fatal(err)
	}

	mixtape = []string{"Push it", "Everywhere", "Jump Around"}
	player := gadget.TapePlayer{}
	err = playList(player, mixtape)
	if err != nil {
		log.Fatal(err)
	}

	mixtape = []string{"Push it", "Everywhere", "Jump Around", "My hero"}
	err = playList(player, mixtape)
	if err != nil {
		log.Fatal(err)
	}

}

type TooManySongsError int
type NumSongs int

// example satisfy the stringer interface
func (n NumSongs) String() string {
	return fmt.Sprintf("Playlist has %d songs", n)
}

// example satisfy the error interface
func (t TooManySongsError) Error() string {
	return fmt.Sprintf("%d songs is too many", t)
}
