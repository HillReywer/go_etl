// Connect to MySQL

func connect() {
	db, err := sql.Open("DbName", "DbUser:DbPassword@tcp(DbHost:DbPort)/DbName")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}

