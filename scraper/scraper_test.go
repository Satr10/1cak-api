package scraper

import (
	"testing"
)

func TestHasil(t *testing.T) {
	scrape := NewScraper()
	hasil, err := scrape.RandomMeme()
	if err != nil {
		t.Fatalf("Error Mendapatkan Meme: %v", err)
	}
	t.Logf("Hasil: %v", hasil)
}
