(function () {
  const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
  const chars = '01·:'.split('');
  const bayer4 = [
    0, 8, 2, 10,
    12, 4, 14, 6,
    3, 11, 1, 9,
    15, 7, 13, 5,
  ];

  function hash(x, y, t) {
    const n = Math.sin(x * 12.9898 + y * 78.233 + t * 37.719) * 43758.5453;
    return n - Math.floor(n);
  }

  function initImage(img) {
    if (img.dataset.matrixReady) return;
    img.dataset.matrixReady = '1';

    function setup() {
      const canvas = document.createElement('canvas');
      canvas.className = img.className.replace('matrix-dither', '').trim();
      canvas.style.cssText = img.style.cssText;
      canvas.style.aspectRatio = `${img.naturalWidth} / ${img.naturalHeight}`;
      canvas.setAttribute('aria-label', img.alt || '');
      canvas.setAttribute('role', 'img');

      const ctx = canvas.getContext('2d', { alpha: true });
      if (!ctx) {
        img.style.visibility = 'visible';
        return;
      }

      const sample = document.createElement('canvas');
      const sampleCtx = sample.getContext('2d', { willReadFrequently: true });
      const cell = Number(img.dataset.cell || 3);
      let raf = null;
      let visible = true;
      let needsResize = true;
      let last = 0;

      function resize() {
        const rect = canvas.getBoundingClientRect();
        const dpr = Math.min(window.devicePixelRatio || 1, 2);
        canvas.width = Math.max(1, Math.floor(rect.width * dpr));
        canvas.height = Math.max(1, Math.floor(rect.height * dpr));
        sample.width = Math.max(1, Math.floor(canvas.width / cell));
        sample.height = Math.max(1, Math.floor(canvas.height / cell));
        needsResize = false;
      }

      function draw(ts) {
        raf = null;
        if (!canvas.isConnected || !visible) return;
        if (!reducedMotion && ts - last < 110) {
          raf = requestAnimationFrame(draw);
          return;
        }
        last = ts;
        if (needsResize) resize();

        sampleCtx.drawImage(img, 0, 0, sample.width, sample.height);
        const data = sampleCtx.getImageData(0, 0, sample.width, sample.height).data;
        const tick = Math.floor(ts / 450);

        ctx.clearRect(0, 0, canvas.width, canvas.height);
        ctx.fillStyle = '#080808';
        ctx.fillRect(0, 0, canvas.width, canvas.height);
        ctx.font = `${Math.max(4, Math.floor(cell * 1.45))}px monospace`;
        ctx.textAlign = 'center';
        ctx.textBaseline = 'middle';

        for (let y = 0; y < sample.height; y++) {
          for (let x = 0; x < sample.width; x++) {
            const i = (y * sample.width + x) * 4;
            const r = data[i];
            const g = data[i + 1];
            const b = data[i + 2];
            let gray = (r * 0.299 + g * 0.587 + b * 0.114) / 255;

            // Small contrast lift, but keep more midtones than the first version.
            gray = Math.max(0, Math.min(1, (gray - 0.08) * 1.55));

            const ordered = (bayer4[(y % 4) * 4 + (x % 4)] + 0.5) / 16;
            const n = hash(x, y, tick);
            const threshold = ordered * 0.62 + n * 0.13;
            if (gray <= threshold) continue;

            const amount = Math.min(1, (gray - threshold) / (1 - threshold));
            const flicker = reducedMotion ? 1 : 0.88 + hash(x + 19, y + 7, tick) * 0.18;
            const tone = Math.floor(78 + amount * 172);
            const alpha = Math.min(0.92, 0.14 + amount * 0.74) * flicker;
            ctx.fillStyle = `rgba(${tone}, ${tone}, ${tone}, ${alpha})`;

            const px = x * cell;
            const py = y * cell;
            if (gray > 0.7 && hash(x + 3, y + 11, tick) > 0.985) {
              ctx.fillText(chars[(x + y + tick) % chars.length], px + cell / 2, py + cell / 2);
            } else {
              const size = Math.max(1, cell * (0.35 + amount * 0.35));
              ctx.fillRect(px + (cell - size) / 2, py + (cell - size) / 2, size, size);
            }
          }
        }

        if (!reducedMotion) raf = requestAnimationFrame(draw);
      }

      function requestDraw() {
        if (!raf) raf = requestAnimationFrame(draw);
      }

      img.parentNode.replaceChild(canvas, img);
      requestDraw();

      window.addEventListener('resize', () => {
        needsResize = true;
        requestDraw();
      });

      const observer = new IntersectionObserver((entries) => {
        visible = entries[0].isIntersecting;
        if (visible) requestDraw();
        else if (raf) {
          cancelAnimationFrame(raf);
          raf = null;
        }
      });
      observer.observe(canvas);
    }

    if (img.complete && img.naturalWidth > 0) setup();
    else img.addEventListener('load', setup, { once: true });
  }

  function init(root) {
    (root || document).querySelectorAll('img.matrix-dither').forEach(initImage);
  }

  if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', () => init(document));
  else init(document);
})();
