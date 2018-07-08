package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/albert-widi/sqlimporter"
	"github.com/albert-widi/sqlimporter/print"
	"github.com/spf13/cobra"
)

// global variable from global flags
var (
	VerboseFlag bool
	timeStart   time.Time

	// for sqlimporter CLI
	dbName          string
	host            string
	port            string
	filesDir        string
	userandpassword string
	waitTime        string
)

func initCMD() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sqlimporter [command]",
		Short: "sqlimporter command line tools",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			print.SetVerbose(VerboseFlag)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			elapsedTime := time.Since(timeStart).Seconds()
			print.Info("Command <<", cmd.CommandPath(), args, ">> running in", fmt.Sprintf("%.3fs", elapsedTime))
		},
	}
	// add flags
	rootCmd.PersistentFlags().BoolVarP(&VerboseFlag, "verbose", "v", false, "sqlimporter verbose output")
	rootCmd.PersistentFlags().StringVar(&dbName, "db", "", "database name")
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "host name")
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "port of host")
	rootCmd.PersistentFlags().StringVarP(&filesDir, "filedir", "f", "", "directory of sql files")
	rootCmd.PersistentFlags().StringVarP(&userandpassword, "user", "u", "", "username of database")
	rootCmd.PersistentFlags().StringVarP(&waitTime, "wait", "w", "", "wait time")

	timeStart = time.Now()
	return rootCmd
}

func main() {
	rootCmd := initCMD()
	registerImporterCommand(rootCmd)
	rootCmd.Execute()
}

func registerImporterCommand(root *cobra.Command) {
	cmds := []*cobra.Command{
		{
			Use:   "import [driver-name] -db dbname -h hostname -p port -f 'directory'",
			Short: "import postgresql/mysql schema from directory",
			Args:  cobra.MinimumNArgs(1),
			Run: func(c *cobra.Command, args []string) {
				driver := args[0]
				if driver == "" {
					print.Error(errors.New("database driver cannot be empty"))
				}
				if port == "" {
					port = "5432"
				}

				dsn := fmt.Sprintf("%s://%s@%s:%s?sslmode=disable", driver, userandpassword, host, port)
				print.Debug("dsn:", dsn)

				// parse wait time
				var waitUntil time.Time
				if waitTime != "" {
					waitDuration, err := time.ParseDuration(waitTime)
					if err != nil {
						print.Fatal(fmt.Errorf("Invalid wait time %v", err))
					}
					waitUntil = time.Now().Add(waitDuration)
				}
				ticker := time.NewTicker(time.Second * 3)

				data := importData{
					Driver:   driver,
					DbName:   dbName,
					DSN:      dsn,
					FilesDir: filesDir,
				}
				err := importToDB(data)
				if err != nil {
					if waitTime == "" {
						print.Fatal(err)
						return
					}
					print.Error(err)
				}

				if err == nil {
					print.Info("Successfully import schema from", filesDir)
					return
				}

				for {
					select {
					case tt := <-ticker.C:
						err := importToDB(data)
						if err != nil {
							if tt.Before(waitUntil) {
								print.Error(err)
								continue
							}
							print.Fatal(err)
						}
						print.Info("Successfully import schema from", filesDir)
						return
					}
				}
			},
		},
		{
			Use:   "test [args]",
			Short: "test command for sqlimporter",
			Run: func(c *cobra.Command, args []string) {
				fmt.Println(args)
			},
		},
	}
	root.AddCommand(cmds...)
}

type importData struct {
	Driver   string
	DbName   string
	DSN      string
	FilesDir string
}

func importToDB(data importData) error {
	db, drop, err := sqlimporter.CreateDB(data.Driver, data.DbName, data.DSN)
	if err != nil {
		return err
	}

	err = sqlimporter.ImportSchemaFromFiles(context.TODO(), db, data.FilesDir)
	if err != nil {
		print.Error(err)
		if err := drop(); err != nil {
			return err
		}
		return fmt.Errorf("Failed to execute sql files, dropping database %s", data.DbName)
	}
	return nil
}
