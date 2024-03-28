package main

const(
	host  = "localhost"
	port = 5432
	user = "root"
	password = "pass123456"
	dbname = "users"
)

type User struct {
	Username string
	PasswordHash string
}

func main(){



}

func setupDatabase() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
	 host, port, username, password, dbname)
	 db, err := sql.Open("postgres", psqlInfo)
	 if err != nil {
		 log.Fatal("Error opening database: ", err)		
	 }
	 err = db.Ping()
	 if err != nil {
		 log.Fatal("Error ping
		 database: ", err)
	 }
	 fmt.Println ("Successfully connected to the PostgreSQL server")
	 return db
		
	 }