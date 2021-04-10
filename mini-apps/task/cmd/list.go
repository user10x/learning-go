package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"task/db"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists the tasks to do",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err :=db.AllTasks()
		if err != nil{
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks! why not take a vacation?🏖🐚👜")
		}
		fmt.Println("You have the following tasks command")
		for i, task := range tasks{
			fmt.Printf("%d. %s, KEY=%d\n", i+1,task.Value, task.Key)
		}

	},
}

func init()  {
	RootCmd.AddCommand(listCmd)
}
