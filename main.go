package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func main() {
	bumpType := "patch"
	verStr := "0.0.0"
	if len(os.Args) > 1 {
		verStr = os.Args[1]
		if len(os.Args) > 2 && len(os.Args[2]) > 1 {
			bumpType = os.Args[2]
		}
	} else {
		log.Println("invalid options provided, must provide VERSION patch|minor|major")
		os.Exit(1)
	}

	var version = Version{}

	verParts := strings.Split(verStr, ".")
	if len(verParts) == 1 {
		// version 1
		major, majErr := strconv.Atoi(verParts[0])
		if majErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verStr[0] to int: %v", majErr))
			return
		}
		version = Version{Major: major}
	} else if len(verParts) == 2 {
		// version 1.1
		major, majErr := strconv.Atoi(verParts[0])
		if majErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verParts[0] to int: %v", majErr))
			return
		}
		minor, minErr := strconv.Atoi(verParts[1])
		if minErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verParts[1] to int: %v", minErr))
			return
		}
		version = Version{Major: major, Minor: minor}
	} else if len(verParts) == 3 {
		major, majErr := strconv.Atoi(verParts[0])
		if majErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verParts[0] to int: %v", majErr))
			return
		}
		minor, majErr := strconv.Atoi(verParts[1])
		if majErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verParts[1] to int: %v", majErr))
			return
		}
		patch, patErr := strconv.Atoi(verParts[2])
		if patErr != nil {
			fmt.Println(fmt.Errorf("failed to convert verParts[2] to int: %v", patErr))
			return
		}
		version = Version{Major: major, Minor: minor, Patch: patch}
	} else {
		// unsupported version
		fmt.Println(fmt.Errorf("unable to parse a 4 dot version number, only Major.Minor.Patch is allowed of ints allowed: %v", errors.New("unsupported version parse")))
		return
	}

	if bumpType == "major" {
		version.Major += 1
	} else if bumpType == "minor" {
		version.Minor += 1
	} else {
		version.Patch += 1
	}

	fmt.Print(version)
}
