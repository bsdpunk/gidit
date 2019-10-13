# Go Image Editing Command Line Tools
golang command line image editor

## gidit
Resize png file
Examples:
```
gidit 1000 file.png
gidit 1000 500 file.png
```
If only one size is given the aspect ratio will be maintained

Sample:
![one](one.png)
```
gidit 30 one.png
```

Result is newone.png
![newone](newone.png)


## gadd
Append an image, to the left of an image

![one](one.png)

<br />

![two](two.png)


```
gadd one.png two.png
```
Result is newtwo.png 
![newtwo](newfist.png)

## gbot
Append an image to the bottom of an image
![newtwo](newtwo.png)
```
gbot newtwo.png newtwo.png
```

Result is newnewtwo.png
![newnewtwo](newnewtwo.png)


