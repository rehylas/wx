function readCookie(name) {
    var nameEQ = name + "=";
    var nameComma = name + ",";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) == 0 || c == name) return c.substring(nameEQ.length,c.length);
    }
    return null;
}

function randStyle() {
    var now = new Date();
    var sec = now.getSeconds();
    var index = sec % 2;
    return index;
}

function showWelcomeAd(skipExternalReferrers) {
    var cookie = readCookie('AdViewed');
    fullURL = parent.document.URL;
    
    if ((skipExternalReferrers == true) && (!document.referrer.match(/http:\/\/[^\/]*investopedia.com/i)) && (document.referrer != "")) {
        return;
    }

    if (!document.referrer.match(/http:\/\/[^\/]*investopedia.com/i)) {
	    document.cookie = "inv_originalReferrer=" + document.referrer + ";path=/;domain=.investopedia.com";
	    document.cookie = "inv_originalDomain=;domain=.investopedia.com; path=/;expires=Thu, 01-Jan-70 00:00:01 GMT";
    }

    var viewed = fullURL.indexOf('viewed=1');
    if ((cookie == null) && (viewed == -1)) {
        //Redirect
        var style = randStyle();
        
        var www = window.location.host;
        var subdomain = www.match(/[a-z0-9]+/i)[0].toLowerCase();
        if (subdomain != "www" && subdomain != "investopedia")
        {
            if (www.search("investopedia.com") >= 0)
                www = www.replace(/[a-z0-9]+/i, "www");
            else if (www.search("equade.com") < 0)
                www = www.replace(/[a-z0-9]+/i, "investopedia");
        }
        PopUpWelcome("http://" + www, "/rotate.aspx?sp=0&backurl=" + escape(fullURL) + "&style=" + style + "&oas_s=" + escape(OAS_sitepage) + "&oas_q=" + escape(OAS_query));
    }
}

jQuery(document).ready(function(){
	// sometimes we skip the ad if the user was referred from an outside website
	var skipExternalReferrers = false;
	if (typeof g_welcomeAdSkipExternalReferrers != "undefined") {
		if (g_welcomeAdSkipExternalReferrers == true) skipExternalReferrers = true;
	}
	// show the ad
	showWelcomeAd(skipExternalReferrers);
});
