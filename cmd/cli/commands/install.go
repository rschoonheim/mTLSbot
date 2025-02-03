package commands

import (
	"errors"
	"mTLS/cmd/cli/state"
	"os"
)

type InstallCliCmd struct {
	// Path - path to store data of the root server.
	Path string `arg:"-p,--path, required" help:"Path to store data of the root server."`

	// RootServerAddress - address of where the root server is running.
	RootServerAddress string `arg:"--root_address, required" help:"Address of where the root server is running."`
}

// InstallCli - installs the root server.
func InstallCli(args *InstallCliCmd) {

	// Validate the arguments
	//
	if args.Path == "" {
		panic("--path is required")
	}

	// Path validation
	//
	if err := validatePath(args.Path); err != nil {
		switch err.Error() {
		case "PATH_CANNOT_BE_EMPTY":
			panic("Installation path cannot be empty")
		case "PATH_NOT_FOUND":
			panic("The installation path was not found on the system")
		case "PATH_NOT_DIRECTORY":
			panic("The installation path was found but it is not a directory")
		case "DIRECTORY_NOT_EMPTY":
			panic("The given installation path is not empty")
		default:
			panic("An unknown error occurred")
		}
	}

	// Root Server validation
	//
	if err := validateRootServer(args.RootServerAddress); err != nil {
		switch err.Error() {
		case "ADDRESS_CANNOT_BE_EMPTY":
			panic("Root server address cannot be empty")
		default:
			panic("An unknown error occurred")
		}
	}

	// Create state object
	//
	s := state.State{
		RootServer: state.RootServer{
			Address: args.RootServerAddress,
		},
	}

	// Write the state to the file system.
	//
	stateString := s.ToJson()
	err := os.WriteFile(args.Path+"/cli.json", []byte(stateString), 0644)
	if err != nil {
		panic("Failed to write state to the file system")
	}

	println("Root server installed successfully")
}

// validateRootServer - ensures that the root server is running.
func validateRootServer(address string) error {

	// The address cannot be empty.
	//
	if address == "" {
		return errors.New("ADDRESS_CANNOT_BE_EMPTY")
	}

	return nil
}

// validatePath - ensures that the path can be used to install the CLI.
func validatePath(path string) error {

	// The path cannot be empty.
	//
	if path == "" {
		return errors.New("PATH_CANNOT_BE_EMPTY")
	}

	// The path must exist.
	//
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("PATH_NOT_FOUND")
	}

	// The path must be a directory.
	//
	if file, _ := os.Stat(path); !file.IsDir() {
		return errors.New("PATH_NOT_DIRECTORY")
	}

	// Directory must be empty.
	//
	if files, _ := os.ReadDir(path); len(files) > 0 {
		return errors.New("DIRECTORY_NOT_EMPTY")
	}

	return nil
}
