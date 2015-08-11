package main

import (	
	"io/ioutil"
	"strings"
)

type Nginx struct {
	
}

func (n *Nginx) Create(name string, ssl bool, bool full) (m MessageModel, err error){

	defaultTmpl := "full.cdn.tmpl"	
	
	if(ssl)
		defaultTmpl = "ssl.cdn.tmpl"
		
	if(full)	
		defaultTmp = "full.cdn.tmpl"

	
	m = MessageModel{}
	m.Success = fase	

	template := n.getTemplateFile("full.cdn.tmpl")
	template := strings.Replace("##DOMAIN##", name)
	
	
}

func (n *Nginx) getTemplateFile(fileName string) (text string, err error){
	
	path := "./tmpl"+ fileName
	b, err := ioutil.ReadFile(path)
	alltext = string(b)

	return alltext, err
}

func (n *Nginx) saveFile(filePath string, content string) (err error){
	
	fileContent := []byte(content)
	err = ioutil.WriteFile(filePath, fileContent, 0644)
		
	return err
}
