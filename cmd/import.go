/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zzzeep/rcndb/parser"
	"github.com/zzzeep/rcndb/storage"
)

var isUrlList bool

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import data",
	Long:  `Import from httpx json output, list of domain, or list of urls`,
	Run: func(cmd *cobra.Command, args []string) {
		if isUrlList {
			importList()
		} else {
			importHttpx()
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolVarP(&isUrlList, "list", "l", false, "input is a simple list of urls / domains")
}

func importHttpx() {
	bytes, _ := io.ReadAll(os.Stdin)
	content := string(bytes)
	records := parser.ParseHttpx(content)

	storage.ImportHttpxRecords(records)
}

func importList() {
	bytes, _ := io.ReadAll((os.Stdin))
	content := string(bytes)
	domains := strings.Split(content, "\n")

	storage.ImportDomainList(domains)
}
