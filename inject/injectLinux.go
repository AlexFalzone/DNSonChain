package inject

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func InjectPKI(name string) {
	exec.Command("certutil", "-d", "sql:$HOME/.pki/nssdb", "-t", "C,,", "-n", name, "-i", "list/certHost.pem")
}

// InjectMozilla finds the default-release profile of Firefox and injects the certificate
func InjectMozilla(name string) {

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

	exec.Command("certutil", "-d", "sql:"+profilePath, "-t", "C,,", "-n", name, "-i", "list/certHost.pem")
}
