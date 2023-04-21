/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zzzeep/rcndb/output"
)

var showOpt output.ShowOptions

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show data from db",
	Long:  `show saved data with various filters/options`,
	Run: func(cmd *cobra.Command, args []string) {
		output.SimplePrintOut(showOpt)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().BoolVarP(&showOpt.Domains, "domains", "d", false, "show domains")
	showCmd.Flags().BoolVarP(&showOpt.Urls, "urls", "u", false, "show urls")
	showCmd.Flags().BoolVarP(&showOpt.IPs, "ips", "i", false, "show web-server")
	showCmd.Flags().BoolVarP(&showOpt.Ports, "ports", "p", false, "show ports")
	showCmd.Flags().BoolVarP(&showOpt.Status, "status-code", "s", false, "show status code")
	showCmd.Flags().BoolVarP(&showOpt.Webserver, "web-server", "w", false, "show web-server")
	showCmd.Flags().BoolVarP(&showOpt.Content, "content-info", "c", false, "show content")
	showCmd.Flags().BoolVarP(&showOpt.LastScan, "last-scan-date", "l", false, "show last scan date")

	showCmd.Flags().UintVar(&showOpt.FilterStatus, "fs", 0, "filter by status code")
	showCmd.Flags().StringVar(&showOpt.FilterIP, "fi", "", "filter by ip")
	showCmd.Flags().StringVar(&showOpt.FilterPort, "fp", "", "filter by port")

	showCmd.Flags().BoolVar(&showOpt.NoColor, "nc", false, "disable colored output")
	showCmd.Flags().BoolVar(&showOpt.NoTruncation, "nt", false, "disable long url truncation")
	showCmd.Flags().BoolVar(&showOpt.NoEmptyResponse, "ne", false, "hide urls with empty response")
}
