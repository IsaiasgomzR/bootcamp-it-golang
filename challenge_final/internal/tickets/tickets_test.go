package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t testing.T){
	ReadDataFile("../../tickets.csv")

	t.Run("should get total person for country", func (t testing.T)  {
		country:= "Brazil"
		valueExpected:= 45

		totalResult, err := GetTotalTickets(country)

		assert.Equal(t, valueExpected, totalResult)
		assert.Nil(t, err)
	})

	t.Run("should get error for country", func (t testing.T)  {
		country:="Guya"
		valueExpected:= 0

		totalResult, err := GetTotalTickets(country)

		assert.Equal(t, valueExpected, totalResult)
	})
}



func TestGetmornings(t testing.T)  {
	t.Run("should get total for a time", func (t testing.T)  {
		time:= "manana"
		valueExpected:= 250

		totalResult,err := GetMornings(time)

		assert.Equal(t , valueExpected, totalResult)
		assert.NotNil(t, err)
	})
	
	t.Run("should get total 0 for time", func (t testing.T)  {
		time:= "medianoche"
		valueExpected:= 0

		totalResult, err := GetMornings(time)

		assert.Equal(t , valueExpected, totalResult)
		assert.NotNil(t, err)
	})
}



func TestAverageDestination(t testing.T)  {
	t.Run("should get average destination for a time", func (t testing.T)  {
		country:= "Brazil"
		total:= 10
		valueExpected:= 4.0

		averageResult, err := AverageDestiantion(country, total)

		assert.Equal(t, valueExpected, averageResult)
		assert.NotNil(t, err)
	})

	t.Run("should get average destination average for a time", func (t testing.T)  {
		country:= "Guya"
		total:= 10
		valueExpected:= 0.0

		averageResult, err := AverageDestiantion(country, total)

		assert.Equal(t, valueExpected, averageResult)
		assert.NotNil(t, err)
	})

	t.Run("should get error cannot divide by zero", func (t testing.T)  {
		country:= "Guya"
		total:= 0
		valueExpected:= 0

		averageResult, err:= AverageDestiantion(country, total)

		assert.Equal(t valueExpected, averageResult)
		assert.NotNil(t, err)
	})
}