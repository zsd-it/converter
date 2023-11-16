package converter

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func loadIni(t *Table2Struct) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 数据库配置
	t.dsn = loadDatabase(cfg).GetDsn()
	t.modelConfig = &ModelConfig{
		Converter: loadConverter(cfg),
	}
}

type DatabaseConfig struct {
	Type          string
	User          string
	Password      string
	Host          string
	DbName        string
	Charset       string
	TablePrefix   string
	MaxIdleConns  int
	MaxOpenConns  int
	PrepareStmt   bool
	SingularTable bool
}

func (c DatabaseConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.DbName, c.Charset)
}

func loadDatabase(cfg *ini.File) DatabaseConfig {
	section, err := cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Failed to get section 'database' :%v ", err)
	}

	return DatabaseConfig{
		Type:          section.Key("TYPE").String(),
		User:          section.Key("USER").String(),
		Password:      section.Key("PASSWORD").String(),
		Host:          section.Key("HOST").String(),
		DbName:        section.Key("NAME").String(),
		Charset:       section.Key("CHARSET").MustString("utf8"),
		TablePrefix:   section.Key("TABLE_PREFIX").String(),
		MaxIdleConns:  section.Key("MAX_IDLE_CONNS").MustInt(10),
		MaxOpenConns:  section.Key("MAX_OPEN_CONNS").MustInt(100),
		PrepareStmt:   section.Key("PREPARE_STMT").MustBool(false),
		SingularTable: section.Key("SINGULAR_TABLE").MustBool(false),
	}
}

type Converter struct {
	DaoPath string // 存储路径
	TagKey  string //
}

func loadConverter(cfg *ini.File) Converter {
	section, err := cfg.GetSection("converter")
	if err != nil {
		log.Fatalf("Failed to get section 'converter' :%v ", err)
	}

	return Converter{
		DaoPath: section.Key("DaoPath").String(),
		TagKey:  section.Key("TagKey").MustString("orm"),
	}
}
