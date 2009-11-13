package main

import "fmt"
import "sqlite"

func main() {
	sqlite.Initialize();
	fmt.Printf("Sqlite3 Version: %s\n", sqlite.LibVersion());

	dbh := new(sqlite.Handle);
	dbh.Open("test.db");

	st,_ := dbh.Prepare("CREATE TABLE foo (i INTEGER, s VARCHAR(20));");
	st.Step();

	dbh.Close();
	sqlite.Shutdown();
}