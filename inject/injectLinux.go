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

// InjectPKI injects a certificate into the NSS database.
//
// Parameters:
//   - name: string representing the name for the certificate.
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func InjectPKI(name string, certPath string) error {

	// Get the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %v", err)
	}
	nssdbPath := filepath.Join(homeDir, ".pki", "nssdb")

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+nssdbPath, "-t", "Pu,,", "-n", name, "-i", certPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}
	return nil

}

// InjectMozilla finds the default-release profile of Firefox and injects the certificate.
//
// Parameters:
//   - name: string representing the name for the certificate.
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
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

	cmd := exec.Command("certutil", "-A", "-d", "sql:"+profilePath, "-t", "Cu,,", "-n", name, "-i", certPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v\nOutput: %s", err, output)
	}

	return nil
}

// InjcetLinux injects a certificate into the NSS database and Firefox profile on Linux.
//
// Parameters:
//   - hostname: string representing the hostname.
//   - certPath: string representing the path to the certificate file.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
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

// fixCertificate adds a new line every 64 characters in the given string.
// This is needed because the certificate is returned in a single line.
//
// Parameters:
//   - cert: string representing the certificate content.
//
// Returns:
//   - string: The modified certificate content.
func fixCertificate(cert string) string {
	header := "-----BEGIN CERTIFICATE-----"
	footer := "-----END CERTIFICATE-----"

	// Trim the header and footer
	cert = strings.TrimPrefix(cert, header)
	cert = strings.TrimSuffix(cert, footer)

	fixedCert := strings.ReplaceAll(cert, " ", "\n")

	fixedCert = header + fixedCert + footer

	return fixedCert
}

// SaveCertToFile saves a certificate to a file.
//
// Parameters:
//   - certContent: string representing the certificate content.
//
// Returns:
//   - string: The path to the saved certificate file.
//   - error: An error if the operation fails, nil otherwise.
func SaveCertToFile(certContent string) (string, error) {
	certDir := "tmp"
	err := os.MkdirAll(certDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create certReq directory: %v", err)
	}

	certPath := filepath.Join(certDir, "tmp.pem")

	// Before writing the certificate to file, add a new line every 64 characters
	certContent = fixCertificate(certContent)

	err = os.WriteFile(certPath, []byte(certContent), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write certificate to file: %v", err)
	}

	return certPath, nil
}
