package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverageNotes(t testing.T)  {
	t.Run("average betwent two values", func (t testing.T)  {
		note1:= 1
		note2 := 5

		valueExpeted:= 2

		result := averageNotes(note1, note2)

		assert.Equal(t, valueExpeted, result)
	})
}