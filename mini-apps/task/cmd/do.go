package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/db"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Marks the task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids [] int
		for _,arg := range args{
			id, err := strconv.Atoi(arg)
			if err != nil{
				fmt.Println("failed to parse the argo", arg)
			}else{
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err !=nil {
			fmt.Println("Something went wrong")
		}
		for _,id := range ids{
			if id <=0 || id > len(tasks){
				fmt.Printf("Invalid task number: %d\n", id)
				continue
			}
			task := tasks[id-1]

			err := db.DeleteTasks(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed %s\n\n", id, err)
			}else {
				fmt.Printf("Marked  \"%d\" as completed \n", id)
			}
		}

	},
}

func init()  {
	RootCmd.AddCommand(doCmd)
}
