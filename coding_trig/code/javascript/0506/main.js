const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

let angleX = 0;
let angleY = 0;
let speedX = Math.random() * 0.1;
let speedY = Math.random() * 0.1;

render();

function render() {
  context.beginPath();

  for (let i = 0; i < 10000; i++) {
    let x = width / 2 + Math.sin(angleX) * width * 0.4;
    let y = height / 2 + Math.sin(angleY) * height * 0.4;

    angleX += speedX;
    angleY += speedY;

    context.lineTo(x, y);
  }

  context.stroke();
}
