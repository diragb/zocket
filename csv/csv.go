package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type Song struct {
	Name  string
	Album string
}

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var songs []Song
	for _, line := range data {
		var song Song
		for index, value := range line {
			if index == 0 {
				song.Name = value
			} else if index == 1 {
				song.Album = value
			}
		}
		songs = append(songs, song)
	}
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(writer, "Name\t Album")
	fmt.Fprintln(writer, "====\t =====")
	for index := range songs {
		fmt.Fprintln(writer, songs[index].Name, "\t", songs[index].Album)
	}
	writer.Flush()
}
