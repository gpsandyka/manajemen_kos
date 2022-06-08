package repository

import (
	"github.com/gpsandyka/manajemen_kos/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}
type ProblemMock struct {
	Mock mock.Mock
}

func (db *UserRepositoryMock) CreateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	arguments := db.Mock.Called(admin)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Admins), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllAdmins() (*[]domain.Admins, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Admins), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAdminByID(ID uint) (*domain.Admins, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Admins), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAdminByUsername(username string) (*domain.Admins, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Admins), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	arguments := db.Mock.Called(admin)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Admins), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteAdmin(admin *domain.Admins) error {
	arguments := db.Mock.Called(admin)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// // RESIDENT CRUD FUNCTIONS
func (db *UserRepositoryMock) CreateResident(Resident *domain.Residents) (*domain.Residents, error) {
	arguments := db.Mock.Called(Resident)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllResidents() (*[]domain.Residents, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadResidentByID(ID uint) (*domain.Residents, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadResidentByUsername(username string) (*domain.Residents, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateResident(Resident *domain.Residents) (*domain.Residents, error) {
	arguments := db.Mock.Called(Resident)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteResident(Resident *domain.Residents) error {
	arguments := db.Mock.Called(Resident)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// // ROOMS CRUD FUNCTIONS

func (db *UserRepositoryMock) CreateRoom(Room *domain.Rooms) (*domain.Rooms, error) {
	arguments := db.Mock.Called(Room)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllRooms() (*[]domain.Rooms, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadRoomByID(ID uint) (*domain.Rooms, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadRoomByUsername(username string) (*domain.Rooms, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateRoom(Room *domain.Rooms) (*domain.Rooms, error) {
	arguments := db.Mock.Called(Room)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteRoom(Room *domain.Rooms) error {
	arguments := db.Mock.Called(Room)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// // INVOICES CRUD FUNCTIONS
func (db *UserRepositoryMock) CreateInvoice(Invoice *domain.Invoices) (*domain.Invoices, error) {
	arguments := db.Mock.Called(Invoice)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Invoices), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllInvoices() (*[]domain.Invoices, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Invoices), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadInvoiceByID(ID uint) (*domain.Invoices, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Invoices), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadInvoiceByUsername(username string) (*domain.Invoices, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Invoices), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateInvoice(Invoice *domain.Invoices) (*domain.Invoices, error) {
	arguments := db.Mock.Called(Invoice)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Invoices), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteInvoice(Invoice *domain.Invoices) error {
	arguments := db.Mock.Called(Invoice)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// //PROBLEMS CRUD FUNCTIONS
func (db *UserRepositoryMock) CreateProblem(Problem *domain.Problems) (*domain.Problems, error) {
	arguments := db.Mock.Called(Problem)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Problems), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllProblems() (*[]domain.Problems, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Problems), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadProblemByID(ID uint) (*domain.Problems, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Problems), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadProblemByUsername(username string) (*domain.Problems, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Problems), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateProblem(Problem *domain.Problems) (*domain.Problems, error) {
	arguments := db.Mock.Called(Problem)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Problems), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteProblem(Problem *domain.Problems) error {
	arguments := db.Mock.Called(Problem)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// //NOTIFICATIONS CRUD FUNCTIONS
func (db *UserRepositoryMock) CreateNotification(Notification *domain.Notifications) (*domain.Notifications, error) {
	arguments := db.Mock.Called(Notification)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllNotifications() (*[]domain.Notifications, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadNotificationByID(ID uint) (*domain.Notifications, error) {
	arguments := db.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadNotificationByUsername(username string) (*domain.Notifications, error) {
	arguments := db.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) UpdateNotification(Notification *domain.Notifications) (*domain.Notifications, error) {
	arguments := db.Mock.Called(Notification)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) DeleteNotification(Notification *domain.Notifications) error {
	arguments := db.Mock.Called(Notification)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Error(0)
}

// //MISC FUNCTION
func (db *UserRepositoryMock) ReadLastResident() (*domain.Residents, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Residents), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadLastNotification() (*domain.Notifications, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*domain.Notifications), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllNotReadyRooms() (*[]domain.Rooms, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Rooms), arguments.Error(1)
}

func (db *UserRepositoryMock) ReadAllUnoccupiedRooms() (*[]domain.Rooms, error) {
	arguments := db.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(*[]domain.Rooms), arguments.Error(1)
}
