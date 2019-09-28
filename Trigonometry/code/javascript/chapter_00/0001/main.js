const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

// sample drawing code:
for (let i = 0; i < 100; i++) {
  context.lineTo(Math.random() * width, Math.random() * height);
}
context.stroke();
