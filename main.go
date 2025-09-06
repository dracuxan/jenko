package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

type Fields struct {
	NumberOfFields int
	FieldNames     []string
	FieldTypes     []string
	FieldValues    []any
}

func main() {
	newF := &Fields{}
	newF.GetFieldsFromUser()
	newF.GenerateRandomData()

	for i := range newF.NumberOfFields {
		fmt.Printf("%v\n", newF.FieldValues[i])
	}
}

func (f *Fields) GetFieldsFromUser() {
	fmt.Print("Enter number of fields: ")
	fmt.Scanln(&f.NumberOfFields)

	f.FieldNames = make([]string, f.NumberOfFields)
	f.FieldTypes = make([]string, f.NumberOfFields)

	for v := range f.NumberOfFields {
		fmt.Printf("Enter name of field %d: ", v+1)
		fmt.Scanln(&f.FieldNames[v])
		fmt.Printf("Enter valid type for the field %d (str/int/float):", v+1)
		fmt.Scanln(&f.FieldTypes[v])
	}
	fmt.Println("Preparing field values....")
}

func (f *Fields) GenerateRandomData() {
	f.FieldValues = make([]any, f.NumberOfFields)

	for i := 0; i < f.NumberOfFields; i++ {
		switch f.FieldTypes[i] {
		case "str":
			f.FieldValues[i] = gofakeit.Name()
		case "int":
			f.FieldValues[i] = gofakeit.Number(1, 100)
		case "float":
			f.FieldValues[i] = gofakeit.Float64()
		}
	}
}
