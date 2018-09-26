package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/astrocorp42/astroflow-go"
	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
	"github.com/astrocorp42/signal/api/server"
	"github.com/astrocorp42/signal/api/util"
	_ "github.com/astrolib/godotenv/autoload"
	"golang.org/x/crypto/ssh/terminal"
)

var RequiredEnv = []string{
	"DATABASE_URL",
	"JWT_SECRET",
}

func checkEnv() {
	for _, required := range RequiredEnv {
		if os.Getenv(required) == "" {
			log.With("var", required).Fatal("Missing environment variable")
		}
	}

	// optional env
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "8080")
	}
}

func deleteAccount(srv *server.Server) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}
	username = strings.TrimSpace(username)
	err = srv.DB.Where("username = ?", username).Delete(&db.Account{}).Error
	if err != nil {
		log.Fatal(err.Error())
	}
}

func registerAccount(srv *server.Server) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err.Error())
	}
	password := strings.TrimSpace(string(bytePassword))
	username = strings.TrimSpace(username)

	hashedPassword, err := util.HashString(password, db.BcryptCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = srv.DB.Create(&db.Account{Username: username, Password: hashedPassword}).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print("\n")
}

func listAccounts(srv *server.Server) {
	accounts := []db.Account{}

	err := srv.DB.Find(&accounts).Error
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, account := range accounts {
		fmt.Println(account.Username)
	}
}

func main() {
	var err error
	// init
	log.Config(
		astroflow.SetFormatter(astroflow.NewConsoleFormatter()),
	)
	checkEnv()

	flagMigrate := flag.Bool("migrate", false, "auto migrate")
	flagRegister := flag.Bool("register", false, "register a new admin")
	flagDelete := flag.Bool("rm", false, "remove an existing admin")
	flagLs := flag.Bool("ls", false, "list all existing admins")
	flagPort := flag.String("p", os.Getenv("PORT"), "port to listen to")
	flag.Parse()

	srv, err := server.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// migrate if the -migrate flag is provided
	if *flagMigrate {
		err = srv.AutoMigrate()
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if *flagRegister {
		registerAccount(srv)
		return
	}

	if *flagDelete {
		deleteAccount(srv)
		return
	}

	if *flagLs {
		listAccounts(srv)
		return
	}

	// run server
	log.With("port", *flagPort).Info("listenning...")
	log.Fatalf("%s", srv.Run(*flagPort))
}
