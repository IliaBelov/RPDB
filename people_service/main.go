package main

import (
	"fmt"

	"github.com/IliaBelov/RPDB/tree/people_service/people_service/service/store"
)

func main() {
	/*ctxwt, cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctxwt,"postgress://belov:ghbrjks02082002@http://95.217.232.188:7777/belov")
	if err != nil{
		fmt.Fprint(os.Stderr,"Unable to connect to database: %v\n",err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())*/

	conn := "postgress://belov@95.217.232.188:7777/belov"
	s := store.NewStore(conn)
	fmt.Printf(s.GetPeopleByID())
}
