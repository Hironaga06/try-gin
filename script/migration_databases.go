package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ngaut/log"
)

const (
	schemaPath  = "/migrations/schema"
	databaseURL = "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
)

var (
	command               = flag.String("exec", "", "set up or down as a argument")
	force                 = flag.Bool("f", false, "force exec fixed sql")
	availableExecCommands = map[string]string{
		"up":   "Execute up schema",
		"down": "Execute down schema",
	}
)

func main() {
	flag.Parse()
	if len(*command) < 1 {
		fmt.Println("no argument")
		showUsageMessge()
		os.Exit(1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	m, err := migrate.New("file:"+currentDir+schemaPath, databaseURL)
	if err != nil {
		fmt.Println(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		fmt.Println(err)
	}
	applyQuery(m, version, dirty)
}

//exec up or down sqls
//with force option if needed
func applyQuery(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	switch *command {
	case "up":
		if err := m.Up(); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	default:
		fmt.Println("error: invalid command" + *command)
		showUsageMessge()
		os.Exit(1)
	}

	fmt.Printf("success exec %s\n", *command)

	version, dirty, err := m.Version()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("updated version info")
	showVersionInfo(version, dirty)
}

func showUsageMessge() {
	fmt.Println("======================================")
	fmt.Println("Usege -exec <command>")
	fmt.Println("Available exec commands")
	for command, detail := range availableExecCommands {
		fmt.Println(command + " : " + detail)
	}
	fmt.Println("======================================")
}

func showVersionInfo(version uint, dirty bool) {
	fmt.Println("======================================")
	fmt.Println("version: ", version)
	fmt.Println("dirty: ", dirty)
	fmt.Println("======================================")
}
