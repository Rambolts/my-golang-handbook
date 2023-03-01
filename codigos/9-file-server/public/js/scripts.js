$(function(){
    
    //DESLISA A PÁGINA!
    $('a').click(function(){
        var goto = $('.' + $(this).attr('href').replace('#', '')).position().top;
        $('html body').animate({scrollTop: goto - $('.main_header').outerHeight()}, 1000);
        return false;
    });
    
    //RETORNA TOPO
    $('.j_back').click(function(){
        $('html body').animate({scrollTop: 0}, 1000);
    });
    
    //FIXA MENU
    $(window).scroll(function(){
        if($(this).scrollTop() > $('.main_header').outerHeight() + 140){
            $('body').css('padding-top' , $('.main_header').outerHeight());
            $('.main_header').addClass('main_header_fixed');
            $('.j_back').fadeIn('slow');
        }else{
            $('body').css('padding-top' , 0);
            $('.main_header').removeClass('main_header_fixed');
            $('.j_back').fadeOut('slow');
        }
    });
    
    //FORM SUBMIT
    $('.j_formsubmit').submit(function(){
        $.ajax({
            beforeSend: function(){
                $('.form_load').fadeIn();
            },
            success: function (data) {
                console.clear();
                console.log(data);
                $('.form_load').fadeOut();
            }
        });
        return false;
    });
});