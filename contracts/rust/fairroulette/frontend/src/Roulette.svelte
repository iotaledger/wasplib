<script lang="ts">
  import { onMount } from 'svelte';

  export let width;
  export let height;
  export let spin;
  export let numbers = 6;

  $: handleSpinSubscription(spin);

  let canvasElement: HTMLElement;
  let canvas: CanvasRenderingContext2D;
  let rotation = 180;
  let rotationHandle;

  function handleSpinSubscription(spin) {
    if (spin) {
      startRotate();
    } else {
      stopRotate();
    }
  }

  function getIndexColor(index: number) {
    if (index % 2 == 1) {
      return { arcColor: '#262e45', fontColor: '#e6e6e7' };
    } else {
      return { arcColor: '#e6e6e7', fontColor: '#262e45' };
    }
  }

  function initializeRouletteWheel() {
    canvas = (<HTMLCanvasElement>canvasElement).getContext('2d');

    canvas.strokeStyle = 'black';
    canvas.lineWidth = 1;
    canvas.imageSmoothingEnabled = true;
    canvas.imageSmoothingQuality = 'high';
    canvas.font = 'bold 42px Arial, sans-serif';
  }

  function drawRouletteWheel() {
    const outsideRadius = Math.round(width / 2);
    const insideRadius = Math.round(outsideRadius * 0.75);
    const fontRadius = Math.round(outsideRadius * 0.8);

    const arc = Math.PI / (numbers / 2);

    const x = Math.round(width / 2);
    const y = Math.round(height / 2);

    let arcNumber = 1;

    for (let i = 0; i < numbers; i++) {
      const angle = rotation + i * arc;
      const color = getIndexColor(i + 1);

      canvas.fillStyle = color.arcColor;

      canvas.beginPath();
      canvas.arc(x, y, outsideRadius, angle, angle + arc, false);
      canvas.arc(x, y, insideRadius, angle + arc, angle, true);
      canvas.stroke();
      canvas.fill();
      canvas.save();

      canvas.translate(
        x + Math.cos(angle + arc / 2) * fontRadius,
        y + Math.sin(angle + arc / 2) * fontRadius
      );

      canvas.rotate(angle + arc / 2 + Math.PI / 2);

      canvas.fillStyle = color.fontColor;

      canvas.fillText(arcNumber.toString(), -canvas.measureText(arcNumber.toString()).width / 2, 0);
      canvas.restore();

      arcNumber++;
    }
  }

  window.addEventListener('resize', drawRouletteWheel);

  onMount(() => {
    initializeRouletteWheel();
    drawRouletteWheel();
  });

  function startRotate() {
    if (!rotationHandle) {
      rotationHandle = setInterval(() => {
        rotation += 0.15;
        drawRouletteWheel();
      }, 15);
    }
  }

  function stopRotate() {
    rotationHandle = clearInterval(rotationHandle);

    rotation = 180;
  }
</script>

<canvas bind:this={canvasElement} {width} {height} />
