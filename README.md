go-colorweave
=============

Extracting dominant color palette from an image

This program extracts the top N dominant colors from an image and display as output, where N can be set by the user. It uses standard image library, [resize library](https://github.com/nfnt/resize) and [webcolors library](https://github.com/jyotiska/go-webcolors). Along with the colors, the program also shows the percentage of the color in the given image. By default, the code displays the color name using CSS 2.1 specifications, but this can be also changes to CSS 3 colors which has wider variety (140) colors.

By default the program takes the <code>test.jpg</code> file in <images> directory. However, it can be changed to any image file.

With the 2 external packages installed, the code can be executed as following:

```bash
go run go-colorweave.go
```

Example
=======

The following image has been used as example:

<img src="http://lokeshdhakar.com/projects/color-thief/img/photo2.jpg"></img>

The program will give the following output when used agains the image above:

```
green 26.77%
silver 19.77%
olive 16.79%
grey 14.47%
aqua 11.01%
```
