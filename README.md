# Go Image Editing Command Line Tools
golang command line image editor
Use one command, gidit or install them all seperately

# Install

```
go get github.com/bsdpunk/gidit
```


## gisize
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


## giadd
Append an image, to the left of an image


![one](one.png)


<br />


![two](two.png)



```
gidit -o add one.png two.png
```
Result is newtwo.png 


![newtwo](newtwo.png)

## giappend
Append an image to the bottom of an image

![newtwo](newtwo.png)

```
gidit -o append newtwo.png newtwo.png
```

Result is newnewtwo.png

![newnewtwo](newnewtwo.png)


