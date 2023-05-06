package aws_profiles

import (
	"os"
	"strings"

	"github.com/go-ini/ini"
)

func ListAWSProfiles() ([]string, error) {
	configFile := os.ExpandEnv("${HOME}/.aws/config")

	// Read the contents of the config file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	// Parse the config file
	cfg, err := ini.Load(data)
	if err != nil {
		return nil, err
	}

	// Get the list of sections (which are the profiles)
	sections := cfg.SectionStrings()
	profiles := make([]string, 0)

	// Extract the profile names
	for _, section := range sections {
		if strings.HasPrefix(section, "profile ") {
			profile := strings.TrimPrefix(section, "profile ")
			profiles = append(profiles, profile)
		}
	}

	return profiles, nil
}
