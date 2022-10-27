package core

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func getDaysForYear(year int) int {
	if year%4 == 0 && year%100 == 0 || year%400 == 0 { // leap year
		return 366
	}
	return 365
}

func computeBirthday(Birthday string) int {
	timeNow := time.Now()

	BirthdayTime := strings.Split(Birthday, "-")
	monthOfBirthTime, err := strconv.Atoi(BirthdayTime[1])
	if err != nil {
		log.Fatal("Failed to convert month of birthday")
	}
	dayOfBirthTime, err := strconv.Atoi(BirthdayTime[2])
	if err != nil {
		log.Fatal("Failed to convert day of birthday")
	}
	BirthdayThisYear := time.Date(timeNow.Year(), time.Month(monthOfBirthTime), dayOfBirthTime, 0, 0, 0, 0, time.UTC)

	if timeNow.After(BirthdayThisYear) {
		return getDaysForYear(timeNow.Year()) - int(timeNow.Sub(BirthdayThisYear).Hours()/24)
	} else {
		return int(timeNow.Sub(BirthdayThisYear).Hours() / 24)
	}
}

func ComposeMessage(username, Birthday string) string {
	nextBirthdayDays := computeBirthday(Birthday)
	switch nextBirthdayDays {
	case 0:
		return "Hello, " + username + "! Happy birthday!"
	case 1:
		return "Hello, " + username + "! " + "Your birthday is in " + strconv.Itoa(nextBirthdayDays) + " day"
	default:
		return "Hello, " + username + "! " + "Your birthday is in " + strconv.Itoa(nextBirthdayDays) + " days"
	}
}
