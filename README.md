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
or
gisize 1000 file.png
gisize 1000 500 file.png
```
If only one size is given the aspect ratio will be maintained

Sample:

![out](one.png)

```
gidit 30 one.png
or
gisize 30 one.png
```

Result is out.png

![newone](newone.png)


## gicat
Append an image, to the right of an image


![one](one.png)


<br />


![two](two.png)



```
gidit cat one.png two.png
or
gicat one.png two.png
```
Result is newtwo.png 


![newtwo](newtwo.png)

## giappend
Append an image to the bottom of an image

![newtwo](newtwo.png)

```
gidit append newtwo.png newtwo.png
or
giappend newtwo.png newtwo.png
```

Result is newnewtwo.png

![newnewtwo](newnewtwo.png)


## gicreate
Create empty image


```
gidit create 200 200
or
gicreate 200 200
```
