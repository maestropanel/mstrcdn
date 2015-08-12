package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Nginx struct {
}

func (n *Nginx) Create(domainName string, ssl bool, full bool, rootPath string) (m MessageModel) {

	defaultTmpl := "full.cdn.tmpl"

	if ssl {
		defaultTmpl = "ssl.cdn.tmpl"
	}

	if !full {
		defaultTmpl = "split.cdn.tmpl"
	}

	m = MessageModel{}
	m.Success = false

	template, err := n.getTemplateFile(defaultTmpl)

	if err != nil {
		m.Message = "Canot read tmpl file. " + err.Error()
		return m
	}

	template = strings.Replace(template, "##DOMAIN##", domainName, -1)

	vhostConfName := fmt.Sprintf("%s.conf", domainName)
	vhostConfPath := fmt.Sprintf("%s/%s", rootPath, vhostConfName)

	err = n.saveFile(vhostConfPath, template)

	if err != nil {
		m.Message = "Cannot save conf file to " + vhostConfPath + " error: " + err.Error()
		return m
	}

	m.Success = true
	m.Message = "Build Success: " + vhostConfPath

	return m
}

func (n *Nginx) getTemplateFile(fileName string) (text string, err error) {

	path := "./tmpl/" + fileName
	b, err := ioutil.ReadFile(path)
	alltext := string(b)

	return alltext, err
}

func (n *Nginx) saveFile(filePath string, content string) (err error) {

	fileContent := []byte(content)
	err = ioutil.WriteFile(filePath, fileContent, 0644)

	return err
}
