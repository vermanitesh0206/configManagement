package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	. "my-app/pkg/config"
	. "my-app/pkg/platform/mongo"
	"my-app/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterConfigRoutes(r)
	http.Handle("/", r)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	uname := os.Getenv("UNAME")
	pswd := os.Getenv("PSWD")
	uri := os.Getenv("URI")

	Db = os.Getenv("DB")
	Col = os.Getenv("COL")
	user := os.Getenv("USERS")

	for _, user := range strings.Split(user, ";") {
		unp := strings.Split(user, ":")
		Users[unp[0]] = unp[1]
	}

	Client, Ctx, err = Connect(fmt.Sprintf("mongodb+srv://%s:%s@%s", uname, pswd, uri))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer Close(Client, Ctx)

	log.Fatal(http.ListenAndServe(":3030", r))
}
