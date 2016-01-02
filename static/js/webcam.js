$(document).ready(function() {
  $(".fancybox").fancybox({
    helpers: {
      title: {
            type: 'inside',
            position: 'bottom'
        }
    },
    afterLoad: function() {
        this.title = '<a href="' + this.href.replace("/img/", "/download/") + '"><img src="/static/img/arrow_down_32.png" ></a> ' + this.title;
    },
    /*afterShow: function() {
        $('<div class="expander"></div>').appendTo(this.inner).click(function() {
            $(document).toggleFullScreen();
        });
    },*/
    loop: false,
    nextEffect: 'none',
    prevEffect: 'none',
    openEffect: 'elastic',
    closeEffect: 'elastic'
  });
});
