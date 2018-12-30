package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Album struct {
	Photo_count, View_count, Created, Last_updated string   // actually numbers, but strings in the JSON
	Id, Url, Title, Description, Cover_photo       string   // these are really strings
	Photos                                         []string // these IDs all look like numbers to me.
}

type Albums struct {
	Albums []Album
}

func check(e error) {
	if e != nil {
		// Pretty basic error handling.
		panic(e)
	}
}

func main() {

	rawfile, err := ioutil.ReadFile(os.Args[1])
	check(err)

	var albums Albums
	err = json.Unmarshal(rawfile, &albums)
	check(err)

	// We need to read the list of all the images we have available to us.
	files, err := ioutil.ReadDir(".")
	check(err)

	for i := 0; i < len(albums.Albums); i++ {
		album := albums.Albums[i]
		// If the album contains illegal characters, fix that. Just do a trivial fix for
		// now -- flickr albums can contain slashes, but filenames (most likely) can't.
		albumTitle := album.Title
		albumTitle = strings.Replace(albumTitle, "/", "", -1)
		albumTitle = strings.Replace(albumTitle, "\\", "", -1)
		fmt.Printf("%v / %v : %v\n", i+1, len(albums.Albums), album.Title)
		os.Mkdir(album.Title, 0777)
		for j := 0; j < len(album.Photos); j++ {
			photoid := album.Photos[j]
			// This is a string of the form "1234567890", which is the flickr ID for this
			// photo.
			// There will be an equivalent file of the form "img1234_1234567890_o.jpg"
			// where the section before the photo name is the original filename for the
			// photo, which we can use to recreate the original file structure used. Note
			// that multiple files can have the same original filename, so if you haven't
			// put them into albums, this will fail.
			//
			// Videos get the filename "vid1234_1234567890.mov", without the _o on the end,
			// so we need to watch for that as well.
			//
			// We could try reading the original filename from the photo_123456789.json
			// file, it's in "name" there, but that seems redundant.

			//fmt.Println("Searching for photo ID " + photoid)

			// So find the photo with that ID in there. Search our current folder for a
			// matching file. Remember to skip the photo_12344576890.json files that are
			// also in the current folder, though.
			for _, file := range files {
				// Let's assume that we only have one filename with _1234567890_ in there.
				// We know (hopefully) that flickr will have only made one file like that,
				// so as long as none of the metadata json matches it, we should be okay.
				filename := file.Name()
				extensionIndex := strings.LastIndex(filename, ".")
				if extensionIndex != -1 {
					extension := filename[extensionIndex:]
					if extension != ".json" {
						idIndex := strings.Index(filename, "_"+photoid)
						if idIndex != -1 {
							fmt.Println("Found: " + filename)
							// We need to extract the original filename.
							origName := filename[:idIndex] + filename[extensionIndex:]
							fmt.Println("  original filename: " + origName + " moving to " + albumTitle)

							// WARNING: If your original album was made by uploading two
							// images with the same original filename, behaviour is not
							// guaranteed to do what you expect here. I don't have an
							// album like that to test with, but there is a risk of something
							// unexpected happening if you do. TODO: make fake data, rename
							// overlapping filenames or something. Maybe just use flickr's
							// filename in this sort of situation.
							os.Rename(filename, albumTitle+string(os.PathSeparator)+origName)
						}
					}
				}
			}
		}
	}

	// Now deal with any images that weren't in any of the albums. We'll put all those into
	// a folder called "No Album" and hope that nobody has called an album by that
	// name.
	os.Mkdir("No Album", 0777)
	files, err = ioutil.ReadDir(".")
	for _, file := range files {
		// Here, we'll just copy all the (non-json) files to
		// that folder, we won't try and fix the filenames, to
		// avoid the risk of clashes.
		filename := file.Name()
		extensionIndex := strings.LastIndex(filename, ".")
		if extensionIndex != -1 {
			extension := filename[extensionIndex:]
			// Remember to skip the zip files as well if they're there.
			if extension != ".json" && extension != ".zip" {
				os.Rename(filename, "No Album"+string(os.PathSeparator)+filename)
			}
		}
	}
}
