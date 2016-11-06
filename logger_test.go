package moviesddl

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	/*
	 * Setup
	 */
	config := createConfig(&LogConfig{
		Filepath: "testdata/logger/normal.log",
		LogLevel: "debug",
	})
	os.Remove(config.Log.Filepath)

	/*
	 * Do
	 */
	logger, logfile := SetupLogger(config)

	/*
	 * Assert
	 */
	defer func() {
		ba, _ := ioutil.ReadFile(config.Log.Filepath)
		if !strings.Contains(string(ba), `"msg":"TEST"`) {
			t.Fatalf("期待値がログファイルに書き込まれていません\n【期待値】\n%s\n【ログ】\n%#+v", `"msg":"TES2T"`, string(ba))
		}
	}()

	/*
	 * Teardown
	 */
	defer logfile.Close()
	logger.Debug("TEST")
}

func createConfig(logConfig *LogConfig) *Config {
	config := &Config{
		Server: &ServerConfig{
			Host: "example.com",
		},
		DB: &DBConfig{
			Server:        "127.0.0.1",
			Ports:         []int{8001, 8001, 8002},
			ConnectionMax: 5000,
			Enabled:       true,
		},
		Log: logConfig,
	}
	return config
}
