Go gallery
========

Go gallery is a simple web gallery written in [golang](https://golang.org) to browse a folder  hierarchy containing pictures. It displays them in a nice and elegant way.

<img src="/docs/folders1.jpg?raw=true" alt="Root folder" width="49%">

It sorts folders by month and day based on theis name and by applying the regular expression `^([0-9]+)-([0-9]+)`. It then displays them chronologically.

<img src="/docs/folders2.jpg?raw=true" alt="Folder listing" width="49%">

It recognizes `jpeg|jpg|gif|png|bmp` pictures and `mp4|m4v|mpeg|mpg|avi` videos and sort them by name, pictures first.

<img src="/docs/pictures.jpg?raw=true" alt="Photo listing" width="49%">
<img src="/docs/pictures2.jpg?raw=true" alt="Photo + Video listing" width="49%">

When clicking on a picture or video, a nice [fancybox](http://fancyapps.com/fancybox/) is displayed. Videos are played in it with the help of the `<video>` introduced in HTML 5.
Keybinding are enabled in the navigation : <kbd>&rarr;</kbd>, <kbd>&larr;</kbd>, <kbd>ESC</kbd>

<img src="/docs/fancybox.jpg?raw=true" alt="Fancybox" width="49%">

The folder hierarchy may be as follows :
```shell
├──2013
├──2014
├──2015
└──2016
    ├── 01-01_NewYear
    │   ├── picture1.jpg
    │   ├── picture2.jpg
    │   ├── video1.mp4
    │   └── ...
    ├── 06-05_Beach
    └── ...
```

Prerequisites
----------

This gallery uses [convert from ImageMagick](http://www.imagemagick.org/script/convert.php) in order to generate thumbnails.

On debian :
```console
apt-get install imagemagick
```

On a mac, the [installation](http://www.imagemagick.org/script/binary-releases.php#macosx) requires [port](https://www.macports.org/) to be available and then :
```console
sudo port install ImageMagick
```

`go-gallery` requires `golang` in version 1.4.2 at least.

Install
----------
Select an empty folder and run :
```console
export GOPATH=`pwd`
go get gitlab.com/coliss86/go-gallery
```
It downloads the source from gitlab and compiles it. The resulting binary is located in `$GOPATH/bin`.
To run it :
```console
$GOPATH/bin/go-gallery <config file>
```
  * 'config file' : *mandatory* path to the config file

Here is a example of this config file :
```
images=/Users/user/Pictures
export=/Users/user/export
cache=/Users/user/temp/cache
port=9090
```

Reporting Issues
----------
  * Please report issues on [Gitlab Issue Tracker](https://gitlab.com/coliss86/go-gallery/issues).
  * In your report, please provide steps to **reproduce** the issue.
  * Before reporting:
     * Make sure you are using the latest master.
     * Check existing open issues.
  * Merge requests, documentation requests, and enhancement ideas are welcome.

License
----------
"Go gallery" is distributed under [GNU GPL v3](http://www.gnu.org/licenses/gpl-3.0.en.html) license, see LICENSE.

It also uses third-party libraries and programs:
  * fancyBox ([Creative Commons Attribution-NonCommercial 3.0 License](http://creativecommons.org/licenses/by-nc/3.0/)) :  http://fancyapps.com/fancybox/
  * famfamfam silk icons ([Creative Commons Attribution 2.5 License](http://creativecommons.org/licenses/by/2.5/) license) : http://www.famfamfam.com/lab/icons/silk/
  * woofunction icons ([GNU General Public License](http://www.gnu.org/licenses/gpl.html) license) : http://www.iconarchive.com/show/woofunction-icons-by-wefunction.html
  * Tag Manager (a jQuery plugin) ([Mozilla Public License 2.0](https://www.mozilla.org/en-US/MPL/2.0/) license) : https://maxfavilli.com/jquery-tag-manager
  * Icon Author: sa-ki (License: Free for personal non-commercial use) : http://sa-ki.deviantart.com
  * A configuration file parser library for Go / Golang (The MIT License) : ~~https://github.com/jimlawless/cfg/~~
  * Gorilla web toolkit ([BSD licensed](https://opensource.org/licenses/BSD-2-Clause)) : http://www.gorillatoolkit.org/
  * Play icon (CC0 1.0 Universal (CC0 1.0) Public Domain Dedication License) : http://www.iconsdb.com/caribbean-blue-icons/video-play-3-icon.html
  * jQuery Lazy (MIT License) : http://jquery.eisbehr.de/lazy/
