/* -*- indent-tabs-mode:nil; coding: utf-8 -*-
 * Copyleft (C) 2015
 * "Tag bao" known as "wackonline" <bb.qnyd@gmail.com>
 * Goway is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License and GNU
 * Lesser General Public License published by the Free Software
 * Foundation, either version 3 of the License, or (at your option)
 * any later version.
 * Goway is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License and GNU Lesser General Public License
 * for more details.
 * You should have received a copy of the GNU General Public License
 * and GNU Lesser General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */
package goway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	CONFIGFILE = "/conf/config.json"
	HTTPSERVER = "0.0.0.0"
	PORT       = "8080"
	ENV        = "development"
	DEBUG      = true
	LOGGER     = "E_ALL"
        STATICPATH ="/public"
        VERSION = "0.0.1"
)

type Config interface {
	Getconf(string, string) interface{}
	Get(string) interface{}
	Set(string, interface{}) Config
}

type config struct {
	configs map[string]interface{}
}

// By key to get the value of the configuration file
func (c *config) Get(key string) interface{} {
	return c.configs[key]
}

// By key to set the value
func (c *config) Set(key string, value interface{}) Config {
	c.configs[key] = value
	return c
}

func (c *config) isFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// reading config file.
func (c *config) readFile(path string) ([]byte, error) {
	if c.isFileExist(path) {
		return ioutil.ReadFile(path)
	}
	return nil, fmt.Errorf("Config file is not exist.")
}

// Access to data in a configuration file
func (c *config) Getconf(path string, key string) interface{} {
	data, err := c.readFile(path)
	results := make(map[string]interface{})
	if err == nil {
		if err := json.Unmarshal(data, results); err != nil {
			return results[key]
		}
	}
	return nil
}

// setting default config infomation
func (c *config) defaultConfig() {
	c.configs["httpServer"] = HTTPSERVER
	c.configs["serverPort"] = PORT
	c.configs["staticPath"] = STATICPATH
	c.configs["logger"] = E_ALL
	c.configs["debug"] = DEBUG
	c.configs["version"] = VERSION
}

// parase config,setting values
func (c *config) paraseConfig() {
	root, oerr := os.Getwd()
	if oerr != nil {
		panic("get root path fail...")
	}
	data, err := c.readFile(path.Join(root, CONFIGFILE))
	if err == nil {
		err := json.Unmarshal(data, c.configs)
		if err != nil {
			c.defaultConfig()
		}
	} else {
		c.defaultConfig()
	}

}

// Initialize the configuration file
// if read config file failer,used default config
func Init() Config {
	configs := &config{}
	configs.paraseConfig()
	return configs
}
