package domain

import (
	"time"

	"github.com/gorilla/schema"
)

var SchemaDecoder = schema.NewDecoder()

//Database Models

/* // UserService : represent the user's services
type UserService interface {
	Validate(user *User) error
	ValidateAge(user *User) bool
	Create(user *User) (*User, error)
	FindAll() ([]User, error)
} */

/* // UserRepository : represent the user's repository contract
type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
} */

type UserRepository interface {
	CreateAdmin(admin *Admins) (*Admins, error)
	ReadAllAdmins() (*[]Admins, error)
	ReadAdminByID(ID uint) (*Admins, error)
	ReadAdminByUsername(username string) (*Admins, error)
	UpdateAdmin(admin *Admins) (*Admins, error)
	DeleteAdmin(admin *Admins) error

	CreateResident(resident *Residents) (*Residents, error)
	ReadAllResidents() (*[]Residents, error)
	ReadResidentByID(ID uint) (*Residents, error)
	ReadResidentByUsername(username string) (*Residents, error)
	UpdateResident(resident *Residents) (*Residents, error)
	DeleteResident(resident *Residents) error

	CreateRoom(room *Rooms) (*Rooms, error)
	ReadAllRooms() (*[]Rooms, error)
	UpdateRoom(room *Rooms) (*Rooms, error)
	DeleteRoom(room *Rooms) error
	ReadRoomByID(ID uint) (*Rooms, error)

	CreateInvoice(invoice *Invoices) (*Invoices, error)
	ReadAllInvoices() (*[]Invoices, error)
	UpdateInvoice(invoice *Invoices) (*Invoices, error)
	DeleteInvoice(invoice *Invoices) error
	ReadInvoiceByID(ID uint) (*Invoices, error)

	CreateProblem(problem *Problems) (*Problems, error)
	ReadAllProblems() (*[]Problems, error)
	UpdateProblem(problem *Problems) (*Problems, error)
	DeleteProblem(problem *Problems) error
	ReadProblemByID(ID uint) (*Problems, error)

	CreateNotification(notification *Notifications) (*Notifications, error)
	ReadAllNotifications() (*[]Notifications, error)
	UpdateNotification(notification *Notifications) (*Notifications, error)
	DeleteNotification(notification *Notifications) error
	ReadNotificationByID(ID uint) (*Notifications, error)

	ReadLastResident() (*Residents, error)
	ReadLastNotification() (*Notifications, error)
	ReadAllNotReadyRooms() (*[]Rooms, error)
	ReadAllUnoccupiedRooms() (*[]Rooms, error)
	/* Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error */
}

type UserService interface {
	CreateAdmin(admin *Admins) (*Admins, error)
	ReadAllAdmins() (*[]Admins, error)
	ReadAdminByID(ID uint) (*Admins, error)
	UpdateAdmin(admin *Admins) (*Admins, error)
	DeleteAdmin(admin *Admins) error

	CreateResident(resident *Residents) (*Residents, error)
	ReadAllResidents() (*[]Residents, error)
	UpdateResident(resident *Residents) (*Residents, error)
	DeleteResident(resident *Residents) error
	ReadResidentByID(ID uint) (*Residents, error)

	CreateRoom(room *Rooms) (*Rooms, error)
	ReadAllRooms() (*[]Rooms, error)
	UpdateRoom(room *Rooms) (*Rooms, error)
	DeleteRoom(room *Rooms) error
	ReadRoomByID(ID uint) (*Rooms, error)

	CreateInvoice(invoice *Invoices) (*Invoices, error)
	ReadAllInvoices() (*[]Invoices, error)
	UpdateInvoice(invoice *Invoices) (*Invoices, error)
	DeleteInvoice(invoice *Invoices) error
	ReadInvoiceByID(ID uint) (*Invoices, error)
	CreateProblem(problem *Problems) (*Problems, error)

	ReadAllProblems() (*[]Problems, error)
	UpdateProblem(problem *Problems) (*Problems, error)
	DeleteProblem(problem *Problems) error
	ReadProblemByID(ID uint) (*Problems, error)

	CreateNotification(notification *Notifications) (*Notifications, error)
	ReadAllNotifications() (*[]Notifications, error)
	UpdateNotification(notification *Notifications) (*Notifications, error)
	DeleteNotification(notification *Notifications) error
	ReadNotificationByID(ID uint) (*Notifications, error)

	LoginAdmin(username string, password string) (*Admins, error)
	LoginResident(username string, password string) (*Residents, error)

	StartDayChangeRoutine()

	FormToCheckInRoom(RoomID string, endDate string, PhoneNumber string,
		permanentStay bool, username string, password string, FullName string, NIKNumber string,
		Address string, Occupation string) error
	CheckInRoom(resident *Residents, roomID *uint, start_date time.Time, end_date time.Time) error
	ReadAllUnoccupiedRooms() (*[]Rooms, error)
	CheckOutRoomByID(id *uint) error
	ReadAllNotReadyRooms() (*[]Rooms, error)
	ReadyUpRoomByID(id *uint) error

	CreateNewResident(stayType StayType, fullName string,
		NIKNumber uint, Address string, Occupation string, PhoneNumber uint,
		username string, password string) (*Residents, error)
	CreateNewNotification(title *string, description *string) (*Notifications, error)
	HowManyNotificationUnsolved() (int, error)
	NotificationLinkDetector(s *string) *string
	SolveNotificationByID(ID uint) error
	ProblemExtractorFromNotification(notif *Notifications) (*Problems, error)
	/* Validate(user *User) error
	ValidateAge(user *User) bool
	Create(user *User) (*User, error)
	FindAll() ([]User, error) */
}
