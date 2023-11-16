package tmp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	ModelName = "{@modelName}"
)

type ModelTmp struct {
	file     *os.File
	name     string // 模块名称
	savePath string // 存储路径
	daoStr   string // dao 路径
}

func (tmp *ModelTmp) Do() {
	tmp.buildModelSavePath()
	var err error
	tmp.file, err = os.Create(tmp.savePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	n, err := tmp.file.WriteString(tmp.structContent())
	if err != nil {
		fmt.Println(n, err.Error())
	}

	tmp.save()
}

func NewModelTmp(name string, daoStr string) *ModelTmp {
	return &ModelTmp{
		name:   name,
		daoStr: daoStr,
	}
}

func (tmp *ModelTmp) buildNewModel() string {
	newStr := `
func New{@modelName}Model(db *gorm.DB) *{@modelName}Model {
	return &{@modelName}Model{
		db: db,
	}
}
`
	return strings.Replace(newStr, ModelName, tmp.name, -1)
}

func (tmp *ModelTmp) save() {
	cmd := exec.Command("gofmt", "-w", tmp.savePath)
	_ = cmd.Run()

	_ = tmp.file.Close()
}

func (tmp *ModelTmp) buildModelSavePath() {
	tmp.savePath = fmt.Sprintf("%s/%sModel.go", tmp.daoStr, tmp.name)
}

func (tmp *ModelTmp) structContent() string {
	str := tmp.buildNewModel()

	return str
}
