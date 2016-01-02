Go gallery
========

Go gallery is a simple web gallery written in [golang](https://golang.org) to browse pictures hierarchy stored in a folder. It allows browsing them with a nice and elegant way. It recognized month in folder name and sort them according.

<img src="/doc/folders.png?raw=true" alt="Folder listing" width="30%">
<img src="/doc/photo.png?raw=true" alt="Photo listing" width="30%">
<img src="/doc/fancybox.png?raw=true" alt="Fancybox" width="30%">

Prerequisites
----------

This gallery use [convert from ImageMagick](http://www.imagemagick.org/script/convert.php) in order to generate thumbnails.

On a debian :
```
apt-get install imagemagick
```

On a mac, the [installation](http://www.imagemagick.org/script/binary-releases.php#macosx) requires [port](https://www.macports.org/) to be installed and then :
```
sudo port install ImageMagick
```

How to use
----------

Clone this repo, set GOPATH and build it :
```
git clone ...
cd ...
export GOPATH=`pwd`
go install webcam
```

The binary is located in `$GOPATH/bin`.
To run it :
```
$GOPATH/bin/webcam <photo folder> <temp folder> [port]
```
  * 'photo folder' : *mandatory* root directory of photos
  * 'temp folder' : *mandatory* temp folder to store thumbnails
  * 'port' : *optional* http port to bind to, default is `9090`


License
----------
"Go gallery" is distributed under GNU GPL v3 license, see LICENSE.

It also uses third-party librarires and programs:
  * fancyBox ([Creative Commons Attribution-NonCommercial 3.0](http://creativecommons.org/licenses/by-nc/3.0/) license) :  http://fancyapps.com/fancybox/
  * famfamfam silk icons ([Creative Commons Attribution 2.5 License](http://creativecommons.org/licenses/by/2.5/) license) : http://www.famfamfam.com/lab/icons/silk/
  * woofunction icons ([GNU General Public License](http://www.gnu.org/licenses/gpl.html) license) : http://www.iconarchive.com/show/woofunction-icons-by-wefunction.html
