# flickr-data-folderify
Processes the files flickr gives you when you do "download all my data", and creates folders for each flickr album for easier reuploading elsewhere.

Usage: 
Flickr will give you a bunch of zip files called something like
123123213123123-12bede123.zip - this contains the metadata
data-download-1.zip
...
data-download-10.zip 
where these contain the images.

Unzip all the files into one folder. You will have a bunch of json files and all the images from your flickr account.
Build the app.
run 'flickr-data-folderify albums.json'
It will create a new folder for each album, and move the images into the folders for each album, and all the images that aren't in an album into a folder called "No Album".

WARNING:
It tries to restore the original filename for images that it moves into per-album folders, but if you had two images with the same original filename, behavior is unpredictable. (I don't have any data like that so haven't tested it myself).

