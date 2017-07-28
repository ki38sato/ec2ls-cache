package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func ec2list(profile string, region string, updateCache bool, cachename string, filters []string) ([]Ec2Info, error) {
	if updateCache {
		ec2s, err := findEc2s(profile, region, filters)
		if err != nil {
			return nil, err
		}

		writeCache(ec2s, cachename)
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

func readFromCache(cachename string) ([]Ec2Info, error) {
	expath, err := expandPath(cachename)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadFile(expath)
	if err != nil {
		return nil, err
	}

	var ec2s []Ec2Info
	if err := json.Unmarshal(jsonBytes, &ec2s); err != nil {
		return nil, err
	}
	return ec2s, nil
}

func writeCache(ec2s []Ec2Info, cachename string) error {
	jsonBytes, err := json.Marshal(ec2s)
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
