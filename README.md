# Cloudinary Image Management System
Is a simple repo which creates routes to upload and get images from your Cloudinary database.
## Routes
"/getImageByPublicID?publicID=`public id`?toDownload=bool" To fetch image via its public id <br>
"/upload" To upload image

## Problem Statement

  Create endpoints to upload images onto cloudinary using a system similar to Document model in NoSQL. So each images is under a collection in our case a directory.
  Function to be created:
  1) To upload photos to a collection
  2) To retreive photos from a collection
