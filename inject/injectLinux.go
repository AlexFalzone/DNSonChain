package inject

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"test/util"
)

func InjectPKI(name string, cert string) error {

	// Get the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %v", err)
	}
	nssdbPath := filepath.Join(homeDir, ".pki", "nssdb")

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+nssdbPath, "-t", "C,,", "-n", name, "-i", cert)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}
	return nil

}

// InjectMozilla finds the default-release profile of Firefox and injects the certificate
func InjectMozilla(name string, cert string) error {

	// Get the current user
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get the path of the Firefox profile
	firefoxProfilesPath := filepath.Join(usr.HomeDir, ".mozilla", "firefox")
	profiles, err := os.ReadDir(firefoxProfilesPath)
	if err != nil {
		log.Fatal(err)
	}

	var profilePath string
	// Find the default-release profile
	for _, profile := range profiles {
		if profile.IsDir() && strings.Contains(profile.Name(), ".default-release") {
			profilePath = filepath.Join(firefoxProfilesPath, profile.Name())
			break
		}
	}

	if profilePath == "" {
		log.Fatal("Impossibile trovare il profilo Firefox")
	}

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+profilePath, "-t", "C,,", "-n", name, "-i", cert)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}

	return nil
}

func InjcetLinux(certString string, hostname string) error {
	err := util.CreateandWriteTemp(certString)
	if err != nil {
		fmt.Println(err)
	}

	errPKI := InjectPKI(hostname, certString)
	if err != nil {
		fmt.Println(err)
	}
	errMozilla := InjectMozilla(hostname, certString)
	if err != nil {
		fmt.Println(err)
	}

	util.RemoveTempFile(errPKI, errMozilla, err)

	return nil
}
