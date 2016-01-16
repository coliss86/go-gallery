$(document).ready(function() {
  $(".fancybox").fancybox({
    helpers: {
      title: {
            type: 'inside',
            position: 'bottom'
        }
    },
    afterLoad: function() {
        this.title = '<a href="' + this.href.replace("/img/", "/download/") + '"><img src="/static/img/arrow_down_32.png" ></a><span title="' + this.title + '"> ' + this.title + '</span><span class="tag"><input type="text" name="tags" placeholder="Dossier" class="tm-input tm-input-small"/></span>';
    },
    afterShow: function() {
      var img = this.href.replace("/img/", "");
      var folders = [];
      for (var tag in tags) {
         folders.push(tag);
      }

      var tab = $(".tm-input").tagsManager({
        prefilled: folders,
        CapitalizeFirstLetter: true,
        deleteTagsOnBackspace: false,
        isClearInputOnEsc: false,
        tagCloseIcon: "&cross;"
      });

      var iof = img.lastIndexOf("/");
      var imgName = img;
      if (iof != -1) {
        imgName = img.substring(iof + 1);
      }
      for (var tag in tags) {
        var idx = $.inArray(imgName, tags[tag]);
        if (idx != -1) {
          tab.tagsManager('selectTag', tag);
        }
      }
      tab.on('tm:pushed', function(e, tag) {
        $.post("/tag/add/" + tag, { img: img });
        folders.push(tag);
      });
      tab.on('tm:spliced', function(e, tag) {
        $.post("/tag/delete/" + tag, { img: img });
        folders.splice(folders.indexOf(tag),1);
      });
      tab.on('tm:selected', function(e, tag) {
        $.post("/tag/select/" + tag, { img: img });
      });
      tab.on('tm:deselected', function(e, tag) {
        $.post("/tag/deselect/" + tag, { img: img });
      });
    },
    beforeLoad : function () {
      if (this.element.data("video")) {
        _videoHref   = this.href;
        _videoPoster = typeof this.element.data("poster")  !== "undefined" ? this.element.data("poster")  :  "";
        _videoWidth  = typeof this.element.data("width")   !== "undefined" ? this.element.data("width")   : 360;
        _videoHeight = typeof this.element.data("height")  !== "undefined" ? this.element.data("height")  : 360;
        _dataCaption = typeof this.element.data("caption") !== "undefined" ? this.element.data("caption") :  "";
        this.title = _dataCaption ? _dataCaption : (this.title ? this.title : "");
        this.content = "<video id='video_player' src='" + _videoHref + "'  poster='" + _videoPoster + "' width='" + _videoWidth + "' height='" + _videoHeight + "' controls='controls' preload='none' autoplay='true'></video>";
      }
    },
    loop: false,
    nextEffect: 'none',
    prevEffect: 'none',
    openEffect: 'elastic',
    closeEffect: 'elastic',
    keys : {
      next : {
  		34 : 'up',   // page down
  		39 : 'left', // right arrow
  		40 : 'up'    // down arrow
	   }
   }
  });
});
