package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Ticket struct{
	ID int
	Name string
	Email string
	Country string
	Date int
	Price int
}

var tickets = [] Ticket{}

func ReadDataFile(path string){
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvlines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvlines {
		var ticket Ticket

		if value, err := strconv.Atoi(line[0]); err != nil{
			ticket.ID = value
		}
		ticket.Name = line[1]
		ticket.Email = line[2]
		ticket.Country = line[3]
		if value, err := strconv.Atoi(strings.Split(line[4], ":")[0]); err != nil{
			ticket.Date = value
		}
		
		if value, err := strconv.Atoi(line[5]); err != nil{
			ticket.Price = value
		}

		tickets = append(tickets, ticket)
	}
}

func GetTotalTickets(destination string) (int, error){
	if len(tickets) <= 0{
		return 0, errors.New("Not Found Data")
	}
	total:= 0
	for _, v := range tickets {
		if v.Country == destination{
			total++
		}
	}
	return total, nil
}

func GetMornings(time string)(int, error){
	if len(tickets) <= 0{
		return 0, errors.New("Data not Found")
	}
	total := 0
	switch time{
	case "madrugada":
		for _, v := range tickets {
			if v.Date >= 0 && v.Date <= 6{
				total++
			}
		}
		return total,nil
	case "manana":
		for _, v := range tickets {
			if v.Date >= 7 && v.Date <= 12{
				total++
			}
		}
		return total,nil
	case "tarde":
		for _, v := range tickets {
			if v.Date >= 13 && v.Date <= 19{
				total++
			}
		}
		return total,nil
	case "noche":
		for _, v := range tickets {
			if v.Date >= 20 && v.Date <= 23{
				total++
			}
		}
		return total,nil
	default:
		return 0,errors.New("Not Found Time")
	}
}

func AverageDestiantion(destination string, total int) (float64, error){
	totalCountryByTickets, err := GetTotalTickets(destination)
	if err != nil{
		return 0, err 
	}
	if total <= 0{
		return 0, errors.New("Cannot divided by cero")
	}
	return float64(totalCountryByTickets)/float64(total), nil
}


