// package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSalaryWithCategory(t testing.T){
	t.Run("Calculate Salary with Caegory A", func (t testing.T)  {
		category:= "A"
		minutesWorked:= 200
		var expectResult int = 200/60 *3000 + 9000*50/100

		result := salaryFunction(minutesWorked, category)
	})
}