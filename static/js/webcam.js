$(document).ready(function() {
  $(".fancybox").fancybox({
    helpers: {
      title: {
            type: 'inside',
            position: 'bottom'
        }
    },
    afterLoad: function() {
        this.title = '<a href="' + this.href.replace("/img/", "/download/") + '"><img src="/static/img/arrow_down_32.png" ></a> ' + this.title + '<span class="tag"><input type="text" name="tags" placeholder="Ajouter un dossier" class="tm-input tm-input-small"/></span>';
    },
    afterShow: function() {
      var img = this.href.replace("/img/", "");

      $(".tm-input").tagsManager({
        prefilled: tags,
        CapitalizeFirstLetter: true,
        deleteTagsOnBackspace: false,
        isClearInputOnEsc: false,
        tagCloseIcon: "&cross;"
      });
      $(".tm-input").on('tm:pushed', function(e, tag) {
        $.post("/tag/add/" + tag, { img: img });
        tags.push(tag);
      });
      $(".tm-input").on('tm:spliced', function(e, tag) {
        $.post("/tag/delete/" + tag, { img: img });
        tags.splice(tags.indexOf(tag),1);
      });
      $(".tm-input").on('tm:selected', function(e, tag) {
        $.post("/tag/select/" + tag, { img: img });
      });
      $(".tm-input").on('tm:deselected', function(e, tag) {
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
