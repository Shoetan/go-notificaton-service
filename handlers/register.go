package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Shoetan/pkg/rabbitmq"
	"github.com/Shoetan/utils"
	"github.com/jmoiron/sqlx"
)

 type payload struct {
	Fname string `json:"first_name"`
	Lname string `json:"last_name"`
	Email string `json:"email_address"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"inserted_at"`
 }

 type response struct {
	UserId int `json:"id"`
	Fname string `json:"first_name"`
	Lname string `json:"last_name"`
	Email string  `json:"email_address"`
	CreatedAt time.Time `json:"inserted_at"`
 }


func RegisterUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {

		var payload payload
		var exitingEmail string
		var userID int

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest )
		}

		hashedPassword, err :=utils.HashPwd(payload.Password)

		if err != nil {
			log.Fatal("could not hash password")
		}

		payload.Password = hashedPassword
		

		err = db.Get(&exitingEmail, "SELECT email_address FROM users WHERE email_address = $1", payload.Email)

		switch {
		case err == sql.ErrNoRows:
			
		case err != nil:

			log.Println("Could not get email from database")
			return
		
		default:
			http.Error(w, "Email already taken", http.StatusBadRequest)
			return
		}

		payload.CreatedAt =  time.Now()

		err = db.QueryRow("INSERT INTO users (first_name, last_name, email_address, password, inserted_at) VALUES ($1, $2, $3, $4, $5) returning id", payload.Fname, payload.Lname, payload.Email, payload.Password, payload.CreatedAt ).Scan(&userID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
		}
		

		response := response{
			UserId: userID,
			Fname: payload.Fname,
			Lname: payload.Lname,
			Email: payload.Email,
			CreatedAt: payload.CreatedAt,
		}

		conn, err := rabbitmq.RabbitMqConn()

		if err != nil {
			log.Println("Could not connect to RabbitMq")
		}

		err = rabbitmq.PublishToQueue(conn)

		if err != nil {
			log.Println("Faild to publish message to rabbitmq")
		} else {
			log.Println("Published message")
		}

		json.NewEncoder(w).Encode(response)
	}
	
}