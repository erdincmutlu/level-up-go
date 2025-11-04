package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

type entry struct {
	song       Song
	albumIndex int
	songIndex  int
}

// An PlaylistHeap is a max-heap of PlaylistEntries
type PlaylistHeap []entry

func (h PlaylistHeap) Len() int {
	return len(h)
}

func (h PlaylistHeap) Less(i, j int) bool {
	// We want Pop to return the highest play count
	return h[i].song.PlayCount > h[j].song.PlayCount
}

func (h PlaylistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PlaylistHeap) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *PlaylistHeap) Pop() any {
	original := *h
	n := len(original)
	x := original[n-1]
	*h = original[0 : n-1]
	return x
}

// makePlaylist makes the merged sorted list of songs
//
// What’s Happening Behind the Scenes:
// Initially insert the first song from each album to the heap.
// Each time you pop, you remove the song with the highest play count.
// You then replace it with the next song from that same album.
// The heap always contains one song from each album → the best remaining song.
// This is a k-way merge using a max-heap.
// (If there were k albums, the heap size would always be k until albums run out.)
func makePlaylist(albums [][]Song) []Song {
	if len(albums) == 0 {
		return nil
	}

	// initialize the heap and add first of each album, since they are the max
	pHeap := &PlaylistHeap{}
	heap.Init(pHeap)
	for i, album := range albums {
		heap.Push(pHeap, entry{album[0], i, 0})
	}

	var playlist []Song
	for pHeap.Len() > 0 {
		// take max elem from the list
		e := heap.Pop(pHeap).(entry)
		playlist = append(playlist, e.song)

		// the next song after the max is a good candidate to look at
		next := e.songIndex + 1
		if next < len(albums[e.albumIndex]) {
			heap.Push(pHeap, entry{albums[e.albumIndex][next], e.albumIndex, next})
		}
	}

	return playlist
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
