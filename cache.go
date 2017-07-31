package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func ec2list(profile string, region string, updateCache bool, cachename string, filters []string, columns string) (map[string]interface{}, error) {
	if updateCache {
		cacheInfo, err := findEc2s(profile, region, filters, columns)
		if err != nil {
			return nil, err
		}

		writeCache(cacheInfo, cachename)
	}

	return readFromCache(cachename)
}

func expandPath(cachename string) (string, error) {
	expath, err := homedir.Expand(cacheBasePath + cachename)
	if err != nil {
		return "", err
	}
	return expath, nil
}

func readFromCache(cachename string) (map[string]interface{}, error) {
	expath, err := expandPath(cachename)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadFile(expath)
	if err != nil {
		return nil, err
	}

	var cacheInfo map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &cacheInfo); err != nil {
		return nil, err
	}
	return cacheInfo, nil
}

func writeCache(cacheInfo map[string]interface{}, cachename string) error {
	jsonBytes, err := json.Marshal(cacheInfo)
	if err != nil {
		return err
	}

	if err = prepareCache(); err != nil {
		return err
	}

	expath, err := expandPath(cachename)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(expath, jsonBytes, 0644); err != nil {
		return err
	}
	return nil
}

func prepareCache() error {
	// TODO: exits dir ?
	expath, err := homedir.Expand(cacheBasePath)
	if err != nil {
		return err
	}
	if _, err := os.Stat(expath); err != nil {
		if err := os.MkdirAll(expath, 0777); err != nil {
			return err
		}
	}
	return nil
}
