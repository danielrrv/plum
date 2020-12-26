package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"flag"
	"os"
	//"log"
	//"io/ioutil"
	//"os/exec"
)

func usage() {
	format := `Usage:
  plum [flags]
Flags:
  --global              Set admin and  password as global.
Author:
  danielrrv <drodrigo567@gmail.com>
`
	fmt.Fprintln(os.Stderr, format)
}
var (
	isGlobal = flag.Bool("global", false, "Set user as global")

	// these are set in build step
	version = "unversioned"
	commit  = "?"
	date    = "?"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	os.Exit(run())
}
const (
	add = "List Databases"
	mount = "Mount database"
	remove = "Remove database"
	
)
func run() int {
	action := promptui.Select{
		Label: "Select action",
		Items: []string{add,mount, remove},
	}
	_, actionType, err := action.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to select action: %v\n", err)
		return 1
	}

	switch actionType {
	case add:
		if databases, err := ListDatabases(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list database: %v\n", err)
			action.Run()
		}
	//case mount:
	//	if err := addUser(); err != nil {
	//		fmt.Fprintf(os.Stderr, "Failed to add user: %v\n", err)
	//		return 1
	//	}
	//case mount:
	//	if err := deleteUser(); err != nil {
	//		fmt.Fprintf(os.Stderr, "Failed to delete user: %v\n", err)
	//		return 1
	//	}
	default:
		fmt.Fprintf(os.Stderr, "Unexpected action type\n")
		return 1
	}

	return 0
}

//func selectDataBase() error {
//	users, err := ListDatabases()
//	if err != nil {
//		return err
//	}

//	if len(users) == 0 {
//		fmt.Println("No users")
//		return nil
//	}

//	user := promptui.Select{
//		Label: "Select git user",
//		Items: UsersToString(users),
//	}
//	selectedUserIndex, _, err := user.Run()
//	if err != nil {
//		return err
//	}

//	option := "--local"
//	if *isGlobal {
//		option = "--global"
//	}

//	cmdName := exec.Command("git", "config", option, "user.name", users[selectedUserIndex].Name)
//	if err := cmdName.Run(); err != nil {
//		return err
//	}
//	cmdMail := exec.Command("git", "config", option, "user.email", users[selectedUserIndex].Email)
//	if err := cmdMail.Run(); err != nil {
//		return err
//	}

//	return nil
//}
