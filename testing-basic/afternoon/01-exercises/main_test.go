package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSalary (t *testing.T)  {
	t.Run(salary less than 50000, func (t *testing.T)  {
		//arrange
		salary := 45600.00
		salaryExpected := 45600.00

		//act
		result := taxSalary(salary)

		assert.Equal(t, salaryExpected, result)
	})
}