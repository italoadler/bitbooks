const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

context.translate(width / 2, height / 2);

let rx = 400;
let ry = 200;
context.beginPath();
for (let i = 0; i < Math.PI * 2; i += 0.1) {
  let x = Math.cos(i) * rx;
  let y = Math.sin(i) * ry;
  context.lineTo(x, y);
}
context.closePath();
context.stroke();
