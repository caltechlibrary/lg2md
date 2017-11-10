// lg2md is a Go lang package exploring processing the XML export document from LibGuides
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package lg2md

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	Version = `v0.0.0-prototype`
)

const (
	DateFmt = "2006-01-02 15:04:05"
)

type LibGuides struct {
	XMLName  xml.Name   `json:"-"`
	Customer *Customer  `json:"customer" xml:"customer"`
	Site     *Site      `json:"site" xml:"site"`
	Accounts []*Account `json:"accounts" xml:"accounts>account"`
	Groups   []*Group   `json:"groups" xml:"groups>group"`
	Subjects []*Subject `json:"subjects" xml:"subjects>subject"`
	Tags     []*Tag     `json:"tags" xml:"tags>tag"`
	//FIXME: Vendors probably needs to be defined as a struct type...
	Vendors string   `json:"vendors" xml:"vendors"`
	Guides  []*Guide `json:"guides" xml:"guides>guide"`
}

type Customer struct {
	XMLName  xml.Name `json:"-"`
	ID       int      `json:"id" xml:"id"`
	Type     string   `json:"type" xml:"type"`
	Name     string   `json:"name" xml:"name"`
	URL      string   `json:"url" xml:"url"`
	City     string   `json:"city" xml:"city"`
	State    string   `json:"state" xml:"state"`
	Country  string   `json:"country" xml:"country"`
	TimeZone string   `json:"time_zone" xml:"time_zone"`
	Created  string   `json:"created" xml:"created"`
	Updated  string   `json:"updated" xml:"updated"`
}

type Site struct {
	XMLName xml.Name `json:"-"`
	ID      int      `json:"id" xml:"id"`
	Type    string   `json:"type" xml:"type"`
	Name    string   `json:"name" xml:"name"`
	Domain  string   `json:"domain" xml:"domain"`
	Admin   string   `json:"admin" xml:"admin"`
	Created string   `json:"created" xml:"created"`
	Updated string   `json:"updated" xml:"updated"`
}

type Account struct {
	XMLName   xml.Name `json:"-"`
	ID        int      `json:"id" xml:"id"`
	EMail     string   `json:"email" xml:"email"`
	FirstName string   `json:"first_name" xml:"first_name"`
	LastName  string   `json:"last_name" xml:"last_name"`
	Nickname  string   `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Signature string   `json:"signature,omitempty" xml:"signature,omitempty"`
	Image     string   `json:"image,omitempty" xml:"image,omitempty"`
	Address   string   `json:"address,omitempty" xml:"address,omitempty"`
	Phone     string   `json:"phone,omitempty" xml:"phone,omitempty"`
	Skype     string   `json:"skype,omitempty" xml:"skype,omitempty"`
	Website   string   `json:"website,omitempty" xml:"website,omitempty"`
	Created   string   `json:"created,omitempty" xml:"created,omitempty"`
	Updated   string   `json:"updated,omitempty" xml:"updated,omitempty"`
}

type Owner struct {
	XMLName   xml.Name `json:"-"`
	ID        int      `json:"id" xml:"id"`
	EMail     string   `json:"email" xml:"email"`
	FirstName string   `json:"first_name" xml:"first_name"`
	LastName  string   `json:"last_name" xml:"last_name"`
}

type Group struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Type        string   `json:"type" xml:"type"`
	Description string   `json:"description,omitempty" xml:"description,innerXML,omitempty"`
	Name        string   `json:"name" xml:"name"`
	Created     string   `json:"created,omitempty" xml:"created,omitempty"`
	Updated     string   `json:"updated,omitempty" xml:"updated,omitempty"`
}

type Subject struct {
	XMLName xml.Name `json:"-"`
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	URL     string   `json:"url" xml:"url"`
}

type Tag struct {
	XMLName xml.Name `json:"-"`
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
}

type Guide struct {
	XMLName     xml.Name   `json:"-"`
	ID          int        `json:"id" xml:"id"`
	Type        string     `json:"type" xml:"type"`
	Name        string     `json:"name" xml:"name"`
	Description string     `json:"description" xml:"description,innerXML"`
	URL         string     `json:"url" xml:"url"`
	Owner       *Owner     `json:"owner" xml:"owner"`
	Group       *Group     `json:"group" xml:"group"`
	Redirect    string     `json:"redirect,omitempty" xml:"redirect,omitempty"`
	Status      string     `json:"status" xml:"status"`
	Published   string     `json:"published" xml:"published"`
	Subjects    []*Subject `json:"subject,omitempty" xml:"subjects>subject,omitepty"`
	Tags        []*Tag     `json:"tags,omitempty" xml:"tags>tag,omitempty"`
	Created     string     `json:"created" xml:"created,omitepmty"`
	Updated     string     `json:"updated" xml:"updated,omitempty"`
	Pages       []*Page    `json:"pages,omitempty" xml:"pages>page,omitempty"`
}

type Page struct {
	XMLName      xml.Name `json:"-"`
	ID           int      `json:"id" xml:"id"`
	Name         string   `json:"name" xml:"name"`
	Description  string   `json:"description" xml:"description,innerXML"`
	URL          string   `json:"url" xml:"url"`
	Redirect     string   `json:"redirect,omitempty" xml:"redirect,omitempty"`
	SourcePageID int      `json:"source_page_id" xml:"source_page_id"`
	ParentPageID int      `json:"parent_page_id" xml:"parent_page_id"`
	Position     int      `json:"position" xml:"position"`
	Hidden       int      `json:"hidden" xml:"hidden"`
	Created      string   `json:"created" xml:"created"`
	Updated      string   `json:"updated" xml:"updated"`
	Boxes        []*Box   `json:"boxes" xml:"boxes>box"`
}

type Box struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Name        string   `json:"name" xml:"name"`
	Type        string   `json:"type" xml:"type"`
	Description string   `json:"description" xml:"description,innerXML"`
	URL         string   `json:"url" xml:"url"`
	Owner       *Owner   `json:"owner" xml:"owner"`
	MapID       int      `json:"map_id" xml:"map_id"`
	Column      int      `json:"column" xml:"column"`
	Position    int      `json:"position" xml:"position"`
	Hidden      int      `json:"hidden" xml:"hidden"`
	Created     string   `json:"created" xml:"created"`
	Updated     string   `json:"updated" xml:"updated"`
	Assets      []*Asset `json:"assets,omitempty" xml:"assets>asset,omitempty"`
}

type Asset struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Name        string   `json:"name" xml:"name"`
	Type        string   `json:"type" xml:"type"`
	Description string   `json:"description" xml:"description,innerXML"`
	URL         string   `json:"url" xml:"url"`
	Redirect    string   `json:"redirect,omitempty" xml:"redirect,omitempty"`
	Owner       *Owner   `json:"owner" xml:"owner"`
	MapID       int      `json:"map_id" xml:"map_id"`
	Position    int      `json:"position" xml:"position"`
	Created     string   `json:"created" xml:"created"`
	Updated     string   `json:"updated" xml:"updated"`

	// Optional??
	Author          string `json:"author,omitempty" xml:"author,omitempty"`
	CallNumber      string `json:"call_number,omitempty" xml:"call_number,omitempty"`
	CoverURL        string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	Enabled         int    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	ISBN            string `json:"isbn,omitempty" xml:"isbn,omitempty"`
	PublicationDate string `json:"publication_date,omitempty" xml:"publication_date,omitempty"`
	FirstName       string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	EMail           string `json:"email,omitempty" xml:"email,omitempty"`
	MoreInfo        string `json:"more_info,omitempty" xml:"more_info,omitempty"`
}

func Slugify(s string) string {
	sluggy := map[string]string{
		" ": "-",
		"'": "",
		"@": "",
		"!": "",
		"?": "",
		"&": "",
		":": "",
		";": "",
		"<": "",
		">": "",
	}
	for k, v := range sluggy {
		s = strings.Replace(s, k, v, -1)
	}
	return strings.ToLower(s)
}

func Clean(src []byte) []byte {
	// Map out offensive chars from MS included content...
	return []byte(strings.Map(func(r rune) rune {
		switch r {
		case 0x0001, 0x000B, 0x000C, 0x0003, 0x0012, 0x0013:
			return -1
		default:
			return r
		}
	}, fmt.Sprintf("%s", src)))
}

func Decode(src []byte) (*LibGuides, error) {
	libguides := new(LibGuides)
	err := xml.Unmarshal(src, &libguides)
	return libguides, err
}

func (lg *LibGuides) ToXML() ([]byte, error) {
	return xml.Marshal(lg)
}

func (lg *LibGuides) ToJSON() ([]byte, error) {
	return json.Marshal(lg)
}

func (lg *LibGuides) ToString() string {
	buf, _ := lg.ToJSON()
	return fmt.Sprintf("%s", buf)
}
