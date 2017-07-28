// Ripple animation
(function($) {
  $(document).on('click', '.ripple', function(e){
    var rippler = $(this);

    // create .ink element if it doesn't exist
    //if(rippler.find(".ink").length == 0) {
    rippler.append('<span class="ink"></span>');
    //}

    var ink = rippler.find('.ink');

    // prevent quick double clicks
    // ink.removeClass("animate");

    for (var i = 0; i < ink.length; i++) {
      var cur = $(ink[i]);
      // set .ink diametr
      if(!cur.height() && !cur.width()) {
        var d = Math.max(rippler.outerWidth(), rippler.outerHeight());
        cur.css({height: d, width: d});
      }

      // get click coordinates
      var x = e.pageX - rippler.offset().left - cur.width()/2;
      var y = e.pageY - rippler.offset().top - cur.height()/2;

      // set .ink position and add class .animate
      cur.css({
        top: y+'px',
        left:x+'px'
      }).addClass("animate").bind('webkitAnimationEnd oanimationend msAnimationEnd animationend', function(e) {
        $(this).remove();
      });
    };

  });
})(jQuery);
