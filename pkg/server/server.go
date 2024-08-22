package server

import (
	"log"
	"net/http"

	"github.com/Shoetan/database"
	"github.com/Shoetan/handlers"
	"github.com/Shoetan/models"
)

type APISERVER struct {
	addr string
}

func NewAPISERVER(addr string) *APISERVER  {
	return &APISERVER{
		addr: addr,
	}
}

func (s* APISERVER) Run() error {


	db, err := database.Database()

	if err != nil {
		log.Fatal("Could not not connect")
	}

	
	db.MustExec(models.TABLES)

router := http.NewServeMux()

//add enpoints here
router.HandleFunc("POST /register", handlers.RegisterUser(db))



server := &http.Server{
	Addr: s.addr,
	Handler: router,
}

log.Printf("The server is running on port: %s", s.addr)
return server.ListenAndServe()
	
}