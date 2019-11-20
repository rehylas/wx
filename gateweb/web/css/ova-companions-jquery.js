/**
 * Advanced companion processing scripts - supports insertion of javascript based
 * companion ad types.
 *
 * Original code contributed by Joe Connor April 28, 2010
 *
 * These javascript functions handle the DIV insertion of companion ad code - they
 * support advanced companion types that use SWF or Javascript embed code to
 * facilitate the display of the companion ad.
 *
 * These methods will only work when the ova plugin has the following config enabled:
 *
 *     "processCompanionsExternally": true
 *
 * When enabled, the OVA plugin will call these javascript functions to insert the
 * companion ad code rather than try to insert the code itself via the ExternalInterface.call
 * method.
 *
 * These methods rely on the prior inclusion of JQuery into the page.
 *
 */

var ova = new function(){
  this.readHTML = function(companion){
    return jQuery('#' + companion).html();
  }

  this.writeCompanion = function(companion, src){
    jQuery('#' + companion).writeCompanion(src);
  }

  this.writeHTML = function(companion, src){
    jQuery('#' + companion).writeOriginal(src);
  }

  /*
  note: not active - SWF treated as pre-formatted HTML with Object tags
  this.flashembed = function(companion, configuration){
    jQuery('#' + companion).flashembed(configuration);
  }
   */
}


// These jQuery functions are trigged from ova class
jQuery.fn.writeCompanion = function(src) {
  // Store jQuery(this) in a variable otherwise it will be out of scope in document.write
  var companionElem = jQuery(this);
  var tmpDocWrite = document.write;
  var tmpDocWriteln = document.writeln;
  document.write = function(arg){ jQuery(companionElem).append(arg); };
  document.writeln = function(arg){ jQuery(companionElem).append(arg) + '\n'; };

  // Function to retrieve a new advert from the server.
  jQuery(companionElem).html(jQuery(src));
};


jQuery.fn.writeOriginal = function(src) {
  var companionElem = jQuery(this);
  jQuery(companionElem).html(src);
};