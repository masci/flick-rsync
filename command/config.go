package command

import (
	"encoding/json"
	"os"
	"os/user"
	"path"
)

const (
	configFileName = ".flick-rsync.cfg"
)

type ConfigFile struct {
	// preserve the same naming of the command line flags inside the config file
	ApiKey           string `json:"api_key"`
	ApiSecret        string `json:"api_secret"`
	OAuthToken       string `json:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret"`
}

// compute the absolute path of the configuration file
func getConfigFilePath() string {
	// get current system user
	u, err := user.Current()
	if err != nil {
		return ""
	}

	return path.Join(u.HomeDir, configFileName)
}

// search for a .flick-rsync.cfg file in user's home directory
// if present, unmarshal file contents and return a ConfigFile instance
func loadConfigFile(filePath string) (*ConfigFile, error) {
	config := ConfigFile{}

	cfg_file, err := os.Open(filePath)
	defer cfg_file.Close()
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(cfg_file)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
