const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

context.translate(0, height / 2);
context.scale(1, -1);

context.beginPath();
for (let i = 0; i <= Math.PI * 2; i += 0.2) {
  let x = i * 100;
  let y = Math.cos(i) * 200;
  context.lineTo(x, y);
}
context.stroke();

context.beginPath();
context.moveTo(0, 0);
context.lineTo(width, 0);
context.stroke();
