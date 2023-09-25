package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomainC75/postgres-test/models"
	"github.com/julienschmidt/httprouter"
)

func AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Print("--------")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	// body := string(bs)

	var data models.Book
	if err := json.Unmarshal(bs, &data); err != nil {
		panic(err)
	}

	// fmt.Print("-> ", data)

	(&data).CreateBook()

	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	log.Fatalln(err)
	// }
}
