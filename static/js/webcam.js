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
      });
      $(".tm-input").on('tm:spliced', function(e, tag) {
        $.post("/tag/delete/" + tag, { img: img });
      });
      $(".tm-input").on('tm:selected', function(e, tag) {
        $.post("/tag/select/" + tag, { img: img });
      });
      $(".tm-input").on('tm:deselected', function(e, tag) {
        $.post("/tag/deselect/" + tag, { img: img });
      });
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
