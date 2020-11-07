package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var days = [7]string{
	"Saturday",
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
}

func main() {
	file, err := os.Open("input-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// reading class infos, including start date, end date, and date of the week
	scanner.Scan()
	classInfos := strings.Split(scanner.Text(), " ")
	startYear, _ := strconv.Atoi(classInfos[0])
	startMonth, _ := strconv.Atoi(classInfos[1])
	startDate, _ := strconv.Atoi(classInfos[2])
	endYear, _ := strconv.Atoi(classInfos[3])
	endMonth, _ := strconv.Atoi(classInfos[4])
	endDate, _ := strconv.Atoi(classInfos[5])
	classDay, _ := strconv.Atoi(classInfos[6])
	classDay %= 7

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid num of students to check in")
	}

	for i := 0; i < num; i++ {
		scanner.Scan()
		checkInInfo := strings.Split(scanner.Text(), " ")
		checkInYear, _ := strconv.Atoi(checkInInfo[0])
		checkInMonth, _ := strconv.Atoi(checkInInfo[1])
		checkInDate, _ := strconv.Atoi(checkInInfo[2])

		valid := false
		if checkInYear >= startYear && checkInYear <= endYear {
			if checkInMonth >= startMonth && checkInMonth <= endMonth {
				if checkInDate >= startDate && checkInDate <= endDate {

					if days[classDay] == getDayOfTheWeek(
						checkInYear,
						checkInMonth,
						checkInDate,
					) {
						valid = true
					}
				}
			}
		}
		if valid {
			fmt.Println("VALID")
		} else {
			fmt.Println("INVALID")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getDayOfTheWeek(year, month, day int) string {
	y := year
	m := month
	if month < 3 {
		y--
		m += 12
	}

	l2y := y % 100
	f2y := y / 100

	d := day

	w := (13 * (m + 1)) / 5
	w += l2y / 4
	w += f2y / 4
	w += d
	w += l2y
	w -= 2 * f2y
	w %= 7
	if w < 0 {
		w += 7
	}
	return days[w]
}
