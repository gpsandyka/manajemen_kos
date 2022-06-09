package controller

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	//"os"
	//"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func NewRouter(u UserController) (r *mux.Router) {
	r = mux.NewRouter()

	// HANDLE LIST
	//CRUD RESIDENT HANDLER
	r.HandleFunc("/residents", u.ReadAllResidents).Methods("GET")
	r.HandleFunc("/residents/{id}", u.ReadResident).Methods("GET")
	r.HandleFunc("/residents", u.CreateResident).Methods("POST")
	r.HandleFunc("/residents/{id}", u.UpdateResident).Methods("PUT")
	r.HandleFunc("/residents/{id}", u.DeleteResident).Methods("DELETE")
	//CRUD ROOMS HANDLER
	r.HandleFunc("/rooms", u.ReadAllRooms).Methods("GET")
	r.HandleFunc("/rooms/{id}", u.ReadRoom).Methods("GET")
	r.HandleFunc("/rooms", u.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms/{id}", u.UpdateRoom).Methods("PUT")
	r.HandleFunc("/rooms/{id}", u.DeleteRoom).Methods("DELETE")
	//CRUD INVOICES HANDLER
	r.HandleFunc("/invoices", u.ReadAllInvoices).Methods("GET")
	r.HandleFunc("/invoices/{id}", u.ReadInvoice).Methods("GET")
	r.HandleFunc("/invoices", u.CreateInvoice).Methods("POST")
	r.HandleFunc("/invoices/{id}", u.UpdateInvoice).Methods("PUT")
	r.HandleFunc("/invoices/{id}", u.DeleteInvoice).Methods("DELETE")
	//CRUD ADMIN HANDLER
	r.HandleFunc("/admins", u.ReadAllAdmins).Methods("GET")
	r.HandleFunc("/admins/{id}", u.ReadAdmin).Methods("GET")
	r.HandleFunc("/admins", u.CreateAdmin).Methods("POST")
	r.HandleFunc("/admins/{id}", u.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admins/{id}", u.DeleteAdmin).Methods("DELETE")
	//CRUD NOTIFICATION HANDLER
	r.HandleFunc("/notifications", u.ReadAllNotifications).Methods("GET")
	r.HandleFunc("/notifications/{id}", u.ReadNotification).Methods("GET")
	r.HandleFunc("/notifications", u.CreateNotification).Methods("POST")
	r.HandleFunc("/notifications/{id}", u.UpdateNotification).Methods("PUT")
	r.HandleFunc("/notifications/{id}", u.DeleteNotification).Methods("DELETE")
	//CRUD PROBLEM HANDLER
	r.HandleFunc("/problems", u.ReadAllProblems).Methods("GET")
	r.HandleFunc("/problems/{id}", u.ReadProblem).Methods("GET")
	r.HandleFunc("/problems", u.CreateProblem).Methods("POST")
	r.HandleFunc("/problems/{id}", u.UpdateProblem).Methods("PUT")
	r.HandleFunc("/problems/{id}", u.DeleteProblem).Methods("DELETE")

	//LOGIN HANDLER
	r.HandleFunc("/admins/login", u.LoginAdmin).Methods("POST")
	r.HandleFunc("/residents/login", u.LoginResident).Methods("POST")
	// r.HandleFunc("/admins/loginbyjson", u.LoginAdminByJSON).Methods("POST")
	// r.HandleFunc("/residents/loginbyjson", u.LoginResidentByJSON).Methods("POST")

	//BEFORE LOGIN HANDLER
	r.HandleFunc("/", u.AreYouLoggedInAsResident).Methods("GET")
	r.HandleFunc("/AdminOnly", u.AreYouLoggedInAsAdmin).Methods("GET")
	r.HandleFunc("/login", u.ResidentLoginForm).Methods("GET")
	r.HandleFunc("/AdminOnly/login", u.AdminLoginForm).Methods("GET")

	//RESIDENT FUNCTION HANDLER

	//ADMIN FUNCTION HANDLER
	r.HandleFunc("/AdminOnly/solve-notification", u.SolveNotificationPage).Methods("GET")
	r.HandleFunc("/AdminOnly/solve-notification", u.SolveNotification).Methods("POST")
	r.HandleFunc("/AdminOnly/check-in-room", u.CheckInPage).Methods("GET")
	r.HandleFunc("/AdminOnly/check-in-room/room-pick", u.CheckInRoomPickPage).Methods("GET")
	r.HandleFunc("/AdminOnly/check-in-room/gather-information", u.CheckInRoomGatherPage).Methods("GET")
	r.HandleFunc("/AdminOnly/check-in-room/fill-in-information", u.CheckInRoomFillinPage).Methods("POST")
	r.HandleFunc("/AdminOnly/check-in-room/complete", u.CheckInRoom).Methods("POST")
	//
	r.HandleFunc("/AdminOnly/check-out-room", u.CheckOutPage).Methods("GET")
	r.HandleFunc("/AdminOnly/check-out-room/confirmation", u.CheckOutConfirmationPage).Methods("GET")
	r.HandleFunc("/AdminOnly/check-out-room/confirmation", u.CheckOutRoom).Methods("POST")
	r.HandleFunc("/AdminOnly/ready-up-room", u.ReadyUpRoomPage).Methods("GET")
	r.HandleFunc("/AdminOnly/ready-up-room", u.ReadyUpRoom).Methods("POST")
	r.HandleFunc("/AdminOnly/information", u.InformationPage).Methods("GET")

	return
}

var wait time.Duration
var srv *http.Server

func RunServer(u UserController, port string) {

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := NewRouter(u)

	fmt.Println("Okay router ready")

	//For testing purpose only

	srv = &http.Server{
		// Addr: "localhost:8080", //<<<< Localhost testing
		Addr: ":" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		// Handler: csrf.Protect([]byte("32-byte-long-auth-key"),
		// 	csrf.Path("/"),
		// 	csrf.Secure(false))(r),
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Should be online")

	//Cookie Option
	store.Options.MaxAge = 3600 //3600 second a.k.a 1 hour
	store.Options.HttpOnly = true
}

func CloseServer() {
	fmt.Println("This message appear if you close it")
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
