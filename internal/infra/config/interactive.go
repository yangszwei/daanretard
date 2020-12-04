package config

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// InteractiveSetup interactively create config file
func InteractiveSetup() (Config, error) {
	var c Config
	c.DataPath = ask("data folder path", ".daanretard")
	if _, err := os.Stat(c.DataPath); os.IsNotExist(err) {
		var create string
		fmt.Printf("\"%s\" does not exist, create now (Y/n)? ", c.DataPath)
		_, _ = fmt.Scanln(&create)
		if strings.ToLower(create) == "y" {
			if err := os.Mkdir(c.DataPath, 0755); err != nil {
				panic(err)
			}
		}
	}
	if _, err := os.Stat(c.DataPath); os.IsNotExist(err) {
		c.Path = ask("envfile path", ".env")
	} else {
		c.Path = ask("envfile path", path.Join(c.DataPath, ".env"))
	}
	c.Addr = ask("server address", ":8000")
	c.DbDsn = ask("database DSN", "")
	c.FbPageID = ask("facebook page id", "")
	c.FbGraphAppID = ask("facebook graph app id", "")
	c.FbGraphAppSecret = ask("facebook graph app secret", "")
	c.Secret = ask("bcrypt secret", "")
	if _, err := os.Stat(c.Path); !os.IsNotExist(err) {
		var o string
		fmt.Printf("An config file was found at %s, do you want to overwrite it (Y/n)? ", c.Path)
		_, err := fmt.Scanln(&o)
		if err != nil && err.Error() != "unexpected newline" {
			panic(err)
		}
		if strings.ToLower(o) != "y" {
			return c, nil
		}
	}
	if err := c.Save(); err != nil {
		return Config{}, err
	}
	fmt.Printf("Config file saved successfully!")
	return c, nil
}

func ask(name, defaultValue string) string {
	var input, d string
	if defaultValue != "" {
		d = fmt.Sprintf(" (default is %s)", defaultValue)
	}
	fmt.Printf("Enter %s%s: ", name, d)
	if _, err := fmt.Scanln(&input); err != nil && err.Error() != "unexpected newline" {
		fmt.Println("An error occurred, please try again.")
		return ask(name, defaultValue)
	}
	if input == "" {
		if defaultValue != "" {
			return defaultValue
		}
		fmt.Println("Invalid value, try again!")
		return ask(name, defaultValue)
	}
	return input
}
