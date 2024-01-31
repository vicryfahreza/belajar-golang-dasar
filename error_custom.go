package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Validation error"}
	}
	if id != "Vicry" {
		return &NotFoundError{"Data Not Found"}
	}
	return nil
}

func main() {
	err := saveData("eko", nil)
	if err != nil {
		// if validationErr, ok := err.(*ValidationError); ok {
		// 	fmt.Println("validation error: ", validationErr.Error())
		// } else if notFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("not found error: ", notFoundErr.Error())
		// } else {
		// 	fmt.Println("unknown error: ", err.Error())
		// }
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("validation error: ", finalError.Error())
		case *NotFoundError:
			fmt.Println("not found error: ", finalError.Error())
		default:
			fmt.Println("unknown error: ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}

}
