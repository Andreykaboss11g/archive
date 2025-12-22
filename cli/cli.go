package cli

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const fileExtension = "vlc"

var rootCmd = &cobra.Command{
	Use:   "archiver",
	Short: "Simple archiver tool",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack files into an archive",
}

func init() {
	rootCmd.AddCommand(packCmd)
}

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Play media files",
	Run:   pack,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal("Usage: archiver vlc <file>")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// packed := Encode(data)

	packed := ""
	log.Printf(string(data))

	err = os.WriteFile(fileName(filePath), []byte(packed), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fileName(path string) string {
	base := filepath.Base(path)

	return strings.TrimSuffix(base, filepath.Ext(base)) + "." + fileExtension
}
