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
	"net/url"
)

const (
	DateFmt = "2006-01-02 15:04:05"
)

type LibGuides struct {
	XMLName  xml.Name   `json:"-"`
	Custmer  *Customer  `json:"customer" xml:"customer"`
	Site     *Site      `json:"site" xml:"libguides"`
	Accounts []*Account `json:"accounts" xml:"accounts"`
	Groups   []*Group   `json:"groups" xml:"groups"`
	Subjects []*Subject `json:"subjects" xml:"subjects"`
	Tags     []*Tag     `json:"tags" xml:"tags"`
	Vendors  string     `json:"vendors" xml:"vendors"` //FIXME: Vendors probably needs to be defined as a struct type...
	Guides   []*Guide   `json:"guides" xml:"guides"`
}

type Customer struct {
	XMLName  xml.Name `json:"-"`
	ID       int      `json:"id" xml:"id"`
	Type     string   `json:"type" xml:"type"`
	Name     string   `json:"name" xml:"name"`
	URL      *url.URL `json:"url" xml:"url"`
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
	Nickname  string   `json:"nickname" xml:"nickname"`
	Signature string   `json:"signature" xml:"signature"`
	Image     string   `json:"image" xml:"image"`
	Address   string   `json:"address" xml:"address"`
	Phone     string   `json:"phone" xml:"phone"`
	Skype     string   `json:"skype" xml:"skype"`
	Website   string   `json:"website" xml:"website"`
	Created   string   `json:"created" xml:"created"`
	Updated   string   `json:"updated" xml:"updated"`
}

type Group struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Type        string   `json:"type" xml:"type"`
	Description string   `json:"description" xml:"description,innerXML"`
	//Password    string     `json:"passord" xml:"password"`
	Name    string `json:"name" xml:"name"`
	Created string `json:"created" xml:"created"`
	Updated string `json:"updated" xml:"updated"`
}

type Subject struct {
	XMLName xml.Name `json:"-"`
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	URL     *url.URL `json:"url" xml:"url"`
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
	Description string     `json:"description" xml:"description,innerXML"`
	URL         *url.URL   `json:"url" xml:"url"`
	Owner       *Account   `json:"owner" xml:"owner"`
	Group       *Group     `json:"group" xml:"group"`
	Redirect    *url.URL   `json:"redirect" xml:"redirect"`
	Status      string     `json:"status" xml:"status"`
	Published   string     `json:"published" xml:"published"`
	Subjects    []*Subject `json:"subject" xml:"subject"`
	Tags        []*Tag     `json:"tags" xml:"tags"`
	Created     string     `json:"created" xml:"created"`
	Updated     string     `json:"updated" xml:"updated"`
	Pages       []*Page    `json:"pages" xml:"pages"`
}

type Page struct {
	XMLName      xml.Name `json:"-"`
	ID           int      `json:"id" xml:"id"`
	Name         string   `json:"name" xml:"name"`
	Description  string   `json:"description" xml:"description,innerXML"`
	URL          *url.URL `json:"url" xml:"url"`
	Redirect     *url.URL `json:"redirect" xml:"redirect"`
	SourcePageID int      `json:"source_page_id" xml:"source_page_id"`
	ParentPageID int      `json:"parent_page_id" xml:"parent_page_id"`
	Position     int      `json:"position" xml:"position"`
	Hidden       int      `json:"hidden" xml:"hidden"`
	Created      string   `json:"created" xml:"created"`
	Updated      string   `json:"updated" xml:"updated"`
	Boxes        []*Box   `json:"boxes" xml:"boxes"`
}

type Box struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Name        string   `json:"name" xml:"name"`
	Type        string   `json:"type" xml:"type"`
	Description string   `json:"description" xml:"description,innerXML"`
	URL         *url.URL `json:"url" xml:"url"`
	Owner       *Account `json:"owner" xml:"owner"`
	MapID       int      `json:"map_id" xml:"map_id"`
	Column      int      `json:"column" xml:"column"`
	Position    int      `json:"position" xml:"position"`
	Hidden      int      `json:"hidden" xml:"hidden"`
	Created     string   `json:"created" xml:"created"`
	Updated     string   `json:"updated" xml:"updated"`
	Assets      []*Asset `json:"assets" xml:"assets"`
}

type Asset struct {
	XMLName     xml.Name `json:"-"`
	ID          int      `json:"id" xml:"id"`
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description,innerXML"`
	URL         *url.URL `json:"url" xml:"url"`
	Redirect    *url.URL `json:"redirect" xml:"redirect"`
	Owner       *Account `json:"owner" xml:"owner"`
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
	PublicationDate string `json:"publication_date" xml:"publication_date"`
	FirstName       string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	EMail           string `json:"email,omitempty" xml:"email,omitempty"`
	MoreInfo        string `json:"more_info,omitempty" xml:"more_info,omitempty"`
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
