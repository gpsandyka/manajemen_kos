package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gpsandyka/manajemen_kos/domain"
)

type userController struct {
	userService domain.UserService
	tp          *template.Template
	//fbService   service.FirebaseService
}

type UserController interface {
	CreateResident(w http.ResponseWriter, r *http.Request)
	ReadResident(w http.ResponseWriter, r *http.Request)
	ReadAllResidents(w http.ResponseWriter, r *http.Request)
	UpdateResident(w http.ResponseWriter, r *http.Request)
	DeleteResident(w http.ResponseWriter, r *http.Request)

	CreateRoom(w http.ResponseWriter, r *http.Request)
	ReadRoom(w http.ResponseWriter, r *http.Request)
	ReadAllRooms(w http.ResponseWriter, r *http.Request)
	UpdateRoom(w http.ResponseWriter, r *http.Request)
	DeleteRoom(w http.ResponseWriter, r *http.Request)

	CreateInvoice(w http.ResponseWriter, r *http.Request)
	ReadInvoice(w http.ResponseWriter, r *http.Request)
	ReadAllInvoices(w http.ResponseWriter, r *http.Request)
	UpdateInvoice(w http.ResponseWriter, r *http.Request)
	DeleteInvoice(w http.ResponseWriter, r *http.Request)

	CreateAdmin(w http.ResponseWriter, r *http.Request)
	ReadAdmin(w http.ResponseWriter, r *http.Request)
	ReadAllAdmins(w http.ResponseWriter, r *http.Request)
	UpdateAdmin(w http.ResponseWriter, r *http.Request)
	DeleteAdmin(w http.ResponseWriter, r *http.Request)

	CreateNotification(w http.ResponseWriter, r *http.Request)
	ReadNotification(w http.ResponseWriter, r *http.Request)
	ReadAllNotifications(w http.ResponseWriter, r *http.Request)
	UpdateNotification(w http.ResponseWriter, r *http.Request)
	DeleteNotification(w http.ResponseWriter, r *http.Request)

	CreateProblem(w http.ResponseWriter, r *http.Request)
	ReadProblem(w http.ResponseWriter, r *http.Request)
	ReadAllProblems(w http.ResponseWriter, r *http.Request)
	UpdateProblem(w http.ResponseWriter, r *http.Request)
	DeleteProblem(w http.ResponseWriter, r *http.Request)

	LoginAdmin(w http.ResponseWriter, r *http.Request)
	LoginResident(w http.ResponseWriter, r *http.Request)
	// LoginAdminByJSON(w http.ResponseWriter, r *http.Request)
	// LoginResidentByJSON(w http.ResponseWriter, r *http.Request)

	ResidentLoginForm(w http.ResponseWriter, r *http.Request)
	AdminLoginForm(w http.ResponseWriter, r *http.Request)

	AreYouLoggedInAsResident(w http.ResponseWriter, r *http.Request)
	AreYouLoggedInAsAdmin(w http.ResponseWriter, r *http.Request)

	SolveNotificationPage(w http.ResponseWriter, r *http.Request)
	SolveNotification(w http.ResponseWriter, r *http.Request)
	CheckInPage(w http.ResponseWriter, r *http.Request)
	CheckInRoomPickPage(w http.ResponseWriter, r *http.Request)
	CheckInRoomGatherPage(w http.ResponseWriter, r *http.Request)
	CheckInRoomFillinPage(w http.ResponseWriter, r *http.Request)
	CheckInRoom(w http.ResponseWriter, r *http.Request)
	CheckOutPage(w http.ResponseWriter, r *http.Request)
	CheckOutConfirmationPage(w http.ResponseWriter, r *http.Request)
	CheckOutRoom(w http.ResponseWriter, r *http.Request)
	ReadyUpRoomPage(w http.ResponseWriter, r *http.Request)
	ReadyUpRoom(w http.ResponseWriter, r *http.Request)
	InformationPage(w http.ResponseWriter, r *http.Request)
}

func NewUserController(s domain.UserService) UserController {
	tp := template.Must(template.ParseGlob("dummyhtml/*.html"))
	return &userController{
		userService: s,
		tp:          tp,
		//fbService:   f,
	}
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//RESIDENT CRUD CONTROLLER
func (u *userController) ReadResident(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, RESIDENTS, u)
}

func (u *userController) ReadAllResidents(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, RESIDENTS, u)
}

func (u *userController) CreateResident(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, RESIDENTS, u)
}

func (u *userController) UpdateResident(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, RESIDENTS, u)
}

func (u *userController) DeleteResident(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, RESIDENTS, u)
}

//ROOMS CRUD CONTROLLER
func (u *userController) ReadRoom(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, ROOMS, u)
}

func (u *userController) ReadAllRooms(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, ROOMS, u)
}

func (u *userController) CreateRoom(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, ROOMS, u)
}
func (u *userController) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, ROOMS, u)
}

func (u *userController) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, ROOMS, u)
}

//INVOICES CRUD CONTROLLER
func (u *userController) ReadInvoice(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, INVOICES, u)
}

func (u *userController) ReadAllInvoices(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, INVOICES, u)
}

func (u *userController) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, INVOICES, u)
}
func (u *userController) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, INVOICES, u)
}

func (u *userController) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, INVOICES, u)
}

//ADMINS CRUD CONTROLLER
func (u *userController) ReadAdmin(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, ADMINS, u)
}

func (u *userController) ReadAllAdmins(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, ADMINS, u)
}

func (u *userController) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, ADMINS, u)
}
func (u *userController) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, ADMINS, u)
}

func (u *userController) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, ADMINS, u)
}

//NOTIFICATIONS CRUD CONTROLLER
func (u *userController) ReadNotification(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, NOTIFICATIONS, u)
}

func (u *userController) ReadAllNotifications(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, NOTIFICATIONS, u)
}

func (u *userController) CreateNotification(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, NOTIFICATIONS, u)
}
func (u *userController) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, NOTIFICATIONS, u)
}

func (u *userController) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, NOTIFICATIONS, u)
}

//PROBLEMS CRUD CONTROLLER
func (u *userController) ReadProblem(w http.ResponseWriter, r *http.Request) {
	CRUDRead(w, r, PROBLEMS, u)
}

func (u *userController) ReadAllProblems(w http.ResponseWriter, r *http.Request) {
	CRUDReadAll(w, r, PROBLEMS, u)
}

func (u *userController) CreateProblem(w http.ResponseWriter, r *http.Request) {
	CRUDCreate(w, r, PROBLEMS, u)
}
func (u *userController) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	CRUDUpdate(w, r, PROBLEMS, u)
}

func (u *userController) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	CRUDDelete(w, r, PROBLEMS, u)
}

//LOG IN CONTROLLER
type Credential struct {
	Username string
	Password string
}

func (u *userController) LoginResident(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	var credential Credential

	// r.PostForm is a map of our POST form values
	err = domain.SchemaDecoder.Decode(&credential, r.PostForm)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	LoginHelper(w, r, RESIDENTS, u, &credential)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (u *userController) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	var credential Credential

	// r.PostForm is a map of our POST form values
	err = domain.SchemaDecoder.Decode(&credential, r.PostForm)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	LoginHelper(w, r, ADMINS, u, &credential)
	http.Redirect(w, r, "/AdminOnly", http.StatusFound)
}

// func (u *userController) LoginResidentByJSON(w http.ResponseWriter, r *http.Request) {
// 	var credential Credential
// 	err := decodeJSONBody(w, r, &credential)
// 	if err != nil {
// 		var mr *malformedRequest
// 		if errors.As(err, &mr) {
// 			http.Error(w, mr.msg, mr.status)
// 		} else {
// 			log.Println(err.Error())
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	LoginHelper(w, r, RESIDENTS, u, &credential)

// 	//Do something
// 	fmt.Fprintf(w, "Welcome %v!!", credential.Username)

// }

// func (u *userController) LoginAdminByJSON(w http.ResponseWriter, r *http.Request) {
// 	var credential Credential
// 	err := decodeJSONBody(w, r, &credential)
// 	if err != nil {
// 		var mr *malformedRequest
// 		if errors.As(err, &mr) {
// 			http.Error(w, mr.msg, mr.status)
// 		} else {
// 			log.Println(err.Error())
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	LoginHelper(w, r, ADMINS, u, &credential)

// 	//Do something
// 	fmt.Fprintf(w, "Welcome %v!!", credential.Username)

// }

//BEFORE LOG IN CONTROLLER
func (u *userController) ResidentLoginForm(w http.ResponseWriter, r *http.Request) {
	if CheckSessionSimple(w, r) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// u.tp.ExecuteTemplate(w, "residentlogin.html", map[string]interface{}{
	// 	csrf.TemplateTag: csrf.TemplateField(r),
	// })

	// token := csrf.Token(r)
	// w.Header().Set("X-CSRF-Token", token)

	u.tp.ExecuteTemplate(w, "residentlogin.html", nil)
}

func (u *userController) AdminLoginForm(w http.ResponseWriter, r *http.Request) {
	if CheckAdminSessionSimple(w, r) {
		http.Redirect(w, r, "/AdminOnly", http.StatusFound)
	}
	// u.tp.ExecuteTemplate(w, "adminlogin.html", map[string]interface{}{
	// 	csrf.TemplateTag: csrf.TemplateField(r),
	// })

	// token := csrf.Token(r)
	// w.Header().Set("X-CSRF-Token", token)

	u.tp.ExecuteTemplate(w, "adminlogin.html", nil)
}

func (u *userController) AreYouLoggedInAsResident(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, RESIDENTS, u, name, ID)

	fmt.Fprintf(w, "Welcome %v with ID %v!! Still in progress!", (*name), (*ID))

	//Main Page
}

func (u *userController) AreYouLoggedInAsAdmin(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	notifNumber, err := u.userService.HowManyNotificationUnsolved()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	output := struct {
		NotifNumber int
		Name        string
	}{NotifNumber: notifNumber,
		Name: *name}

	//fmt.Fprintf(w, "Welcome %v!! Still in progress!", "Random Admin")
	u.tp.ExecuteTemplate(w, "adminmainpage.html", output)
	//Main Admin Page
}

//ADMIN FUNCTION CONTROLLER
func (u *userController) SolveNotificationPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something
	notifications, err := u.userService.ReadAllNotifications()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	type Output struct {
		Notifications domain.Notifications
		DetectedLink  string
	}

	var output []Output

	for _, e := range *notifications {
		output = append(output, Output{e, *u.userService.NotificationLinkDetector(&e.Description)})

	}
	// "<a href=\"/AdminOnly/solve-notification/" + strconv.Itoa(int(e.ID)) + "\" class=\"button\">Click here if it's solved</a>"
	u.tp.ExecuteTemplate(w, "adminsolvenotificationpage.html", output)
}

func (u *userController) SolveNotification(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// temp, err := strconv.ParseUint(vars["id"], 10, 0)
	// if err != nil {
	// 	fmt.Fprintf(w, "Error : %v", err.Error())
	// 	return
	// }
	// ID := uint(temp)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	temp := r.PostForm["NotifID"]
	ID2, err := strconv.Atoi(temp[0])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	ID := uint(ID2)
	err = u.userService.SolveNotificationByID(ID)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	// u.SolveNotificationPage(w, r)
	u.tp.ExecuteTemplate(w, "successpage.html", "Solving Notification")
}

func (u *userController) CheckInPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something
	u.tp.ExecuteTemplate(w, "admincheckinchoosedatepage.html", nil)
}

func (u *userController) CheckInRoomPickPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	// fmt.Println(r.Form)

	endDate := r.Form["EndDate"]

	permanentStay := ""
	if len(r.Form["PermanentStay"]) > 0 {
		permanentStay = "True"
	}

	rooms, err := u.userService.ReadAllUnoccupiedRooms()

	type Output struct {
		Room          domain.Rooms
		EndDate       string
		PermanentStay string
	}

	var output []Output

	for _, e := range *rooms {
		output = append(output, Output{e, endDate[0], permanentStay})

	}

	//Do Something
	u.tp.ExecuteTemplate(w, "admincheckinchooseroompage.html", output)
}

func (u *userController) CheckInRoomGatherPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	// fmt.Println(r.Form)

	RoomID := r.Form["RoomID"]
	endDate := r.Form["EndDate"]
	permanentStay := ""
	if r.Form["PermanentStay"][0] != "" {
		permanentStay = "True"
	}

	output := struct {
		RoomID        string
		EndDate       string
		PermanentStay string
	}{RoomID: RoomID[0], EndDate: endDate[0], PermanentStay: permanentStay}

	//Do Something
	u.tp.ExecuteTemplate(w, "admincheckingatherpage.html", output)
}

func (u *userController) CheckInRoomFillinPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	// fmt.Println(r.Form)

	RoomID := r.PostForm["RoomID"]
	endDate := r.PostForm["EndDate"]
	PhoneNumber := r.PostForm["PhoneNumber"]
	permanentStay := ""
	username := ""
	password := ""
	if r.PostForm["PermanentStay"][0] != "" {
		permanentStay = "True"
		username = r.PostForm["username"][0]
		password = r.PostForm["password"][0]
	}

	output := struct {
		RoomID        string
		EndDate       string
		PhoneNumber   string
		PermanentStay string
		Username      string
		Password      string
	}{RoomID: RoomID[0], EndDate: endDate[0], PermanentStay: permanentStay,
		PhoneNumber: PhoneNumber[0], Username: username, Password: password}

	//Do Something
	u.tp.ExecuteTemplate(w, "admincheckinfillinpage.html", output)
}

func (u *userController) CheckInRoom(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	// fmt.Println(r.Form)

	RoomID := r.PostForm["RoomID"]
	endDate := r.PostForm["EndDate"]
	PhoneNumber := r.PostForm["PhoneNumber"]
	permanentStay := false
	username := ""
	password := ""
	if r.Form["PermanentStay"][0] != "" {
		permanentStay = true
		username = r.PostForm["username"][0]
		password = r.PostForm["password"][0]
	}
	FullName := r.PostForm["FullName"]
	NIKNumber := r.PostForm["NIKNumber"]
	Address := r.PostForm["Address"]
	Occupation := r.PostForm["Occupation"]

	err = u.userService.FormToCheckInRoom(RoomID[0], endDate[0], PhoneNumber[0],
		permanentStay, username, password, FullName[0], NIKNumber[0], Address[0], Occupation[0])

	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	//Do Something
	u.tp.ExecuteTemplate(w, "successpage.html", "Check in Mr./Ms. "+FullName[0]+" on Room No. "+RoomID[0])
}

func (u *userController) CheckOutPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something
	u.tp.ExecuteTemplate(w, "admincheckoutpage.html", nil)
}

func (u *userController) CheckOutConfirmationPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	temp := r.Form["RoomID"]
	ID2, err := strconv.Atoi(temp[0])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	roomID := uint(ID2)

	room, err := u.userService.ReadRoomByID(roomID)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	resident, err := u.userService.ReadResidentByID(room.ResidentID)

	output := struct {
		Room     domain.Rooms
		Resident domain.Residents
	}{Room: *room, Resident: *resident}

	u.tp.ExecuteTemplate(w, "admincheckoutconfirmationpage.html", output)
}

func (u *userController) CheckOutRoom(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	temp := r.Form["RoomID"]
	ID2, err := strconv.Atoi(temp[0])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	roomID := uint(ID2)

	err = u.userService.CheckOutRoomByID(&roomID)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	u.tp.ExecuteTemplate(w, "successpage.html", "Check Out Room No. "+temp[0])
}

func (u *userController) ReadyUpRoomPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something
	rooms, err := u.userService.ReadAllNotReadyRooms()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	u.tp.ExecuteTemplate(w, "adminreadyuproompage.html", rooms)
}

func (u *userController) ReadyUpRoom(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	temp := r.Form["RoomID"]
	ID2, err := strconv.Atoi(temp[0])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	roomID := uint(ID2)

	err = u.userService.ReadyUpRoomByID(&roomID)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	u.tp.ExecuteTemplate(w, "successpage.html", "Ready Up Room No. "+temp[0])
}

func (u *userController) InformationPage(w http.ResponseWriter, r *http.Request) {
	name, ID := CheckAndTakeAdminSession(w, r)
	if name == nil {
		http.Redirect(w, r, "/AdminOnly/login", http.StatusFound)
		return
	}
	CookiesSetter(w, r, ADMINS, u, name, ID)

	//Do Something
	fmt.Fprintf(w, "Still in Progress, tanya yang megang aplikasinya aja ya :V")
}

//HELPER FUNCTION
type DataType int

const (
	RESIDENTS = iota
	ROOMS
	INVOICES
	ADMINS
	NOTIFICATIONS
	PROBLEMS
)

func LoginHelper(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController, credential *Credential) {
	switch dataType {
	case RESIDENTS:
		info, err := u.userService.LoginResident(credential.Username, credential.Password)
		if err != nil {
			fmt.Fprintf(w, "Error : %v", err.Error())
			//Alert or do something when login failed
		} else {
			CookiesSetter(w, r, RESIDENTS, u, &info.FullName, &info.ID)
		}
	case ADMINS:
		info, err := u.userService.LoginAdmin(credential.Username, credential.Password)
		if err != nil {
			fmt.Fprintf(w, "Error : %v", err.Error())
			//Alert or do something when login failed
		} else {
			CookiesSetter(w, r, ADMINS, u, &info.AdminName, &info.ID)
		}
	}
}

func CookiesSetter(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController, secret *string, secretID *uint) {
	switch dataType {
	case RESIDENTS:
		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, _ := store.Get(r, "session")
		// Set some session values.
		session.Values["Name"] = (*secret)
		session.Values["ID"] = (*secretID)
		// Save it before we write to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case ADMINS:
		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, _ := store.Get(r, "admin-session")
		// Set some session values.
		session.Values["AdminName"] = (*secret)
		session.Values["AdminID"] = (*secretID)
		// Save it before we write to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	//fmt.Fprintf(w, "Welcome %v!", secret)
}

func CheckAndTakeSession(w http.ResponseWriter, r *http.Request) (*string, *uint) {
	if !CheckSessionSimple(w, r) {
		return nil, nil
	}
	session, _ := store.Get(r, "session")
	temp := session.Values["ID"]
	ID, ok := temp.(uint)
	if !ok {
		fmt.Fprintf(w, "Error : Something wrong on assertion")
		return nil, nil
	}
	temp = session.Values["Name"]
	name, ok := temp.(string)
	if !ok {
		fmt.Fprintf(w, "Error : Something wrong on assertion")
		return nil, nil
	}
	return &name, &ID
}

func CheckAndTakeAdminSession(w http.ResponseWriter, r *http.Request) (*string, *uint) {
	if !CheckAdminSessionSimple(w, r) {
		return nil, nil
	}
	session, _ := store.Get(r, "admin-session")
	temp := session.Values["AdminID"]
	ID, ok := temp.(uint)
	if !ok {
		fmt.Fprintf(w, "Error : Something wrong on assertion")
		return nil, nil
	}
	temp = session.Values["AdminName"]
	name, ok := temp.(string)
	if !ok {
		fmt.Fprintf(w, "Error : Something wrong on assertion")
		return nil, nil
	}
	return &name, &ID
}

func CheckSessionSimple(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["ID"]
	// fmt.Println("ok:", ok)
	// if !ok {
	// 	fmt.Fprintf(w, "Error : Please login first!")
	// }
	return ok
}

func CheckAdminSessionSimple(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "admin-session")
	_, ok := session.Values["AdminID"]
	// fmt.Println("ok:", ok)
	// if !ok {
	// 	fmt.Fprintf(w, "Error : Please login first!")
	// }
	return ok
}

//CRUD SPECIALIZED HELPER FUNCTION

func CRUDRead(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController) {
	if !CheckAdminSessionSimple(w, r) {
		fmt.Fprintf(w, "Error : Please login first!")
		return
	}

	vars := mux.Vars(r)
	temp, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	ID := uint(temp)

	var info any
	switch dataType {
	case RESIDENTS:
		info, err = u.userService.ReadResidentByID(ID)
	case ROOMS:
		info, err = u.userService.ReadRoomByID(ID)
	case INVOICES:
		info, err = u.userService.ReadInvoiceByID(ID)
	case ADMINS:
		info, err = u.userService.ReadAdminByID(ID)
	case NOTIFICATIONS:
		info, err = u.userService.ReadNotificationByID(ID)
	case PROBLEMS:
		info, err = u.userService.ReadProblemByID(ID)
	}

	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	data, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
	} else {
		fmt.Fprintf(w, "%v", string(data))
	}
}

func CRUDReadAll(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController) {
	if !CheckAdminSessionSimple(w, r) {
		fmt.Fprintf(w, "Error : Please login first!")
		return
	}

	var info any
	var err error
	switch dataType {
	case RESIDENTS:
		info, err = u.userService.ReadAllResidents()
	case ROOMS:
		info, err = u.userService.ReadAllRooms()
	case INVOICES:
		info, err = u.userService.ReadAllInvoices()
	case ADMINS:
		info, err = u.userService.ReadAllAdmins()
	case NOTIFICATIONS:
		info, err = u.userService.ReadAllNotifications()
	case PROBLEMS:
		info, err = u.userService.ReadAllProblems()
	}

	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}
	data, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
	} else {
		fmt.Fprintf(w, "%v", string(data))
	}
}

func CRUDCreate(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController) {
	if !CheckAdminSessionSimple(w, r) {
		fmt.Fprintf(w, "Error : Please login first!")
		return
	}

	var info any
	var err error
	switch dataType {
	case RESIDENTS:
		temp := domain.Residents{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case ROOMS:
		temp := domain.Rooms{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case INVOICES:
		temp := domain.Invoices{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case ADMINS:
		temp := domain.Admins{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case NOTIFICATIONS:
		temp := domain.Notifications{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case PROBLEMS:
		temp := domain.Problems{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	}
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	var ID uint
	switch dataType {
	case RESIDENTS:
		info, ok := info.(domain.Residents)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Resident: %+v\n", info)
		result, errtemp := u.userService.CreateResident(&(info))
		err = errtemp
		ID = result.ID
	case ROOMS:
		info, ok := info.(domain.Rooms)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Room: %+v\n", info)
		result, errtemp := u.userService.CreateRoom(&(info))
		err = errtemp
		ID = result.ID
	case INVOICES:
		info, ok := info.(domain.Invoices)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Invoice: %+v\n", info)
		result, errtemp := u.userService.CreateInvoice(&(info))
		err = errtemp
		ID = result.ID
	case ADMINS:
		info, ok := info.(domain.Admins)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Admin: %+v\n", info)
		result, errtemp := u.userService.CreateAdmin(&(info))
		err = errtemp
		ID = result.ID
	case NOTIFICATIONS:
		info, ok := info.(domain.Notifications)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Notification: %+v\n", info)
		result, errtemp := u.userService.CreateNotification(&(info))
		err = errtemp
		ID = result.ID
	case PROBLEMS:
		info, ok := info.(domain.Problems)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		fmt.Fprintf(w, "Problem: %+v\n", info)
		result, errtemp := u.userService.CreateProblem(&(info))
		err = errtemp
		ID = result.ID
	}
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
	} else {
		fmt.Fprintf(w, "%T with ID %v created successfully", info, ID)
	}
}

func CRUDUpdate(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController) {
	if !CheckAdminSessionSimple(w, r) {
		fmt.Fprintf(w, "Error : Please login first!")
		return
	}

	var info any
	var err error
	switch dataType {
	case RESIDENTS:
		temp := domain.Residents{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case ROOMS:
		temp := domain.Rooms{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case INVOICES:
		temp := domain.Invoices{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case ADMINS:
		temp := domain.Admins{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case NOTIFICATIONS:
		temp := domain.Notifications{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	case PROBLEMS:
		temp := domain.Problems{}
		errtemp := decodeJSONBody(w, r, &temp)
		info = temp
		err = errtemp
	}
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	switch dataType {
	case RESIDENTS:
		info, ok := info.(domain.Residents)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Resident: %+v\n", info)
		_, errtemp := u.userService.UpdateResident(&info)
		err = errtemp
	case ROOMS:
		info, ok := info.(domain.Rooms)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Room: %+v\n", info)
		_, errtemp := u.userService.UpdateRoom(&info)
		err = errtemp
	case INVOICES:
		info, ok := info.(domain.Invoices)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Invoice: %+v\n", info)
		_, errtemp := u.userService.UpdateInvoice(&info)
		err = errtemp
	case ADMINS:
		info, ok := info.(domain.Admins)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Admin: %+v\n", info)
		_, errtemp := u.userService.UpdateAdmin(&info)
		err = errtemp
	case NOTIFICATIONS:
		info, ok := info.(domain.Notifications)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Notification: %+v\n", info)
		_, errtemp := u.userService.UpdateNotification(&info)
		err = errtemp
	case PROBLEMS:
		info, ok := info.(domain.Problems)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Problem: %+v\n", info)
		_, errtemp := u.userService.UpdateProblem(&info)
		err = errtemp
	}
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
	} else {
		fmt.Fprintf(w, "%T with ID %v updated successfully", info, ID)
	}
}

func CRUDDelete(w http.ResponseWriter, r *http.Request, dataType DataType, u *userController) {
	if !CheckAdminSessionSimple(w, r) {
		fmt.Fprintf(w, "Error : Please login first!")
		return
	}

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
		return
	}

	var info any
	switch dataType {
	case RESIDENTS:
		info = domain.Residents{}
		info, ok := info.(domain.Residents)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Resident: %+v\n", info)
		errtemp := u.userService.DeleteResident(&info)
		err = errtemp
	case ROOMS:
		info = domain.Rooms{}
		info, ok := info.(domain.Rooms)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Room: %+v\n", info)
		errtemp := u.userService.DeleteRoom(&info)
		err = errtemp
	case INVOICES:
		info = domain.Invoices{}
		info, ok := info.(domain.Invoices)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Invoice: %+v\n", info)
		errtemp := u.userService.DeleteInvoice(&info)
		err = errtemp
	case ADMINS:
		info = domain.Admins{}
		info, ok := info.(domain.Admins)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Admin: %+v\n", info)
		errtemp := u.userService.DeleteAdmin(&info)
		err = errtemp
	case NOTIFICATIONS:
		info = domain.Notifications{}
		info, ok := info.(domain.Notifications)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Notification: %+v\n", info)
		errtemp := u.userService.DeleteNotification(&info)
		err = errtemp
	case PROBLEMS:
		info = domain.Problems{}
		info, ok := info.(domain.Problems)
		if !ok {
			fmt.Fprintf(w, "Error : Something wrong on assertion")
			return
		}
		info.ID = uint(ID)
		fmt.Fprintf(w, "Problem: %+v\n", info)
		errtemp := u.userService.DeleteProblem(&info)
		err = errtemp
	}
	if err != nil {
		fmt.Fprintf(w, "Error : %v", err.Error())
	} else {
		fmt.Fprintf(w, "%T with ID %v deleted successfully", info, ID)
	}
}
