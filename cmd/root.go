package cmd

import (
	"fmt" // for printing messages
	"os"
	"strconv" // for converting string args to int (for task IDs)
	"github.com/spf13/cobra" // Cobra CLI framework
)

var root = &cobra.Command{
	Use: "todo",
	Short: "A simple todo-cli",
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Root command executed!")
  },
}

func Execute(){
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
