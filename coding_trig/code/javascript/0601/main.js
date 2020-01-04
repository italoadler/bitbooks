const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

context.translate(0, height / 2);
context.scale(1, -1); // to flip the sine wave to Cartesian coords.

context.beginPath();
for (let i = 0; i <= Math.PI * 10; i += 0.01) {
  let x = i * 100;
  let y = Math.tan(i) * 20;
  context.lineTo(x, y);
}
context.stroke();

context.lineWidth = 0.25;
context.beginPath();
context.moveTo(0, 0);
context.lineTo(width, 0);
context.stroke();
