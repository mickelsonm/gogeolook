package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mickelsonm/go-helpers/geocoding"
	"github.com/spf13/cobra"
)

func main() {
	var lookup geocoding.Lookup

	var doAddress = &cobra.Command{
		Use:   "address",
		Short: "Lookup by address",
		Long:  "Looks up geocoding information for a given address.",
		Run: func(cmd *cobra.Command, args []string) {
			lookup = geocoding.Lookup{
				Address: strings.Join(args, " "),
			}

			resp, err := lookup.Search()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", resp)
		},
	}

	var doLatLong = &cobra.Command{
		Use:   "latlng [lat] [long]",
		Short: "Lookup by latitude and longitude coordinate",
		Long:  "Looks up geocoding information by a given latitude and longitude coordinates.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println(cmd.UseLine())
				return
			}

			var lat, lng float64
			var err error

			lat, err = strconv.ParseFloat(args[0], 10)
			if err != nil {
				fmt.Println(err)
				return
			}

			lng, err = strconv.ParseFloat(args[1], 10)
			if err != nil {
				fmt.Println(err)
				return
			}

			lookup = geocoding.Lookup{
				Location: &geocoding.Point{
					Latitude:  lat,
					Longitude: lng,
				},
			}

			resp, err := lookup.Search()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", resp)
		},
	}

	var app = &cobra.Command{Use: "gogeolook"}
	app.AddCommand(doAddress)
	app.AddCommand(doLatLong)
	app.Execute()
}
