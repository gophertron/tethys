package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

type Page struct {
	Name    string
	Title   string
	Content []byte
}

func (p *Page) Path() string {
	fileName := fmt.Sprintf("%s.md", p.Name)
	return path.Join(GetConfig().RepoRoot, fileName)
}

func (p *Page) Save() error {
	return ioutil.WriteFile(p.Path(), p.Content, 0600)
}

func (p *Page) Update(title string, content string) error {
	prevPage := p.Path()

	p.Title = title
	p.Name = nameFromTitle(title)
	p.Content = []byte(content)

	err := p.Save()

	if err != nil {
		return err
	}

	if prevPage != p.Path() {
		os.Remove(prevPage)
	}

	return nil
}

func (p *Page) Markup() template.HTML {
	return template.HTML(string(Markdown(p.Content)))
}

func (p *Page) Markdown() string {
	return string(p.Content)
}

func (p *Page) Url() string {
	return fmt.Sprintf("/wiki/%s", p.Name)
}

func LoadPage(name string) (*Page, error) {
	fileName := fmt.Sprintf("%s.md", name)
	filePath := path.Join(GetConfig().RepoRoot, fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	title := titleFromName(name)
	return &Page{Name: name, Title: title, Content: content}, nil
}
