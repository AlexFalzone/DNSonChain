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

func SaveCertToFileWindows(certContent string) (string, error) {
	certDir := "tmp"
	err := os.MkdirAll(certDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create certReq directory: %v", err)
	}

	certPath := filepath.Join(certDir, "tmpWindows.crt")

	// Before writing the certificate to file, add a new line every 64 characters
	certContent = fixCertificate(certContent)

	err = os.WriteFile(certPath, []byte(certContent), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write certificate to file: %v", err)
	}

	return certPath, nil
}

// InjectTrustStore injects a certificate into the Windows trust store.
//
// Parameters:
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func InjectTrustStore(certPath string) error {
	cmd := exec.Command("certutil.exe", "-user", "-addstore", "Root", certPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}
	return nil
}

// InjectMozillaWindows finds the default-release profile of Firefox on Windows and injects the certificate.
//
// Parameters:
//   - name: string representing the name for the certificate.
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func InjectMozillaWindows(name string, certPath string) error {

	// Get the current user
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get the path of the Firefox profile
	firefoxProfilesPath := filepath.Join(usr.HomeDir, "AppData", "Roaming", "Mozilla", "Firefox")
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

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+profilePath, "-t", "Cu,,", "-n", name, "-i", certPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}

	return nil
}

// InjectWindows injects a certificate into the Windows trust store and Firefox profile.
//
// Parameters:
//   - hostname: string representing the hostname.
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func InjectWindows(hostname string, certPath string) error {
	errTrustStore := InjectTrustStore(certPath)
	if errTrustStore != nil {
		fmt.Println(errTrustStore)
	}

	errMozilla := InjectMozillaWindows(hostname, certPath)
	if errMozilla != nil {
		fmt.Println(errMozilla)
	}

	return nil
}
