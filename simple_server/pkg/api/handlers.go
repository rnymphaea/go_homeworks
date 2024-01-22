package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_server/pkg/db"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	songs, err := db.GetSongs()

	if err != nil && len(songs) == 0 {
		w.Write([]byte("Internal server error!\n"))
	}

	switch r.Method {
	case http.MethodGet:
		for _, songInfo := range songs {
			fmt.Fprintf(w, "'%s' by %s\n", songInfo.Title, songInfo.Artist)
		}

	case http.MethodPost:
		data, err := io.ReadAll(r.Body)

		if err != nil {
			w.Write([]byte("Плохое тело запроса!\n"))
			return
		}

		song := db.Song{}

		if err = json.Unmarshal(data, &song); err != nil {
			log.Println(err)
			w.Write([]byte("Произошла ошибка при декодировании запроса!\n"))
			return
		}

		err = db.AddSong(&song)

		if err != nil {
			w.Write([]byte("Произошла ошибка при попытке записать песню в базу данных!\n"))
			return
		}

		w.Write([]byte("Песня успешно добавлена!\n"))
		return
	}
}
