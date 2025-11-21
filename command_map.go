package main

import (
	"fmt"
	"errors"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
} // End commandMapf() func

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you have not explored any of the map yet to look at previous locations")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
} // End commandMapb() func