const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

function circle(context, x, y, radius) {
  context.beginPath();
  for (let i = 0; i < Math.PI * 2; i += 0.1) {
    context.lineTo(x + Math.cos(i) * radius, y + Math.sin(i) * radius);
  }
  context.closePath();
  context.stroke();
}

for (let i = 0; i < 50; i++) {
  circle(context, Math.random() * width, Math.random() * height, 10 + Math.random() * 100);
}
