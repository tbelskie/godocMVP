/* godoc theme · Pagefind UI mount
 * Runs after pagefind-ui.js (defer). If the index has not been built
 * yet, PagefindUI is undefined and the header fallback input stays.
 */
(function () {
  'use strict';

  function init() {
    if (typeof window.PagefindUI !== 'function') {
      return;
    }
    var mount = document.getElementById('godoc-search');
    if (!mount) {
      return;
    }
    try {
      new window.PagefindUI({
        element: '#godoc-search',
        showImages: false
      });
    } catch (_) {
      /* keep fallback input */
    }
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
