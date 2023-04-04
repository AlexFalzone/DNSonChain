package inject

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func InjectPKI(name string, certPath string) error {

	// Get the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %v", err)
	}
	nssdbPath := filepath.Join(homeDir, ".pki", "nssdb")

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+nssdbPath, "-t", "C,,", "-n", name, "-i", certPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}
	return nil

}

// InjectMozilla finds the default-release profile of Firefox and injects the certificate
func InjectMozilla(name string, certPath string) error {

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

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+profilePath, "-t", "C,,", "-n", name, "-i", certPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}

	return nil
}

func SaveCertToFile(certContent string) (string, error) {
	certDir := "tmp"
	err := os.MkdirAll(certDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create certReq directory: %v", err)
	}

	certPath := filepath.Join(certDir, "tmp.pem")
	err = os.WriteFile(certPath, []byte(certContent), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write certificate to file: %v", err)
	}

	return certPath, nil
}

func InjcetLinux(hostname string, certPath string) error {

	errPKI := InjectPKI(hostname, certPath)
	if errPKI != nil {
		fmt.Println(errPKI)
	}

	errMozilla := InjectMozilla(hostname, certPath)
	if errMozilla != nil {
		fmt.Println(errMozilla)
	}

	return nil
}
