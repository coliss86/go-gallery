$(document).ready(function() {
  $(".fancybox").fancybox({
    helpers: {
      title: {
            type: 'inside',
            position: 'bottom'
        }
    },
    afterLoad: function() {
        this.title = '<a href="' + this.href + '">Download</a> ' + this.title;
    },
    afterShow: function() {
        $('<div class="expander"></div>').appendTo(this.inner).click(function() {
            $(document).toggleFullScreen();
        });
    },
    loop: false,
    nextEffect: 'none',
    prevEffect: 'none',
    openEffect: 'elastic',
    closeEffect: 'elastic'
  });
});
