package autocfg

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName    = "init_config"
	emTblName = "env_module"
	mcTblName = "module_config"

	queryALLModuleConfig = "SELECT b.name, b.address, b.timeout" +
		" FROM " + emTblName + " AS a, " + mcTblName + " AS b" +
		" WHERE env=? AND module=? AND a.`id`=b.`emid`"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig(env, module string) (mods []DependModule, err error) {
	stmt, err := db.Prepare(queryALLModuleConfig)
	if err != nil {
		log.Println(err)
		return mods, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(env, module)
	if err != nil {
		log.Println(err)
		return mods, err
	}
	defer rows.Close()

	var mod DependModule
	for rows.Next() {
		err = rows.Scan(&mod.Name, &mod.Address, &mod.Timeout)
		if err != nil {
			log.Println(err)
			return mods, ErrScanDependModule
		}

		mods = append(mods, mod)
	}

	log.Printf("get %s module in %s config success: %+v\n", module, env, mods)
	return mods, nil
}
