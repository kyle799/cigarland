(function () {
  const canvas = document.getElementById('bg-canvas');
  const ctx    = canvas.getContext('2d');
  let W, H, nodes;

  function resize() {
    W = canvas.width  = window.innerWidth;
    H = canvas.height = window.innerHeight;
  }

  function mkNodes(n) {
    return Array.from({ length: n }, () => ({
      x:  Math.random() * W,
      y:  Math.random() * H,
      vx: (Math.random() - 0.5) * 0.28,
      vy: (Math.random() - 0.5) * 0.28,
      r:  Math.random() * 1.4 + 0.4,
    }));
  }

  function tick() {
    ctx.clearRect(0, 0, W, H);
    const D = 155;

    for (let i = 0; i < nodes.length; i++) {
      const a = nodes[i];
      a.x += a.vx; a.y += a.vy;
      if (a.x < 0 || a.x > W) a.vx *= -1;
      if (a.y < 0 || a.y > H) a.vy *= -1;

      for (let j = i + 1; j < nodes.length; j++) {
        const b  = nodes[j];
        const dx = a.x - b.x, dy = a.y - b.y;
        const d  = Math.sqrt(dx * dx + dy * dy);
        if (d < D) {
          ctx.beginPath();
          ctx.moveTo(a.x, a.y);
          ctx.lineTo(b.x, b.y);
          ctx.strokeStyle = `rgba(138,181,200,${(1 - d / D) * 0.14})`;
          ctx.lineWidth   = 0.5;
          ctx.stroke();
        }
      }

      ctx.beginPath();
      ctx.arc(a.x, a.y, a.r, 0, Math.PI * 2);
      ctx.fillStyle = 'rgba(138,181,200,0.38)';
      ctx.fill();
    }

    requestAnimationFrame(tick);
  }

  resize();
  nodes = mkNodes(70);
  tick();
  window.addEventListener('resize', resize);
})();
