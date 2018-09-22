package main 

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"reflect"
)

func GetAgeInYear(age string) string {
    fmt.Println("from get Age function dob:",age)
    str := strings.Split(age,"-")
    fmt.Println(str[0])
    fmt.Println(str[1])
    fmt.Println(str[2])

    var year,currYear int
    var month,currMonth int
    var day,currDay int

    if i,err := strconv.Atoi(str[0]);err == nil{
    	year = i
    	//fmt.Println("i: ",i)
    }
    fmt.Println("year: ",year)

    if i,err := strconv.Atoi(str[1]);err == nil{
    	month = i
    	//fmt.Println("i: ",i)
    }
	fmt.Println("month: ",month)

	if i,err := strconv.Atoi(str[2]);err == nil{
    	day = i
    	//fmt.Println("i: ",i)
    }
	fmt.Println("day: ",day)
//-------------------------------

	currentTime := time.Now().Format("2006-01-02")
	fmt.Println(reflect.TypeOf(currentTime))
	str = strings.Split(currentTime,"-")

	 if i,err := strconv.Atoi(str[0]);err == nil{
    	currYear = i
    	//fmt.Println("i: ",i)
    }
    fmt.Println("currYear: ",currYear)

    if i,err := strconv.Atoi(str[1]);err == nil{
    	currMonth = i
    	//fmt.Println("i: ",i)
    }
	fmt.Println("currMonth: ",currMonth)

	if i,err := strconv.Atoi(str[2]);err == nil{
    	currDay = i
    	//fmt.Println("i: ",i)
    }
	fmt.Println("currDay: ",currDay)

//----------------------------------
// start logic

//logic refernce https://www.geeksforgeeks.org/program-calculate-age/
	
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	//fmt.Println(days)

	if day > currDay {
		currMonth = currMonth - 1
		currDay = currDay + days[month - 1]
	}

	if month > currMonth {
		currYear = currYear - 1
		currMonth = currMonth + 12
	}

	var result_day = currDay - day
	var result_month = currMonth - month
	var result_year = currYear - year
	return strconv.Itoa(result_year)+" Years "+strconv.Itoa(result_month)+" Months "+strconv.Itoa(result_day)+" days"
}
