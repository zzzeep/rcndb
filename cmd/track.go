/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zzzeep/rcndb/output"
	"github.com/zzzeep/rcndb/storage"
)

var trackOpt output.TrackOptions

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "display changes",
	Long:  `display changes for a given domain/url if any`,
	Run: func(cmd *cobra.Command, args []string) {
		if !trackOpt.All {
			trackDomain()
		} else {
			trackAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)

	trackCmd.Flags().StringVarP(&trackOpt.Domain, "domain", "d", "", "show changes for Domain")
	trackCmd.Flags().StringVarP(&trackOpt.Url, "url", "u", "/", "show changes for URL")
	trackCmd.Flags().BoolVar(&trackOpt.All, "dump", false, "show all changes")
	trackCmd.Flags().BoolVar(&trackOpt.NoColor, "nc", false, "show all changes")
}

func trackDomain() {
	output.InitTable(trackOpt.NoColor)
	if len(trackOpt.Domain) > 0 {
		d := storage.Domain{
			Domain: trackOpt.Domain,
		}

		dChanges, err := storage.GetDomainChanges(&d)
		if err == nil {
			output.PrintDomainChanges(dChanges)

			uChanges, err := storage.GetUrlChanges(d, trackOpt.Url)
			if err == nil {
				output.PrintURLChanges(uChanges)
			}
		}
	}
	output.Render()
}

func trackAll() {
	output.InitTable(trackOpt.NoColor)
	domains, err := storage.GetAllDomains()
	if err == nil {
		for _, d := range domains {
			dChanges, err := storage.GetDomainChanges(&d)
			output.PrintDomainChanges(dChanges)

			urls, err := storage.GetDomainUrls(d)
			if err == nil {
				for _, u := range urls {
					uChanges, err := storage.GetUrlChanges(d, u.Path)
					if err == nil {
						output.PrintURLChanges(uChanges)
					}
				}
			}
		}
	}
	output.Render()
}
