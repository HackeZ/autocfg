package autocfg

import "errors"

// ErrNoExistModule when no module setting in config
var ErrNoExistModule = errors.New("can not found setting module in config")

// ErrNoSettingENVORMod when the env or module not set,
// you must set the 'env' 'module' in config file and
// it is the necessary conditions for searching config
var ErrNoSettingENVORMod = errors.New("please setting 'env' and 'module' in config file")

// ErrScanDependModule
var ErrScanDependModule = errors.New("scan value to depend module failed:")

// ErrNoSettingCfg can not find this computer config in db
var ErrNoSettingCfg = errors.New("can not find any config in center")
