package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/gpsandyka/manajemen_kos/domain"
	"github.com/gpsandyka/manajemen_kos/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRandomNumberGenerator(t *testing.T) {
	n := 8

	t.Run("SingleTest", func(t *testing.T) {
		a := RandomNumberGenerator(n)
		time.Sleep(999999)
		b := RandomNumberGenerator(n)

		assert.NotEqual(t, a, b, "The two number should not be the same.")
	})
	t.Run("MultipleTest", func(t *testing.T) {
		howMuchTest := 8

		RandomTable := make([]int, howMuchTest)

		errorCount := 0

		for i := range RandomTable {
			temp := RandomNumberGenerator(n)
			t.Log("Random on try ", i, "get ", temp)
			for ii := range RandomTable {
				if RandomTable[ii] == temp {
					t.Log("Somehow failed randoming on try number ", ii, ", it's the same on try number ", i, ", coincidence?")
					errorCount++
				}
			}
			RandomTable[i] = temp
			time.Sleep(999999)
		}

		if errorCount > howMuchTest/2 {
			t.Errorf("On %v test attempt, it failed %v times, more than half of the test", howMuchTest, errorCount)
		}
	})

}

func TestDateParser(t *testing.T) {

	t.Run("SimpleDate", func(t *testing.T) {
		s := "2022-06-06"
		timetest, err := DateParser(s)
		assert.Equal(t, nil, err, "Error is not nil, should be nil")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Year(), timetest.Year(), "Year doesn't match")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Month(), timetest.Month(), "Month doesn't match")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Day(), timetest.Day(), "Day doesn't match")
	})
	t.Run("LooseDate", func(t *testing.T) {
		s := "2022 - 06 - 06"
		timetest, err := DateParser(s)
		assert.Equal(t, nil, err, "Error is not nil, should be nil")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Year(), timetest.Year(), "Year doesn't match")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Month(), timetest.Month(), "Month doesn't match")
		assert.Equal(t, time.Date(2022, 6, 6, 0, 0, 0, 0, time.Local).Day(), timetest.Day(), "Day doesn't match")
	})
	t.Run("WrongLayoutDate", func(t *testing.T) {
		s := "06-06-2022"
		_, err := DateParser(s)
		assert.NotEqual(t, nil, err, "Mismatched date, should raise error")
	})
	t.Run("NotEnoughDate", func(t *testing.T) {
		s := "06- -2022"
		_, err := DateParser(s)
		assert.NotEqual(t, nil, err, "Missing month, should raise error")
	})
	t.Run("WrongDate", func(t *testing.T) {
		s := "06-2022-2022"
		_, err := DateParser(s)
		assert.NotEqual(t, nil, err, "No 2022 month, should raise error")
	})
}

var uRepo = &repository.UserRepositoryMock{mock.Mock{}}
var uServi = userService{uRepo}

func TestProblemExtractorFromNotification(t *testing.T) {
	uintID := uint(123)
	problem123 := domain.Problems{}
	uRepo.Mock.On("ReadProblemByID", uintID).Return(&problem123, nil)

	t.Run("EmptyTitleNotif", func(t *testing.T) {
		notif := domain.Notifications{}

		result, err := uServi.ProblemExtractorFromNotification(&notif)

		fmt.Println(result)
		fmt.Printf("%T\n", result)

		assert.NotEqual(t, nil, err, "Should be not nil because error no problem object")
	})
	t.Run("Notif#123", func(t *testing.T) {
		notif := domain.Notifications{Title: "Masalah#123#"}

		result, err := uServi.ProblemExtractorFromNotification(&notif)

		fmt.Println(result)
		fmt.Printf("%T\n", result)

		assert.Equal(t, nil, err, "Should be nil because no error")
		assert.Equal(t, &problem123, result, "Should be empty problem because mocked so")

	})
}
