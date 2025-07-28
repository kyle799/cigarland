const themeSwitch   = document.getElementById('themeSwitch');
const prefersDarkMq = window.matchMedia('(prefers-color-scheme: dark)');
function applyTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  localStorage.setItem('theme', theme);
  themeSwitch.checked = (theme === 'dark');
}
(function initTheme() {
  const stored = localStorage.getItem('theme');
  const initial = stored || (prefersDarkMq.matches ? 'dark' : 'light');
  applyTheme(initial);
})();
themeSwitch.addEventListener('change', () => {
  applyTheme(themeSwitch.checked ? 'dark' : 'light');
});
prefersDarkMq.addEventListener('change', e => {
  if (!localStorage.getItem('theme')) applyTheme(e.matches ? 'dark' : 'light');
});
document.querySelectorAll('.spoiler-toggle').forEach(btn => {
  btn.addEventListener('click', () => {
    const box = document.getElementById(btn.dataset.target);
    const isHidden = box.classList.toggle('hidden');
    if (!isHidden) {
      box.classList.add('fade-in');
      setTimeout(() => box.classList.remove('fade-in'), 250);
    }
    btn.textContent = isHidden ? 'Show notes' : 'Hide notes';
  });
});