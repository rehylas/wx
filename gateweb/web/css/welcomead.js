function PopUpWelcome(host, path) {
       
    var url = host + path + '#' + encodeURIComponent(document.location.href);

    var overlay = document.createElement("div");
    jQuery(overlay).css({
        'position': (jQuery.browser.msie && jQuery.browser.version == "6.0") ? 'absolute' : 'fixed', // IE6 = absolute
        'top': 0,
        'left': 0,
        'background-color': '#000',
        'width': '100%',
        'height': '100%',
        'z-index': 601,
        'opacity': 0.8,
        'display': 'none'
    });
    jQuery(document.body).append(overlay);

    var container = document.createElement("div");
    jQuery(container).css({
        'position': (jQuery.browser.msie && jQuery.browser.version == "6.0") ? 'absolute' : 'fixed', // IE6 = absolute
        'top': '50%',
        'left': '50%',
        'width': '100px',
        'height': '0px',
        'margin-left': 0,
        'margin-top': 0,
        'z-index': 602,
        'display': 'block',
        'border-width': 0,
        'background-color': '#fff'
    });
    jQuery(document.body).append(container);

    var iframe = document.createElement("iframe");
    jQuery(iframe).attr({
        'id': 'welcomead_iframe',
        'frameborder': '0',
        'hspace': 0,
        'scrolling': 'no',
        'src': url
    }).css({
        'width': '100%',
        'height': '100%'
    });
    jQuery(container).append(iframe);

    XD.receiveMessage(function(message) {
        if (message.data == "close") {
            closeWelcomeAd();
        }
        else {
            var width = 800, height = 600;

            var dimension = message.data.split('x');
            width = dimension[0];
            height = dimension[1];

            jQuery(container).hide();

            if (jQuery(window).height() > parseInt(height)) {
                jQuery(container).css({
                    'width': width + 'px',
                    'height': height + 'px',
                    'margin-left': '-' + (width / 2) + 'px',
                    'margin-top': '-' + (height / 2) + 'px'
                });
            }
            else {
                jQuery(container).css({
                    'position': 'absolute',
                    'width': width + 'px',
                    'height': height + 'px',
                    'top': '0px',
                    'left': '0px'
                });
            }

            if (jQuery(window).height() > jQuery(document.body).height()) {
                jQuery(overlay).css({ 'height': jQuery(window).height() });
            }
            else {
                jQuery(overlay).css({ 'height': jQuery(document.body).height() });
            }

            jQuery(overlay).fadeIn();
            jQuery(container).fadeIn();
        }
    }, host);

    function closeWelcomeAd() {
        jQuery(container).fadeOut('slow', function() { jQuery(container).remove(); });
        jQuery(overlay).fadeOut('slow', function() { jQuery(container).remove(); });
    }

}
