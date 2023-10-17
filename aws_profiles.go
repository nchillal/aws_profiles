package aws_profiles

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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

func ListAWSRegions(awsProfile string) ([]string, error) {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(awsProfile),
	)
	if err != nil {
		return []string{}, nil
	}

	// Create an EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Call DescribeRegions to get a list of regions
	resp, err := ec2Client.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{})
	if err != nil {
		return []string{}, nil
	}

	// Get list of regions
	regions := make([]string, 0)
	for _, region := range resp.Regions {
		regions = append(regions, *region.RegionName)
	}
	return regions, nil
}
