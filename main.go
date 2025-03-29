package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Preset struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type Config struct {
	Presets map[string]Preset `yaml:"presets"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: git-preset <preset-name>")
	}

	presetName := os.Args[1]
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	preset, exists := config.Presets[presetName]
	if !exists {
		log.Fatalf("Preset '%s' not found. Available presets: %v",
			presetName, getPresetNames(config))
	}

	fmt.Printf("Loaded preset %s\n", presetName)
	setGitConfig("user.name", preset.Name)
	setGitConfig("user.email", preset.Email)
}

func loadConfig() (*Config, error) {
	dotConfigPath, err := getUserConfigDirectory()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(filepath.Dir(dotConfigPath), "git-preset", "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid config format: %v", err)
	}

	return &config, nil
}

func setGitConfig(key, value string) {
	cmd := exec.Command("git", "config", "--local", key, value)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to set %s: %v", key, err)
	}
	fmt.Printf("Set %s to %s\n", key, value)
}

func getUserConfigDirectory() (string, error) {
	currentUser, err := user.Current()

	if err != nil {
		return "", err
	}

	configDir := filepath.Join(currentUser.HomeDir, ".config", "git-preset")

	return configDir, nil
}

func getPresetNames(config *Config) []string {
	names := make([]string, 0, len(config.Presets))
	for name := range config.Presets {
		names = append(names, name)
	}
	return names
}
