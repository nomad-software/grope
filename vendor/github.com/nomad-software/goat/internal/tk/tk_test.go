package tk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tclListTest struct {
	str    string
	result []string
}

func TestParseTclList(t *testing.T) {
	tests := []tclListTest{
		{"{{The Shawshank Redemption}} 1994 1", []string{"{The Shawshank Redemption}", "1994", "1"}},
		{"{\"The Godfather\"} 1972 2", []string{"\"The Godfather\"", "1972", "2"}},
		{"{[The Godfather: Part II]} 1974 3", []string{"[The Godfather: Part II]", "1974", "3"}},
		{"{$The Dark Knight} 2008 4", []string{"$The Dark Knight", "2008", "4"}},
		{"{Pulp Fiction} 1994 5", []string{"Pulp Fiction", "1994", "5"}},
		{"{The Good, the Bad and the Ugly} 1966 6", []string{"The Good, the Bad and the Ugly", "1966", "6"}},
		{"{Schindler's List} 1993 7", []string{"Schindler's List", "1993", "7"}},
		{"{Angry Men} 1957 8", []string{"Angry Men", "1957", "8"}},
		{"{The Lord of the Rings: The Return of the King} 2003 9", []string{"The Lord of the Rings: The Return of the King", "2003", "9"}},
		{"{Fight Club} 1999 10", []string{"Fight Club", "1999", "10"}},
		{"{{Fight}} 1999 11", []string{"{Fight}", "1999", "11"}},
		{"{[Fight]} 1999 12", []string{"[Fight]", "1999", "12"}},
		{"{\"Fight\"} 1999 13", []string{"\"Fight\"", "1999", "13"}},
		{"{$Fight} 1999 14", []string{"$Fight", "1999", "14"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := parseTclList(test.str)
		assert.Equal(t, test.result, result)
	}
}
