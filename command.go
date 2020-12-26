package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"fmt"
)

type Database struct {
	Name string
}
type Table struct{
	Name string
}

func (db Database) GetTables() ([]Tables error) {
	var tables []Table
	//sudo mysql -e "SHOW TABLES FROM chocoro"
	cmd := exec.Command(path, "-e", fmt.Sprintf("SHOW TABLES %v", db.Name))	
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Output() failed with %s\n", err)
	}
	for _, value:=range strings.Split(string(out), "\n"){
		  tables=append(tables,value)
   } 
   return tables err
}


func ListDatabases()([]Database error){

	path, err := exec.LookPath("mysql")
	if err != nil {
		log.Fatal("Mysql not installed!")
	}
	cmd := exec.Command(path, "-e", "SHOW DATABASES")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Output() failed with %s\n", err)
	}
	databases := make([]Database, 20, 20)
	for _, value:=range strings.Split(string(out), "\n"){
		 databases =append(database, value)
	} 
	return databases err	
}
//func DatabaseProfile( db Database){

//}


