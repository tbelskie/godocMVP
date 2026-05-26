/* Pendragon theme · MVP 1.0
 * Three behaviors, no dependencies, no build:
 *  1. theme toggle  — flips [data-theme] on <html>, persists in localStorage
 *  2. sidebar collapse — toggles aria-expanded + [hidden] on its <ul>
 *  3. mobile hamburger — toggles [data-open] on the sidebar
 *
 * Flash-of-unstyled-content prevention runs inline in <head>; this
 * script only handles user-initiated interactions, so it is loaded
 * with `defer`.
 */
(function () {
  'use strict';

  var root = document.documentElement;
  var STORAGE_KEY = 'pendragon-theme';

  function safeRead() {
    try { return localStorage.getItem(STORAGE_KEY); } catch (_) { return null; }
  }
  function safeWrite(value) {
    try { localStorage.setItem(STORAGE_KEY, value); } catch (_) { /* noop */ }
  }

  // ---- theme toggle ----
  var toggle = document.querySelector('[data-theme-toggle]');
  if (toggle) {
    toggle.addEventListener('click', function () {
      var stored = safeRead();
      var systemLight =
        window.matchMedia &&
        window.matchMedia('(prefers-color-scheme: light)').matches;
      var current =
        root.getAttribute('data-theme') ||
        (stored === 'light' || stored === 'dark' ? stored : (systemLight ? 'light' : 'dark'));
      var next = current === 'dark' ? 'light' : 'dark';
      if (next === 'light') {
        root.setAttribute('data-theme', 'light');
      } else {
        root.setAttribute('data-theme', 'dark');
      }
      safeWrite(next);
    });
  }

  // ---- sidebar collapsible sections ----
  Array.prototype.forEach.call(
    document.querySelectorAll('[data-nav-toggle]'),
    function (btn) {
      btn.addEventListener('click', function () {
        var expanded = btn.getAttribute('aria-expanded') === 'true';
        btn.setAttribute('aria-expanded', String(!expanded));
        var listId = btn.getAttribute('aria-controls');
        var list = listId && document.getElementById(listId);
        if (list) list.hidden = expanded;
      });
    }
  );

  // ---- mobile hamburger ----
  var hamburger = document.querySelector('[data-hamburger]');
  var sidebar = document.querySelector('[data-sidebar]');
  if (hamburger && sidebar) {
    hamburger.addEventListener('click', function () {
      var open = sidebar.getAttribute('data-open') === 'true';
      sidebar.setAttribute('data-open', String(!open));
      hamburger.setAttribute('aria-expanded', String(!open));
    });
    // Close the menu when a nav link is tapped on mobile.
    sidebar.addEventListener('click', function (event) {
      if (event.target && event.target.closest('a')) {
        sidebar.setAttribute('data-open', 'false');
        hamburger.setAttribute('aria-expanded', 'false');
      }
    });
  }
})();
