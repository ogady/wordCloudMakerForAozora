package aozora

import "time"

// Author is author Info
type Author struct {
	PersonID  int    `json:"person_id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

// Book is book Info
type Book struct {
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
	BaseBookPublisher1          string    `json:"base_book_1_publisher"`
	BaseBookFirstEdition1       string    `json:"base_book_1_1st_edition"`
	BaseBookEditionInput1       string    `json:"base_book_1_edition_input"`
	BaseBookEditionProofing1    string    `json:"base_book_1_edition_proofing"`
	BaseBookParent1             string    `json:"base_book_1_parent"`
	BaseBookParentPublisher1    string    `json:"base_book_1_parent_publisher"`
	BaseBookParentFirstEdition1 string    `json:"base_book_1_parent_1st_edition"`
	BaseBook2                   string    `json:"base_book_2"`
	BaseBookPublisher2          string    `json:"base_book_2_publisher"`
	BaseBookFirstEdition2       string    `json:"base_book_2_1st_edition"`
	BaseBookEditionInput2       string    `json:"base_book_2_edition_input"`
	BaseBookEditionProofing2    string    `json:"base_book_2_edition_proofing"`
	BaseBookParent2             string    `json:"base_book_2_parent"`
	BaseBookParentPublisher2    string    `json:"base_book_2_parent_publisher"`
	BaseBookParentFirstEdition2 string    `json:"base_book_2_parent_1st_edition"`
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
	Translators                 []Author  `json:"translators"`
	Authors                     []Author  `json:"authors"`
}
