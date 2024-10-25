package xml

import (
	"encoding/xml"
	"os"

	"github.com/lights-T/goutils"
	"github.com/lights-T/goutils/files"
)

func Create(xmlName string, v any) error {
	// 重新序列化XML数据
	updatedData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return goutils.Errorf("XML数据序列化失败: %s", err.Error())
	}

	// 写入修改后的XML数据
	if err = files.WriteFileByOptional(xmlName, os.O_CREATE, xml.Header+string(updatedData)); err != nil {
		return goutils.Errorf("写入XML文件失败: %s", err.Error())
	}
	return nil
}

func Update(xmlName string, v any) error {
	// 读取XML文件
	_, err := files.ReadFile(xmlName)
	if err != nil {
		return goutils.Errorf("读取XML文件失败: %s", err.Error())
	}
	if err = files.Remove(xmlName); err != nil {
		return goutils.Errorf("删除XML文件失败: %s", err.Error())
	}
	// 重新序列化XML数据
	updatedData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return goutils.Errorf("XML数据序列化失败: %s", err.Error())
	}

	// 写入修改后的XML数据
	if err = files.WriteFileByOptional(xmlName, os.O_CREATE, xml.Header+string(updatedData)); err != nil {
		return goutils.Errorf("写入XML文件失败: %s", err.Error())
	}
	return nil
}
