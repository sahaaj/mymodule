package main

import (
	"fmt"
	sample "github.com/elliotforbes/test-package"
	"github.com/spf13/cobra"
	"mymodule/mypackage"
)

func main() {
	fmt.Println("Hello From GO!")
	mypackage.PrintHello()
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Modules")
		},
	}
	fmt.Println("Calling cmd.Execute()")
	cmd.Execute()
	test-package.MySampleFunc()
}
