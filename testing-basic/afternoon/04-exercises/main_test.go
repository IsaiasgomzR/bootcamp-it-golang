package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatics(t testing.T){
	t.Run("minimmum califications", func (t testing.T)  {
		oper:= operation("minimum")
		valueExpected:= 2
		
		result := oper(2,3,4,5)


		assert.Equal(t, valueExpected, result)
	})
}