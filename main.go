package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mickelsonm/go-helpers/geocoding"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("gogeolook", "0.0.2")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"address": func() (cli.Command, error) {
			return new(AddressCommand), nil
		},
		"latlng": func() (cli.Command, error) {
			return new(LatLongCommand), nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

type AddressCommand struct {
	cli.Command
}

func (cmd AddressCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println("usage: ", cmd.Help())
		return 0
	}
	lookup := geocoding.Lookup{
		Address: strings.Join(args, " "),
	}

	resp, err := lookup.Search()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	js, err := json.MarshalIndent(resp, " ", " ")
	if err != nil {
		fmt.Println(fmt.Errorf("error marshaling json: %s", err))
		return 0
	}

	fmt.Printf("%s\n", string(js))
	return 1
}

func (cmd AddressCommand) Help() string {
	return "address <location to lookup>"
}
func (cmd AddressCommand) Synopsis() string {
	return "address <location to lookup>"
}

type LatLongCommand struct {
	cli.Command
}

func (cmd LatLongCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Println("usage: ", cmd.Help())
		return 0
	}

	var lat, lng float64
	var err error

	lat, err = strconv.ParseFloat(args[0], 10)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	lng, err = strconv.ParseFloat(args[1], 10)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	lookup := geocoding.Lookup{
		Location: &geocoding.Point{
			Latitude:  lat,
			Longitude: lng,
		},
	}

	resp, err := lookup.Search()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	js, err := json.MarshalIndent(resp, " ", " ")
	if err != nil {
		fmt.Println(fmt.Errorf("error marshaling json: %s", err))
		return 0
	}

	fmt.Printf("%s\n", string(js))
	return 1
}
func (cmd LatLongCommand) Help() string {
	return "latlng <latitude> <longitude>"
}
func (cmd LatLongCommand) Synopsis() string {
	return "latlng <lat> <lng>"
}
