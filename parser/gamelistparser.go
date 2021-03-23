package parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//GameListItem 数据项
type GameListItem struct {
	XMLName     xml.Name `xml:"game"`
	Path        string   `xml:"path"`
	Name        string   `xml:"name"`
	Image       string   `xml:"image,omitempty"`
	Video       string   `xml:"video,omitempty"`
	PlayCount   int32    `xml:"playcount,omitempty"`
	LastPlayed  string   `xml:"lastplayed,omitempty"`
	Lang        string   `xml:"lang,omitempty"`
	Desc        string   `xml:"desc,omitempty"`
	Rating      float32  `xml:"rating,omitempty"`
	ReleaseDate string   `xml:"releasedate,omitempty"`
	Developer   string   `xml:"developer,omitempty"`
	Publisher   string   `xml:"publisher,omitempty"`
	Genre       string   `xml:"genre,omitempty"`
	Players     string   `xml:"players,omitempty"`
	Marquee     string   `xml:"marquee,omitempty"`
}

//GameList 游戏列表
type GameListInfo struct {
	XMLName xml.Name       `xml:"gameList"`
	Games   []GameListItem `xml:"game"`
}

//GameListParser parser 解析gamelist.xml
type GameListParser struct {
	file  string
	gl    GameListInfo
	store map[string]*GameListItem
}

//NewGameListParser 创建新的parser
func NewGameListParser(f string) *GameListParser {
	g := &GameListParser{file: f, store: make(map[string]*GameListItem)}
	return g
}

//Parse parse
func (g *GameListParser) Parse() error {
	data, err := ioutil.ReadFile(g.file)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, &g.gl)
	if err != nil {
		return err
	}
	for index := range g.gl.Games {
		item := &g.gl.Games[index]
		if len(item.Name) == 0 || len(item.Path) == 0 {
			continue
		}
		g.store[item.Name] = item
	}
	return nil
}

//GetAll 读取整个map
func (g *GameListParser) GetAll() map[string]*GameListItem {
	return g.store
}

//Get 获取对应项
func (g *GameListParser) Get(name string) (*GameListItem, bool) {
	if item, ok := g.store[name]; ok {
		return item, true
	}
	return nil, false
}

//Remove 删除对应的项
func (g *GameListParser) Remove(name string) {
	delete(g.store, name)
}

//Save 将数据进行保存
func (g *GameListParser) Save() error {
	//save old
	err := os.Rename(g.file, g.file+".bak")
	if err != nil {
		return err
	}
	g.gl.Games = nil
	for _, item := range g.store {
		g.gl.Games = append(g.gl.Games, *item)
	}
	data, err := xml.MarshalIndent(&g.gl, " ", "  ")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(g.file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(xml.Header)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
