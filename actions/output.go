package actions

import (
	"fmt"
)

// echo "::workflow-command parameter1={data},parameter2={data}::{command value}"

func SetOutput(o string) error {
	_, err := fmt.Printf("::set-output name=link_objects::%s\n", o)
	if err != nil {
		return err
	}

	return nil
}
