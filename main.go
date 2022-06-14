package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gpsandyka/manajemen_kos/controller"
	"github.com/gpsandyka/manajemen_kos/repository"
	"github.com/gpsandyka/manajemen_kos/services"
)

func startCode() {
	db, err := repository.NewDBConnection()
	if err != nil {
		fmt.Println("Failed Connection, maybe forgot to sudo service postgresql start?")
		os.Exit(1)
	}

	ur := repository.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controller.NewUserController(us)
	controller.RunServer(uc, os.Getenv("PORT"))

	us.StartDayChangeRoutine()
}

func waitCode() {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
}

func closeCode() {
	controller.CloseServer()
	os.Exit(0)
}

func testingCode() {
	db, err := repository.NewDBConnection()
	if err != nil {
		fmt.Println("Failed Connection, maybe forgot to sudo service postgresql start?")
		os.Exit(1)
	}

	ur := repository.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controller.NewUserController(us)
	controller.RunServer(uc, os.Getenv("PORT"))

	us.StartDayChangeRoutine()

	// tp := template.Must(template.ParseFiles("dummyhtml/residentlogin.html"))

	// tp.Execute(os.Stdout, nil)

}

func main() {
	//Harus pake command "go run main.go" ya

	// fmt.Println(time.Now())

	// testingCode()

	startCode()
	// fmt.Println()
	waitCode()
	// fmt.Println()
	closeCode()

}
