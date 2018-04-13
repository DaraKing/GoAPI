package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func openDB()  *sql.DB{
	db, err := sql.Open("mysql","root:dario123@/accounts")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func QueryItAndJSON() ([]User, error) {

	db := openDB()
	var user User
	var users []User
	//QUERY DATABASE
	rows, err := db.Query("SELECT id,first_name,last_name,age,gender,email FROM Users")
	checkErr(err)

	for rows.Next() {

		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.Email)

		users = append(users, user)
	}

	db.Close()

	return users, err
}

func insertInto(firstName string, lastName string,age int, gender string, email string) bool {

	//INSERTING INTO
	db := openDB()
	stmIns, err := db.Prepare("INSERT INTO Users(first_name, last_name, age, gender ,email) VALUES (?,?,?,?,?)")
	checkErr(err)

	defer stmIns.Close()

	_ ,err = stmIns.Exec(firstName, lastName, age, gender, email)

	if err != nil {
		panic(err.Error())
		return false
	}

	return true
}

func deleteUserInDB(id int) bool{

	db := openDB()
	stmt, err := db.Prepare("DELETE FROM Users WHERE id=?")

	checkErr(err)

	_, result := stmt.Exec(id)

	defer stmt.Close()

	if result != nil {
		panic(result.Error())
		return false
	}

	return true
}

func updateUser(id int, firstName string, lastName string, age int, email string, gender string) bool{

	db := openDB()
	stmt, err := db.Prepare("UPDATE Users SET first_name=?, last_name=?, age=?, email=?, gender=? WHERE id=?")

	checkErr(err)

	result, err := stmt.Exec(firstName,lastName,age,email,gender,id)

	defer stmt.Close()

	var numResults, _ = result.RowsAffected()

	if err != nil || numResults != 1 {
		return false
	}

	return true
}