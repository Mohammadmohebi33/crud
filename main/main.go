package main

import (
	"encoding/json"
	 "fmt"
	"log"
	"os"
	"workWithDatabase/dbtools"
)


type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}



func main(){

	file , err := os.Open("../config/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	conf := new(Configuration)
	json.NewDecoder(file).Decode(conf)
	dbtools.DBInitializer(conf.DriverName , conf.DataSourceName)
	


	//student := modles.Student{ID: 3 , Name: "amirrrrrrrrr" , Age: 20}


	// t := dbtools.Delete(3)
	// fmt.Println(t)

	
    // dbtools.Save(student)

	students := dbtools.GetAllStudents()

	for _ , val := range students {

		fmt.Println(val.Name , " " , val.Age)
	}

	// student_3 := dbtools.GetAllStudentById(1)

	// fmt.Println(student_3)
}
