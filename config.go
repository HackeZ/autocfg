package autocfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	configPath = "./config/init.json"
)

var cfg Config

func init() {
	bs, err := readCfg()
	if err != nil {
		os.Exit(1)
	}

	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	if "" == cfg.ENV || "" == cfg.Module {
		log.Println(ErrNoSettingENVORMod)
		os.Exit(3)
	}
}

// readCfg return content in config.
func readCfg() (bs []byte, e error) {
	if _, e = os.Stat(configPath); os.IsNotExist(e) {
		log.Println(e)
		return bs, e
	}

	f, e := os.Open(configPath)
	if e != nil {
		log.Println(e)
		return bs, e
	}

	bs, e = ioutil.ReadAll(f)
	if e != nil {
		log.Println(e)
		return bs, e
	}
	return bs, e
}

func Get() (mods []DependModule, err error) {
	mods, err = getConfig(cfg.ENV, cfg.Module)
	if err != nil {
		log.Println(err)
		return mods, err
	}
	return mods, nil
}
