package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func GetSongs() ([]Song, error) {
	file, err := os.ReadFile("pkg/db/db.json")
	if err != nil {
		log.Println("Error: cannot open db file!", err)
		return nil, errors.New("Cannot open db file!")
	}

	var songs []Song

	err = json.Unmarshal([]byte(file), &songs)

	if err != nil {
		log.Println("Error: cannot unmarshall songs from json!")
		return songs, errors.New("Cannot unmarshall songs from json!")
	}

	return songs, nil

}

func AddSong(s *Song) error {
	songs, err := GetSongs()

	if err != nil && len(songs) > 0 {
		log.Println("Error: cannot get information from GetSongs()!")
		return errors.New("Cannot get info from GetSongs()!")
	}

	for _, songInfo := range songs {
		if s.Title == songInfo.Title && s.Artist == songInfo.Artist {
			log.Println("Song already exists!")
			return errors.New("Song already exists!")
		}
	}

	songs = append(songs, *s)
	fmt.Printf("songs: %v\n", songs)

	file, err := os.Open("pkg/db/db.json")

	if err != nil {
		log.Println("Error: cannot open db.json for add song!")
		return nil
	}

	defer file.Close()

	data, err := json.Marshal(songs)

	if err != nil {
		log.Println("Error: cannot marshall song to json!")
		return nil
	}

	err = os.WriteFile(file.Name(), data, 0644)

	if err != nil {
		log.Println("Error: cannot wrtie song info to db!")
		return nil
	}

	return nil
}
