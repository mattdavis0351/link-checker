package actions

import (
	"fmt"
)

// echo "::workflow-command parameter1={data},parameter2={data}::{command value}"

func SetOutput() error {
	_, err := fmt.Println("::set-output name=link_objects::an array of objects")
	if err != nil {
		return err
	}
	return nil
}
