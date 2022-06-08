package repository

import (
	"github.com/gpsandyka/manajemen_kos/domain"
	//"time"

	//"github.com/jinzhu/gorm"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO dibikin function buat nyalain connection ke db

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository : get injected database
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{DB: db}
}

/*  func (u *userRepository) Save(user *domain.User) (*domain.User, error) {
	return user, u.DB.Create(user).Error
}
func (u *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := u.DB.Find(&users).Error
	return users, err
}
func (u *userRepository) Delete(user *domain.User) error {
	return u.DB.Delete(&user).Error
}
func (u *userRepository) Migrate() error {
	return u.DB.AutoMigrate(&domain.User{}).Error
}
*/

// ADMIN CRUD FUNCTIONS
func (db *userRepository) CreateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	result := db.DB.Create(&admin)
	return admin, result.Error
}

func (db *userRepository) ReadAllAdmins() (*[]domain.Admins, error) {
	var admins []domain.Admins
	result := db.DB.Find(&admins)
	return &admins, result.Error
}

func (db *userRepository) ReadAdminByID(ID uint) (*domain.Admins, error) {
	var admin = domain.Admins{ID: uint(ID)}
	result := db.DB.First(&admin)
	return &admin, result.Error
}

func (db *userRepository) ReadAdminByUsername(username string) (*domain.Admins, error) {
	var admin = domain.Admins{}
	result := db.DB.Where("Username = ?", (username)).First(&admin)
	return &admin, result.Error
}

func (db *userRepository) UpdateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	result := db.DB.Model(&admin).Updates(admin)
	return admin, result.Error
}

func (db *userRepository) DeleteAdmin(admin *domain.Admins) error {
	result := db.DB.Delete(&admin)
	return result.Error
}

// RESIDENT CRUD FUNCTIONS
func (db *userRepository) CreateResident(resident *domain.Residents) (*domain.Residents, error) {
	result := db.DB.Create(&resident)
	return resident, result.Error
}

func (db *userRepository) ReadAllResidents() (*[]domain.Residents, error) {
	var residents []domain.Residents
	result := db.DB.Find(&residents)
	return &residents, result.Error
}

func (db *userRepository) ReadResidentByID(ID uint) (*domain.Residents, error) {
	var resident = domain.Residents{ID: uint(ID)}
	result := db.DB.First(&resident)
	return &resident, result.Error
}

func (db *userRepository) ReadResidentByUsername(username string) (*domain.Residents, error) {
	var resident = domain.Residents{}
	result := db.DB.Where("Username = ?", (username)).First(&resident)
	return &resident, result.Error
}

func (db *userRepository) UpdateResident(resident *domain.Residents) (*domain.Residents, error) {
	result := db.DB.Model(&resident).Updates(resident)
	return resident, result.Error
}

func (db *userRepository) DeleteResident(resident *domain.Residents) error {
	result := db.DB.Delete(&resident)
	return result.Error
}

// ROOMS CRUD FUNCTIONS
func (db *userRepository) CreateRoom(room *domain.Rooms) (*domain.Rooms, error) {
	result := db.DB.Create(&room)
	return room, result.Error
}

func (db *userRepository) ReadAllRooms() (*[]domain.Rooms, error) {
	var rooms []domain.Rooms
	result := db.DB.Find(&rooms)
	return &rooms, result.Error
}

func (db *userRepository) ReadRoomByID(ID uint) (*domain.Rooms, error) {
	var room = domain.Rooms{ID: uint(ID)}
	result := db.DB.First(&room)
	return &room, result.Error
}

func (db *userRepository) UpdateRoom(room *domain.Rooms) (*domain.Rooms, error) {
	result := db.DB.Model(&room).Updates(room)
	return room, result.Error
}

func (db *userRepository) DeleteRoom(room *domain.Rooms) error {
	result := db.DB.Delete(&room)
	return result.Error
}

// INVOICES CRUD FUNCTIONS
func (db *userRepository) CreateInvoice(invoice *domain.Invoices) (*domain.Invoices, error) {
	result := db.DB.Create(&invoice)
	return invoice, result.Error
}

func (db *userRepository) ReadAllInvoices() (*[]domain.Invoices, error) {
	var invoices []domain.Invoices
	result := db.DB.Find(&invoices)
	return &invoices, result.Error
}

func (db *userRepository) ReadInvoiceByID(ID uint) (*domain.Invoices, error) {
	var invoice = domain.Invoices{ID: uint(ID)}
	result := db.DB.First(&invoice)
	return &invoice, result.Error
}

func (db *userRepository) UpdateInvoice(invoice *domain.Invoices) (*domain.Invoices, error) {
	result := db.DB.Model(&invoice).Updates(invoice)
	return invoice, result.Error
}

func (db *userRepository) DeleteInvoice(invoice *domain.Invoices) error {
	result := db.DB.Delete(&invoice)
	return result.Error
}

//PROBLEMS CRUD FUNCTIONS
func (db *userRepository) CreateProblem(problem *domain.Problems) (*domain.Problems, error) {
	result := db.DB.Create(&problem)
	return problem, result.Error
}

func (db *userRepository) ReadAllProblems() (*[]domain.Problems, error) {
	var problems []domain.Problems
	result := db.DB.Find(&problems)
	return &problems, result.Error
}

func (db *userRepository) ReadProblemByID(ID uint) (*domain.Problems, error) {
	var problem = domain.Problems{ID: uint(ID)}
	result := db.DB.First(&problem)
	return &problem, result.Error
}

func (db *userRepository) UpdateProblem(problem *domain.Problems) (*domain.Problems, error) {
	result := db.DB.Model(&problem).Updates(problem)
	return problem, result.Error
}

func (db *userRepository) DeleteProblem(problem *domain.Problems) error {
	result := db.DB.Delete(&problem)
	return result.Error
}

//NOTIFICATIONS CRUD FUNCTIONS
func (db *userRepository) CreateNotification(notification *domain.Notifications) (*domain.Notifications, error) {
	result := db.DB.Create(&notification)
	return notification, result.Error
}

func (db *userRepository) ReadAllNotifications() (*[]domain.Notifications, error) {
	var notifications []domain.Notifications
	result := db.DB.Find(&notifications)
	return &notifications, result.Error
}

func (db *userRepository) ReadNotificationByID(ID uint) (*domain.Notifications, error) {
	var notification = domain.Notifications{ID: uint(ID)}
	result := db.DB.First(&notification)
	return &notification, result.Error
}

func (db *userRepository) UpdateNotification(notification *domain.Notifications) (*domain.Notifications, error) {
	result := db.DB.Model(&notification).Updates(notification)
	return notification, result.Error
}

func (db *userRepository) DeleteNotification(notification *domain.Notifications) error {
	result := db.DB.Delete(&notification)
	return result.Error
}

//MISC FUNCTION
func (db *userRepository) ReadLastResident() (*domain.Residents, error) {
	var resident domain.Residents
	result := db.DB.Last(&resident)
	return &resident, result.Error
}

func (db *userRepository) ReadLastNotification() (*domain.Notifications, error) {
	var notification domain.Notifications
	result := db.DB.Last(&notification)
	return &notification, result.Error
}

func (db *userRepository) ReadAllNotReadyRooms() (*[]domain.Rooms, error) {
	var rooms []domain.Rooms
	result := db.DB.Where("status = ?", domain.NOTREADY).Find(&rooms)
	return &rooms, result.Error
}

func (db *userRepository) ReadAllUnoccupiedRooms() (*[]domain.Rooms, error) {
	var rooms []domain.Rooms
	result := db.DB.Where("status = ?", domain.UNOCCUPIED).Find(&rooms)
	return &rooms, result.Error
}
