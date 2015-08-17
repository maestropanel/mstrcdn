package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Nginx struct {
}

func (n *Nginx) Create(domainName string, IpAddr string, Port string, ssl bool, full bool, rootPath string) (m MessageModel) {

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
		m.Message = "Cannot read tmpl file: " + err.Error()
		return m
	}

	template = strings.Replace(template, "##DOMAIN##", domainName, -1)
	template = strings.Replace(template, "##IP##", IpAddr, -1)
	template = strings.Replace(template, "##PORT##", Port, -1)

	vhostConfName := fmt.Sprintf("%s.conf", domainName)
	vhostConfPath := filepath.Join(rootPath, vhostConfName)

	err = n.saveFile(vhostConfPath, template)

	if err != nil {
		m.Message = "Cannot save conf file to " + vhostConfPath + " error: " + err.Error()
		return m
	}

	m.Success = true
	m.Message = "Build Success: " + vhostConfPath

	return m
}

func (n *Nginx) Delete(domainName string, rootPath string) (m MessageModel) {
	m = MessageModel{}
	m.Success = false

	vhostConfName := fmt.Sprintf("%s.conf", domainName)
	vhostConfPath := filepath.Join(rootPath, vhostConfName)

	err := n.deleteFile(vhostConfPath)

	if err != nil {
		m.Message = "File cannot be delete: " + vhostConfPath
		return
	}

	m.Success = true
	m.Message = "Domain deleted: " + domainName

	return m
}

func (n *Nginx) List(rootPath string) (m MessageModelList) {
	m = MessageModelList{}
	m.Success = false
	m.Vhosts = []string{}

	filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		m.Vhosts = append(m.Vhosts, path)
		return nil
	})

	m.Success = true
	m.Message = "Success"

	return m
}

func (n *Nginx) getTemplateFile(fileName string) (text string, err error) {

	currentPath, err := os.Getwd()

	if err != nil {
		return "", err
	}

	path := filepath.Join(currentPath, "tmpl", fileName)
	b, err := ioutil.ReadFile(path)
	alltext := string(b)

	return alltext, err
}

func (n *Nginx) saveFile(filePath string, content string) (err error) {

	fileContent := []byte(content)
	err = ioutil.WriteFile(filePath, fileContent, 0644)

	return err
}

func (n *Nginx) deleteFile(fileName string) (err error) {

	isExists := n.fileExists(fileName)

	if !isExists {
		return err
	}

	err = os.Remove(fileName)

	return err
}

func (n *Nginx) fileExists(fileName string) bool {
	finfo, err := os.Stat(fileName)

	if err != nil {
		return false
	}

	return (finfo.IsDir() == false)
}
