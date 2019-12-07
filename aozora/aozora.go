package aozora

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BOOKS_ENDPOINT = "http://pubserver2.herokuapp.com/api/v0.1/books/"
)

type Author struct {
	PersonID  int    `json:"person_id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type BookInfo struct {
	BookID                      int       `json:"book_id"`
	Title                       string    `json:"title"`
	TitleYomi                   string    `json:"title_yomi"`
	TitleSort                   string    `json:"title_sort"`
	Subtitle                    string    `json:"subtitle"`
	SubtitleYomi                string    `json:"subtitle_yomi"`
	OriginalTitle               string    `json:"original_title"`
	FirstAppearance             string    `json:"first_appearance"`
	NDCCode                     string    `json:"ndc_code"`
	FontKanaType                string    `json:"font_kana_type"`
	Copyright                   bool      `json:"copyright"`
	ReleaseDate                 time.Time `json:"release_date"`
	LastModified                time.Time `json:"last_modified"`
	CardURL                     string    `json:"card_url"`
	BaseBook1                   string    `json:"base_book_1"`
	BaseBook1Publisher          string    `json:"base_book_1_publisher"`
	BaseBook1FirstEdition       string    `json:"base_book_1_1st_edition"`
	BaseBook1EditionInput       string    `json:"base_book_1_edition_input"`
	BaseBook1EditionProofing    string    `json:"base_book_1_edition_proofing"`
	BaseBook1Parent             string    `json:"base_book_1_parent"`
	BaseBook1ParentPublisher    string    `json:"base_book_1_parent_publisher"`
	BaseBook1ParentFirstEdition string    `json:"base_book_1_parent_1st_edition"`
	BaseBook2                   string    `json:"base_book_2"`
	BaseBook2Publisher          string    `json:"base_book_2_publisher"`
	BaseBook2FirstEdition       string    `json:"base_book_2_1st_edition"`
	BaseBook2EditionInput       string    `json:"base_book_2_edition_input"`
	BaseBook2EditionProofing    string    `json:"base_book_2_edition_proofing"`
	BaseBook2Parent             string    `json:"base_book_2_parent"`
	BaseBook2ParentPublisher2   string    `json:"base_book_2_parent_publisher"`
	BaseBook2ParentFirstEdition string    `json:"base_book_2_parent_1st_edition"`
	Input                       string    `json:"input"`
	Proofing                    string    `json:"proofing"`
	TextURL                     string    `json:"text_url"`
	TextLastModified            time.Time `json:"text_last_modified"`
	TextEncoding                string    `json:"text_encoding"`
	TextCharset                 string    `json:"text_charset"`
	TextUpdated                 int       `json:"text_updated"`
	HTMLURL                     string    `json:"html_url"`
	HTMLLastModified            time.Time `json:"html_last_modified"`
	HTMLEncoding                string    `json:"html_encoding"`
	HTMLCharset                 string    `json:"html_charset"`
	HTMLUpdated                 int       `json:"html_updated"`
	Authors                     []Author  `json:"authors"`
}

func GetBookInfoByTitleName(titleName string) (string, error) {

	url := fmt.Sprintf("%s?title=%s", BOOKS_ENDPOINT, titleName)
	// APIを叩いてデータを取得
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 取得したデータをJSONデコード
	var bookInfo BookInfo
	err = json.Unmarshal(body, &bookInfo)
	if err != nil {
		return "", err
	}

	return bookInfo.HTMLURL, nil
}
