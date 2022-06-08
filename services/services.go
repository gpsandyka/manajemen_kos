package services

import (
	//"errors"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gpsandyka/manajemen_kos/domain"
	"golang.org/x/crypto/bcrypt"

	//"gorm.io/gorm"

	//"log"
	"sync"
	"time"
)

var once sync.Once

type userService struct {
	userRepository domain.UserRepository
}

var instance *userService

//NewUserService: construction function, injected by user repository
func NewUserService(r domain.UserRepository) domain.UserService {
	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})
	return instance
}

/* func (*userService) Validate(user *domain.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The name of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("The email of user is empty")
		return err
	}
	if user.DOB == "" {
		err := errors.New("The DOB of user is empty")
		return err
	}
	return nil
}
func (*userService) ValidateAge(user *domain.User) bool {
	ageLimit := 13
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return false
	}
	diff := now.Sub(dob)
	diffInYears := int(diff.Hours() / (24 * 7 * 4 * 12))
	if diffInYears < ageLimit {
		return false
	} else {
		return true
	}
}
func (u *userService) Create(user *domain.User) (*domain.User, error) {
	return u.userRepository.Save(user)
}
func (u *userService) FindAll() ([]domain.User, error) {
	return u.userRepository.FindAll()
} */

//RESIDENT CRUD SERVICE
func (u *userService) CreateResident(resident *domain.Residents) (*domain.Residents, error) {
	if (*resident).Password != "" {
		text, err := bcrypt.GenerateFromPassword([]byte((*resident).Password), 0)
		if err != nil {
			return resident, err
		}
		(*resident).Password = string(text)
	}
	return u.userRepository.CreateResident(resident)
}

func (u *userService) ReadAllResidents() (*[]domain.Residents, error) {
	return u.userRepository.ReadAllResidents()
}

func (u *userService) ReadResidentByID(ID uint) (*domain.Residents, error) {
	return u.userRepository.ReadResidentByID(ID)
}

func (u *userService) UpdateResident(resident *domain.Residents) (*domain.Residents, error) {
	if (*resident).Password != "" {
		text, err := bcrypt.GenerateFromPassword([]byte((*resident).Password), 0)
		if err != nil {
			return resident, err
		}
		(*resident).Password = string(text)
	}
	return u.userRepository.UpdateResident(resident)
}

func (u *userService) DeleteResident(resident *domain.Residents) error {
	return u.userRepository.DeleteResident(resident)
}

//ROOMS CRUD SERVICE
func (u *userService) CreateRoom(room *domain.Rooms) (*domain.Rooms, error) {
	return u.userRepository.CreateRoom(room)
}

func (u *userService) ReadAllRooms() (*[]domain.Rooms, error) {
	return u.userRepository.ReadAllRooms()
}

func (u *userService) ReadRoomByID(ID uint) (*domain.Rooms, error) {
	return u.userRepository.ReadRoomByID(ID)
}

func (u *userService) UpdateRoom(room *domain.Rooms) (*domain.Rooms, error) {
	return u.userRepository.UpdateRoom(room)
}

func (u *userService) DeleteRoom(room *domain.Rooms) error {
	return u.userRepository.DeleteRoom(room)
}

//INVOICES CRUD SERVICE
func (u *userService) CreateInvoice(invoice *domain.Invoices) (*domain.Invoices, error) {
	return u.userRepository.CreateInvoice(invoice)
}

func (u *userService) ReadAllInvoices() (*[]domain.Invoices, error) {
	return u.userRepository.ReadAllInvoices()
}

func (u *userService) ReadInvoiceByID(ID uint) (*domain.Invoices, error) {
	return u.userRepository.ReadInvoiceByID(ID)
}

func (u *userService) UpdateInvoice(invoice *domain.Invoices) (*domain.Invoices, error) {
	return u.userRepository.UpdateInvoice(invoice)
}

func (u *userService) DeleteInvoice(invoice *domain.Invoices) error {
	return u.userRepository.DeleteInvoice(invoice)
}

//ADMINS CRUD SERVICE
func (u *userService) CreateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	if (*admin).Password != "" {
		text, err := bcrypt.GenerateFromPassword([]byte((*admin).Password), 0)
		if err != nil {
			return admin, err
		}
		(*admin).Password = string(text)
	}
	return u.userRepository.CreateAdmin(admin)
}

func (u *userService) ReadAllAdmins() (*[]domain.Admins, error) {
	return u.userRepository.ReadAllAdmins()
}

func (u *userService) ReadAdminByID(ID uint) (*domain.Admins, error) {
	return u.userRepository.ReadAdminByID(ID)
}

func (u *userService) UpdateAdmin(admin *domain.Admins) (*domain.Admins, error) {
	if (*admin).Password != "" {
		text, err := bcrypt.GenerateFromPassword([]byte((*admin).Password), 0)
		if err != nil {
			return admin, err
		}
		(*admin).Password = string(text)
	}
	return u.userRepository.UpdateAdmin(admin)
}

func (u *userService) DeleteAdmin(admin *domain.Admins) error {
	return u.userRepository.DeleteAdmin(admin)
}

//PROBLEMS CRUD SERVICE
func (u *userService) CreateProblem(problem *domain.Problems) (*domain.Problems, error) {
	return u.userRepository.CreateProblem(problem)
}

func (u *userService) ReadAllProblems() (*[]domain.Problems, error) {
	return u.userRepository.ReadAllProblems()
}

func (u *userService) ReadProblemByID(ID uint) (*domain.Problems, error) {
	return u.userRepository.ReadProblemByID(ID)
}

func (u *userService) UpdateProblem(problem *domain.Problems) (*domain.Problems, error) {
	return u.userRepository.UpdateProblem(problem)
}

func (u *userService) DeleteProblem(problem *domain.Problems) error {
	return u.userRepository.DeleteProblem(problem)
}

//NOTIFICATIONS CRUD SERVICE
func (u *userService) CreateNotification(notification *domain.Notifications) (*domain.Notifications, error) {
	return u.userRepository.CreateNotification(notification)
}

func (u *userService) ReadAllNotifications() (*[]domain.Notifications, error) {
	return u.userRepository.ReadAllNotifications()
}

func (u *userService) ReadNotificationByID(ID uint) (*domain.Notifications, error) {
	return u.userRepository.ReadNotificationByID(ID)
}

func (u *userService) UpdateNotification(notification *domain.Notifications) (*domain.Notifications, error) {
	return u.userRepository.UpdateNotification(notification)
}

func (u *userService) DeleteNotification(notification *domain.Notifications) error {
	return u.userRepository.DeleteNotification(notification)
}

//LOG IN SERVICE

func (u *userService) LoginAdmin(username string, password string) (*domain.Admins, error) {
	admin, err := u.userRepository.ReadAdminByUsername(username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte((password))) != nil {
		return nil, errors.New("username and password combination invalid")
	}
	return admin, nil
}

func (u *userService) LoginResident(username string, password string) (*domain.Residents, error) {
	resident, err := u.userRepository.ReadResidentByUsername(username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(resident.Password), []byte((password))) != nil {
		return nil, errors.New("username and password combination invalid")
	}
	return resident, nil
}

//GENERAL SERVICE

func (u *userService) StartDayChangeRoutine() {
	err := DayChangeRoutine(u)
	if err != nil {
		fmt.Println(err)
		//Bikin notification error, tapi masih jalan routinenya
	}
	go func() {
		var temp time.Time
		//for true
		//Delay sampe tengah malam besok (mungkin besok dikurang time.now)
		//DayChangeRoutine()
		for {
			temp = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 1, 0, 0, time.Local)
			time.Sleep(time.Until(temp))
			err = DayChangeRoutine(u)
			if err != nil {
				fmt.Println(err)
				//Bikin notification error, tapi masih jalan routinenya
			}
		}
	}()
}

func DayChangeRoutine(u *userService) error {
	//Periksa semua kamar, dan update tagihan
	err := BillsUpdate(u)
	if err != nil {
		fmt.Println(err)
	}
	return err

}

func BillsUpdate(u *userService) error {
	rooms, err := u.ReadAllRooms()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, e := range *rooms {
		//if end_date null && time.Now().Day == startDate.Day && no notif
		if e.EndOccupiedDate.Before(time.Date(100, 1, 1, 0, 0, 0, 0, time.UTC)) &&
			time.Now().Day() == e.StartOccupiedDate.Day() {

			//In case kejalanin dobel, jadi di check semua notif yang nyala dulu
			title := NotificationTitleGenerator(BILLS, e.ID, 0)
			notifs, err := u.ReadAllNotifications()
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				skip := false
				for _, e := range *notifs {
					if e.Title == title {
						skip = true
					}
				}
				if skip {
					continue
				}
			}

			resident, errtemp := u.ReadResidentByID(e.ResidentID)
			if errtemp != nil {
				err = errtemp
				fmt.Println(err)
				continue
			}
			resident.Bills += e.Price
			_, errtemp = u.UpdateResident(resident)
			if errtemp != nil {
				err = errtemp
				fmt.Println(err)
				continue
			}
			//Add Notif baru
			link := WhatsappAPIGenerate(resident.PhoneNumber,
				"Yth. "+resident.FullName+", tagihan bulanan sebesar Rp. "+strconv.Itoa(int(resident.Bills))+
					" telah tersedia. Dimohon untuk membayar segera. Terimakasih.")
			text := "Tolong ingatkan penghuni tersebut dengan cara membuka whatsapp di PC dan mengklik link : \"" + link + "\""

			_, errtemp = u.CreateNewNotification(&title, &text)
			if errtemp != nil {
				err = errtemp
				fmt.Println(err)
				continue
			}
			//Dan bikin semacam trigger (mungkin?) ke PC kalau ada notif baru
		}
	}

	return err
}

// func (u *userService) ReportProblem() {
// 	//1. Dapet problemnya apa, kalau belum dibikin, bikin baru CreateBewProblem

// 	//2. Bikin notifikasi baru yang berkaitan dengan problemnya

// 	//3. Kasih trigger notif baru(?)
// }

// func (u *userService) PayBills {
// 		//1. GetResidentByID(GetRoomByID(xxx).residentID)

// 		//2. Tampilin informasinya, termasuk tagihannya berapa

// 		//3. Setelah pencet bayar, bikin invoice baru, terus lanjut sampai selesai bayar

// 		//4, Bills nya berkurang jadi 0, terus update

// }

// func (u *userService) ExtendRoom {
// 		//1. Tampilin kalendar, pilih mau di extend sampai tanggal berapa
//		//Note : Yang X tandanya sudah di booking orang lain

// 		//2. Bayar tagihannya

// 		//3. Room.endOccupiedDate = end_date, terus save
// }

//ADMIN FUNCTION SERVICE

func (u *userService) SolveNotificationByID(ID uint) error {
	notif, err := u.userRepository.ReadNotificationByID(ID)
	if err != nil {
		return err
	}
	problem, err := u.ProblemExtractorFromNotification(notif)
	if err != nil && err.Error() != "Problem object not found" {
		return err
	}
	if problem != nil {
		err = u.userRepository.DeleteProblem(problem)
		if err != nil {
			return err
		}
	}
	err = u.userRepository.DeleteNotification(notif)
	return err
}

func (u *userService) FormToCheckInRoom(RoomID string, endDate string, PhoneNumber string,
	permanentStay bool, username string, password string, FullName string, NIKNumber string,
	Address string, Occupation string) error {
	intRoomID, err := strconv.Atoi(RoomID)
	if err != nil {
		return err
	}
	uintRoomID := uint(intRoomID)

	dateEndDate, err := DateParser(endDate)
	if err != nil {
		return err
	}

	intPhoneNumber, err := strconv.Atoi(PhoneNumber)
	if err != nil {
		return err
	}
	uintPhoneNumber := uint(intPhoneNumber)

	stayType := domain.TEMPORARY
	if permanentStay {
		stayType = domain.PERMANENT
	}

	intNIKNumber, err := strconv.Atoi(NIKNumber)
	if err != nil {
		return err
	}
	uintNIKNumber := uint(intNIKNumber)

	resident, err := u.CreateNewResident(stayType, FullName, uintNIKNumber, Address, Occupation, uintPhoneNumber, username, password)
	if err != nil {
		return err
	}

	err = u.CheckInRoom(resident, &uintRoomID, time.Now(), dateEndDate)
	return err
}

func (u *userService) CheckInRoom(resident *domain.Residents, roomID *uint, start_date time.Time, end_date time.Time) error {
	room, err := u.userRepository.ReadRoomByID(*roomID)
	if err != nil {
		return err
	}

	(*room).StartOccupiedDate = start_date
	if resident.StayType == domain.PERMANENT {
		(*room).EndOccupiedDate = time.Date(50, 1, 1, 0, 0, 0, 0, time.UTC)
	} else {
		(*room).EndOccupiedDate = end_date
	}
	(*room).ResidentID = resident.ID
	(*room).Status = domain.OCCUPIED
	_, err = u.userRepository.UpdateRoom(room)
	return err
}

func (u *userService) CheckOutRoomByID(id *uint) error {
	room, err := u.userRepository.ReadRoomByID(*id)
	if err != nil {
		return err
	}
	room.Status = domain.NOTREADY
	room.EndOccupiedDate = time.Now()
	_, err = u.userRepository.UpdateRoom(room)
	return err
}

func (u *userService) ReadAllNotReadyRooms() (*[]domain.Rooms, error) {
	return u.userRepository.ReadAllNotReadyRooms()
}

func (u *userService) ReadAllUnoccupiedRooms() (*[]domain.Rooms, error) {
	return u.userRepository.ReadAllUnoccupiedRooms()
}

func (u *userService) ReadyUpRoomByID(id *uint) error {
	room, err := u.userRepository.ReadRoomByID(*id)
	if err != nil {
		return err
	}
	room.Status = domain.UNOCCUPIED
	_, err = u.userRepository.UpdateRoom(room)
	return err
}

// func (u *userService) BookingRoom {
// 		//1. Check tanggal mau kapan sampai kapan (start_date end_date), atau jadi permanent resident

// 		//2. Pilih kamar yang tersedia (room, nanti kan comot semua)
//		//Note :
//		//-Ijo itu artinya tersedia, ditandai dengan room.Status = UNOCCUPIED (terlepas isinya ada residentID atau nggak)
//		//-Kuning itu artinya ditempatin orang lain, tapi cuman sebagian tanggal, ditandai dengan room.Status = OCCUPIED dan start_date < end_occupied_date < end_date
//		//-Merah itu artinya kuning tapi full semua tanggal, ditandai dengan room.Status = OCCUPIED sama end_date < end_occupied_date
//		//-Item itu artinya ditempatin sama yang tinggal permanen, ditandai dengan room.Status = OCCUPIED sama end_date == nil (dan juga resident.stayType nya permanent)

// 		//2a. CreateNewResident(StayType) (resident),

// 		//3. Bayar	(cuman check berhasil atau nggak)

// 		//3a. terus CheckInRoom()

// 		//4. Username dan Password (kalau Permanent Resident)

// 		//5. Isi informasi resident, terus update informasinya ke resident dummy tadi
// }

//HELPER FUNCTION
func WhatsappAPIGenerate(number uint, text string) string {
	return "https://api.whatsapp.com/send?phone=+62" + strconv.Itoa(int(number)) + "&text=" + text
}

func (u *userService) CreateNewResident(stayType domain.StayType, fullName string,
	NIKNumber uint, Address string, Occupation string, PhoneNumber uint,
	username string, password string) (*domain.Residents, error) {
	temp, err := u.userRepository.ReadLastResident()
	if err != nil {
		return nil, err
	}
	UsernameCheck := false
	if username != "" {
		UsernameCheck = true
	}

	done := false
	i := uint(1)
	var resident domain.Residents
	for !done {
		if UsernameCheck {
			resident = domain.Residents{ID: temp.ID + i,
				StayType: stayType, FullName: fullName,
				NIKNumber: NIKNumber, Address: Address,
				Occupation: Occupation, PhoneNumber: PhoneNumber,
				Username: username, Password: password}

		} else {
			resident = domain.Residents{ID: temp.ID + i,
				StayType: stayType, FullName: fullName,
				NIKNumber: NIKNumber, Address: Address,
				Occupation: Occupation, PhoneNumber: PhoneNumber,
				Username: strconv.Itoa(RandomNumberGenerator(6)), Password: password}
		}
		_, err = u.CreateResident(&resident)
		if err != nil {
			if "ERROR: duplicate key value violates unique constraint \"residents_pkey\" (SQLSTATE 23505)" != err.Error() {
				return nil, err
			} else {
				i++
			}
		} else {
			done = true
		}
	}

	return &resident, err
}

func (u *userService) CreateNewNotification(title *string, description *string) (*domain.Notifications, error) {
	temp, err := u.userRepository.ReadLastNotification()
	if err != nil {
		return nil, err
	}
	notification := domain.Notifications{ID: (*temp).ID + 1, Title: *title, Description: *description}
	_, err = u.userRepository.CreateNotification(&notification)
	return &notification, err
}

func RandomNumberGenerator(n int) int {
	return int(time.Now().UnixMilli() % int64(math.Pow10(n)))
}

type NotificationType int

const (
	BILLS = iota
	PROBLEM
)

func NotificationTitleGenerator(notiftype NotificationType, RoomID uint, ProblemID uint) string {
	switch notiftype {
	case BILLS:
		return "Tagihan untuk kamar No." + strconv.Itoa(int(RoomID)) + " pada tanggal " + time.Now().Format("15/02/2006")
	case PROBLEM:
		return "Masalah#" + strconv.Itoa(int(ProblemID)) + "# pada kamar No. " + strconv.Itoa(int(RoomID))
	}
	return ""
}

func (u *userService) ProblemExtractorFromNotification(notif *domain.Notifications) (*domain.Problems, error) {
	temp := ""
	if !strings.Contains(notif.Title, "Masalah#") {
		return nil, errors.New("Problem object not found")
	}
	start := strings.Index(notif.Title, "Masalah#") + 8
	temp = notif.Title[start : start+strings.Index(notif.Title[start:], "#")]
	ID, err := strconv.Atoi(temp)
	temp2 := uint(ID)
	if err != nil {
		return nil, err
	}
	problem, err := u.ReadProblemByID(temp2)
	return problem, err
}

func (u *userService) HowManyNotificationUnsolved() (int, error) {
	notifications, err := u.ReadAllNotifications()
	if err != nil {
		return 0, err
	}
	return len(*notifications), err
}

func (u *userService) NotificationLinkDetector(s *string) *string {
	temp := ""
	if strings.Contains((*s), "\"http") {
		start := strings.Index((*s), "\"http") + 1
		temp = (*s)[start : start+strings.Index((*s)[start:], "\"")]
	}
	return &temp
}

func DateParser(s string) (time.Time, error) {
	// return time.Parse("2022-06-15", *s)

	temp := strings.Split(s, "-")
	year, err := strconv.Atoi(temp[0])
	if err != nil {
		return time.Time{}, err
	}
	month, err := strconv.Atoi(temp[1])
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(temp[2])
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 1, 0, time.Local), nil

}
