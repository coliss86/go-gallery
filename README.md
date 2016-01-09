Go gallery
========

Go gallery is a simple web gallery written in [golang](https://golang.org) to browse pictures hierarchy stored in a folder. It allows browsing them with a nice and elegant way. It recognized month in folder name and sort them according.

<img src="/doc/folders.png?raw=true" alt="Folder listing" width="30%">
<img src="/doc/photo.png?raw=true" alt="Photo listing" width="30%">
<img src="/doc/fancybox.png?raw=true" alt="Fancybox" width="30%">

Prerequisites
----------

This gallery use [convert from ImageMagick](http://www.imagemagick.org/script/convert.php) in order to generate thumbnails.

On debian :
```
apt-get install imagemagick
```

On a mac, the [installation](http://www.imagemagick.org/script/binary-releases.php#macosx) requires [port](https://www.macports.org/) to be installed and then :
```
sudo port install ImageMagick
```

Install
----------

```console
$ export GOPATH=`pwd`
$ go get github.com/gmembre/go-gallery
```

The binary is located in `$GOPATH/bin`.
To run it :
```
$GOPATH/bin/go-gallery <config file>
```
  * 'config file' : *mandatory* path to the config file in [toml](https://github.com/toml-lang/toml) format

Here is a example of this config file :
```toml
images = "/Users/user/Pictures"
export = "/Users/user/export"
cache = "/Users/user/temp/cache"
#port=9090
```

Reporting Issues
----------
  * Please report issues on [Github Issue Tracker](https://github.com/gmembre/go-gallery/issues).
  * In your report, please provide steps to **reproduce** the issue.
  * Before reporting:
     * Make sure you are using the latest master.
     * Check existing open issues.
  * Pull requests, documentation requests, and enhancement ideas are welcome.

License
----------
"Go gallery" is distributed under [GNU GPL v3](http://www.gnu.org/licenses/gpl-3.0.en.html) license, see LICENSE.

It also uses third-party libraries and programs:
  * fancyBox ([Creative Commons Attribution-NonCommercial 3.0 License](http://creativecommons.org/licenses/by-nc/3.0/)) :  http://fancyapps.com/fancybox/
  * famfamfam silk icons ([Creative Commons Attribution 2.5 License](http://creativecommons.org/licenses/by/2.5/) license) : http://www.famfamfam.com/lab/icons/silk/
  * woofunction icons ([GNU General Public License](http://www.gnu.org/licenses/gpl.html) license) : http://www.iconarchive.com/show/woofunction-icons-by-wefunction.html
  * Tag Manager (a jQuery plugin) ([Mozilla Public License 2.0](https://www.mozilla.org/en-US/MPL/2.0/) license) : https://maxfavilli.com/jquery-tag-manager
  * Icon Author: sa-ki (License: Free for personal non-commercial use) : http://sa-ki.deviantart.com
  * TOML parser and encoder for Go ([wtfpl License](http://www.wtfpl.net/)) : https://github.com/BurntSushi/toml
  * Gorilla web toolkit ([BSD licensed](https://opensource.org/licenses/BSD-2-Clause)) http://www.gorillatoolkit.org/
