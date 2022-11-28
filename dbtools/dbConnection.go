package dbtools

import (
	"database/sql"
	"fmt"
	"log"
	"workWithDatabase/modles"

	_ "github.com/go-sql-driver/mysql"

)

var driverName string
var dataSourceName string

func DBInitializer(dn, dsn string) {

	driverName = dn
	dataSourceName = dsn
}


func Test(){
	fmt.Println("test")
}

func connect()(db *sql.DB){

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}




func GetAllStudents() []modles.Student{
    db := connect()
	defer db.Close()
	rows , err := db.Query("select * from students")

	if err != nil {
		log.Fatal(err.Error())
	}

	students := [] modles.Student{}

		
	for rows.Next(){
		student := modles.Student{}
		err := rows.Scan(&student.ID , &student.Name , &student.Age)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		students = append(students , student)
	}

	return students
}


func GetAllStudentById(id int) modles.Student{

	db := connect()
	defer db.Close()
	row := db.QueryRow("select * from students where id = ?", id)
	
	student := modles.Student{}
	err := row.Scan(&student.ID, &student.Name, &student.Age)

	if err != nil{
		log.Fatal(err.Error())
	}

    return student
}


func Save(student modles.Student) {

	db := connect()
	defer db.Close()

	save, err := db.Prepare("insert into students (id, name, age) values(?, ? , ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	res  ,err := save.Exec(student.ID , student.Name , student.Age)

	if err != nil {
		log.Fatal(err.Error())
	}

	id, err := res.LastInsertId()
      
	if err != nil {
		log.Fatal(err.Error())
	}
 
    fmt.Println("Insert id", id)
}





func UpdateStudent(student modles.Student) int64{
	db := connect()
	defer db.Close()

	update, err := db.Prepare("update students set name=?, age=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

    res , err := update.Exec(student.Name, student.Age , student.ID)

	if err != nil {
		log.Fatal(err.Error())
	}
    
	row ,err := res.RowsAffected()


	if err != nil {
		log.Fatal(err.Error())
	}
    
	return row
}


func Delete(id int) int64{

	db := connect()
	defer db.Close()

	delete , err := db.Prepare("delete from students where id=?")

	if err != nil{
		log.Fatal(err.Error())
	}

	res , err := delete.Exec(id)


	if err != nil{
		log.Fatal(err.Error())
	}

	row ,err := res.RowsAffected()

	if err != nil{
		log.Fatal(err.Error())
	}

	return row


}