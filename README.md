# flickr-data-folderify
Processes the files flickr gives you when you do "download all my data", and creates folders for each flickr album for easier reuploading elsewhere.

Usage: 
Flickr will give you a bunch of zip files called something like

* 123123213123123-12bede123.zip -- this contains the metadata.
* data-download-1.zip
* ...
* data-download-99.zip -- these data-download-...zip files contain the images.

Unzip all the files into one folder. You will have a bunch of json files and all the images from your flickr account, but the filenames for those images will have flickr's extra image ID in there.

Build main.go.

run 

    [whatever go built] albums.json
    
in the folder you extracted to.

It will create a new folder for each album, and move the images into the folders for each album, and all the images that aren't in an album into a folder called "No Album".

WARNING:
It tries to restore the original filename for images that it moves into per-album folders, but if you had two images with the same original filename, behavior is unpredictable. (I didn't have any files like that so haven't tested it myself).

WARNING 2:
It's probably best to make a copy of the original zip files you get from flickr in another folder, and work with those, just to be on the safe side.


Note: The uploading part isn't automated, at least not for google photos. Google Photos does have an API, but if you use the API to upload images, they count towards your storage quota -- the 'high quality but free storage' part of google photos is only available if you upload through a browser. See [the starred note here](https://developers.google.com/photos/library/guides/upload-media) under 'The upload process'.
