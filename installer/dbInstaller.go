package installer

import (
	"airdb/config"
	"airdb/helpers"
	"airdb/models"
	"database/sql"
	"errors"
	"fmt"
	"kyaps-api/app/libraries"
	"os"
	"strings"
	"time"

	"github.com/howeyc/gopass"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func dbInstaller() {
	fmt.Println("Creating and Configuring airdb database... ")

	fmt.Print("Enter a db name: ")
	var dbName string
	_, err := fmt.Scanf("%s\n", &dbName)
	if err != nil {
		fmt.Println("Enter database info to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("dbName", dbName)
	fmt.Print("Enter your db username: ")

	var dbUsername string
	_, err = fmt.Scanf("%s\n", &dbUsername)
	if err != nil {
		fmt.Println("Enter database info to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("dbUsername", dbUsername)
	if err != nil {
		fmt.Println("Exception occured while setting your db username error " + err.Error())
		os.Exit(103)
	}

	fmt.Print("Enter your db password: ")
	var dbPassword []byte
	dbPassword, err = gopass.GetPasswd()
	if err != nil {
		fmt.Println("Enter database info to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("dbPassword", string(dbPassword))
	if err != nil {
		fmt.Println("Exception occured while setting your db password error " + err.Error())
		os.Exit(103)
	}

	fmt.Print("Enter your db host: ")
	var dbHost string
	_, err = fmt.Scanf("%s\n", &dbHost)
	if err != nil {
		fmt.Println("Enter database info to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("dbHost", dbHost)
	if err != nil {
		fmt.Println("Exception occured while setting your db host error " + err.Error())
		os.Exit(103)
	}

	fmt.Print("Enter your db host port: ")
	var dbPort string
	_, err = fmt.Scanf("%s\n", &dbPort)
	if err != nil {
		fmt.Println("Enter database info to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("dbPort", dbPort)
	if err != nil {
		fmt.Println("Exception occured while setting your db host port error " + err.Error())
		os.Exit(103)
	}
	err = initDB()
	if err != nil {
		fmt.Println("Exception occured while creating airdb database " + err.Error())
		os.Exit(103)
	}
	fmt.Println("Database setup complete.. \n")
	fmt.Println("Create your master web credentials \n")
	fmt.Print("Enter your account name: ")
	var name string
	_, err = fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("Enter name info to continue error " + err.Error())
		os.Exit(103)
	}
	fmt.Print("Enter your username: ")
	var username string
	_, err = fmt.Scanf("%s\n", &username)
	if err != nil {
		fmt.Println("Enter username info to continue error " + err.Error())
		os.Exit(103)
	}
	fmt.Print("Enter your password: ")
	var password []byte
	password, err = gopass.GetPasswd()
	if err != nil {
		fmt.Println("Enter password info to continue error " + err.Error())
		os.Exit(103)
	}
	err = createMasterUser(name, username, string(password))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(103)
	}
	fmt.Println("Master user setup complete.. \n")
}

func createMasterUser(name string, username string, password string) error {
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name, config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		return err
	}
	user := models.User{}
	if db.Where("roles = ?", "gaurdian").First(&user).RecordNotFound() {
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			pass, err := libraries.HashPassword(password)
			if err != nil {
				return err
			}
			user.Username = username
			user.Password = pass
			user.Name = name
			user.Roles = "gaurdian"
			user.Status = "Active"
			user.CreatedAt = time.Now()
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("Master Username already exits")
		}
	} else {
		return errors.New("Master user already exists on this database")
	}
}

func initDB() error {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbusername, err := helpers.Getenv("dbUsername")
	if err != nil {
		return errors.New("db username not found")
	}
	dbpassword, err := helpers.Getenv("dbPassword")
	if err != nil {
		return errors.New("db password not found")
	}
	dbhost, err := helpers.Getenv("dbHost")
	if err != nil {
		return errors.New("db host not found")
	}
	dbname, err := helpers.Getenv("dbName")
	if err != nil {
		return errors.New("db name not found")
	}
	dbport, err := helpers.Getenv("dbPort")
	if err != nil {
		return errors.New("db port not found")
	}
	var (
		userSchema = `CREATE TABLE IF NOT EXISTS ` + dbname + `.users (
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(200) COLLATE utf8_unicode_ci NOT NULL,
			status varchar(100) COLLATE utf8_unicode_ci NOT NULL,
			username varchar(30) COLLATE utf8_unicode_ci NOT NULL,
			password varchar(300) COLLATE utf8_unicode_ci NOT NULL,
			roles varchar(300) COLLATE utf8_unicode_ci NOT NULL,
			created_at datetime NULL DEFAULT NULL,	
			updated_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		 ) ENGINE=InnoDB`
		dbSchema = `CREATE TABLE IF NOT EXISTS ` + dbname + `.dbs (
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(200) COLLATE utf8_unicode_ci NOT NULL,
			dbschema varchar(100) COLLATE utf8_unicode_ci NOT NULL,
			host varchar(30) COLLATE utf8_unicode_ci NOT NULL,
			port varchar(20) COLLATE utf8_unicode_ci NOT NULL,
			username varchar(300) COLLATE utf8_unicode_ci NOT NULL,
			password varchar(300) COLLATE utf8_unicode_ci NOT NULL,
			userId int(10) COLLATE utf8_unicode_ci NOT NULL,
			created_at datetime NULL DEFAULT NULL,	
			updated_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB`
		URL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
	)
	var d *sql.DB
	pr := strings.Split(URL, "/")
	if len(pr) != 2 {
		return errors.New("Invalid database schema url")
	}
	if len(pr[1]) == 0 {
		return errors.New("Invalid database name")
	}
	url := pr[0]
	database := pr[1]
	if d, err = sql.Open("mysql", url+"/"); err != nil {
		return err
	}
	if _, err := d.Exec("CREATE DATABASE IF NOT EXISTS " + database); err != nil {
		return err
	}
	if _, err := d.Exec(userSchema); err != nil {
		return err
	}
	if _, err := d.Exec(dbSchema); err != nil {
		return err
	}
	return nil
}
