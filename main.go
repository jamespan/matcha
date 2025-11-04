package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mmcdole/gofeed"
	_ "modernc.org/sqlite"
)

func main() {
	bootstrapConfig()

	fp := gofeed.NewParser()
	writer := getWriter()
	displayWeather(writer)
	displaySunriseSunset(writer)
	generateAnalysis(fp, writer)

	for _, feed := range myFeeds {
		parsedFeed := parseFeed(fp, feed.url, feed.limit)

		if parsedFeed == nil {
			continue
		}

		items := generateFeedItems(writer, parsedFeed, feed)
		if items != "" {
			writeFeed(writer, parsedFeed, items)
		}
	}

	if !terminalMode {
		markdown_file_name := mdPrefix + currentDate + mdSuffix + ".md.tmp"
		if err := os.Rename(filepath.Join(markdownDirPath, markdown_file_name), filepath.Join(markdownDirPath, mdPrefix+currentDate+mdSuffix+".md")); err != nil {
			log.Fatal(err)
		}
	}

	defer db.Close()
}
