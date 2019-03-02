var callback = function(){

    /* Drop down button */
    $(".button a").click(function(){
        $(".overlay").fadeToggle(200);
       $(this).toggleClass('btn-open').toggleClass('btn-close');
    });

    var swiper = new Swiper('.swiper-container', {
        pagination: {
          el: '.swiper-pagination',
          dynamicBullets: true,
        },
    });

    /* scroll */
    new WOW().init();

    $('#fullpage').fullpage({
        scrollBar: true,
        navigation: true,
        navigationTooltips: ['Slide 1', 'Slide 2', 'Slide 3', 'Slide 4'],
        loopBottom: true,
        sectionSelector: 'section'
    });
    licenseKey: 'OPEN-SOURCE-GPLV3-LICENSE'
};
  
if (
    document.readyState === "complete" ||
    (document.readyState !== "loading" && !document.documentElement.doScroll)
) {
    callback();
} else {
    document.addEventListener("DOMContentLoaded", callback);
}

$('.overlay').on('click', function(){
    $(".overlay").fadeToggle(200);   
    $(".button a").toggleClass('btn-open').toggleClass('btn-close');
    open = false;
});
