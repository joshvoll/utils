/*
	Developer: josue rodriguez
	Email: joshvoll@yahoo.com

	// convert struct interface{}  to 2d array []interface{}
*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

// COLTAG is the title of the tag properties
const COLTAG = "col"

// ToSlice Will generate slice from and struct, is going to resutn and []interface{}
func ToSlice(src interface{}) []interface{} {

	// first defining a type interface or string
	ret := []interface{}{}

	// Using reflect the go standar librery to modify different types
	if v := reflect.ValueOf(src); v.Kind() == reflect.Slice {
		// looping throw the struct
		for i := 0; i < v.Len(); i++ {
			// append the struct to the string of []interface{}
			ret = append(ret, v.Index(i).Interface())
		}
	} else {

		ret = append(ret, v.Interface())
	}

	return ret

}

// GenerateRows method is going to generate a row from []interface{}, and genearte a 2d array [][]string
// This was a special problem that i have for generateing a csv
func GenerateRows(src interface{}) [][]string {

	// First generate the slice of the struct for example [{name, point, age}, {name, point, age}], using the ToSlice method
	sl := ToSlice(src)

	// define the rows and ingoreColIndex
	// rows will hold the result of the data slice on the new type [][]string
	// ignoreColIndex this is use for if you want to use it as title of pdf or csv(which was the problem from the start )
	rows := make([][]string, 1)
	ignoreColIndex := map[int]bool{}

	// loop throw the slice
	for n, d := range sl {
		// 1.- append the string to rows
		rows = append(rows, []string{})
		v := reflect.ValueOf(d)

		// 2.- loop to get the field from the slice
		for i := 0; i < v.NumField(); i++ {
			// if the first row is 0, that mean is the title of the array, remember this is for pdf or csv ;)
			if n == 0 {
				// Header
				columnName := v.Type().Field(i).Tag.Get(COLTAG)

				// if there is no column name defined
				if columnName == "" {
					columnName = strings.ToLower(v.Type().Field(i).Name)

				} else if columnName == "-" {
					// we're going to ignore the index
					ignoreColIndex[i] = true
					continue
				}

				// append the first row(row with 0 index to the rows)
				//remember this append is waiting for type string not type [][]string at the moment
				// so we can't append the entire rows b'cos rows are => []string{}
				rows[0] = append(rows[0], columnName)
			}

			if !ignoreColIndex[i] {
				// the secret salse, converting the slice byte to [][]string{}
				rows[len(rows)-1] = append(rows[len(rows)-1], fmt.Sprint(v.Field(i).Interface()))
			}
		}
	}

	// return everything
	return rows
}
