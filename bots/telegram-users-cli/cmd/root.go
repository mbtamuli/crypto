package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mbtamuli/crypto/sqlite"
	"github.com/spf13/cobra"
)

var (
	dbFile   string
	store    *sqlite.Store
	username string
	chat_id  string
)

var rootCmd = &cobra.Command{
	Use:   "telegram-users-cli",
	Short: "telegram-users-cli gives the ability to create, update, delete users",
	Long:  "telegram-users-cli gives the ability to create, update, delete users",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize()
	var err error

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	defaultDBPath := fmt.Sprintf("%s/data/telegram.sqlite", exPath)

	rootCmd.PersistentFlags().StringVar(&dbFile, "dbfile", defaultDBPath, "sqlite database file")

	store, err = sqlite.NewStore(defaultDBPath)
	if err != nil {
		log.Fatal(err)
	}
}
