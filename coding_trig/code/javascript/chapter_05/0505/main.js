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
  context.clearRect(0, 0, width, height);

  let x = width / 2 + Math.sin(angleX) * width * 0.4;
  let y = height / 2 + Math.sin(angleY) * height * 0.4;

  angleX += speedX;
  angleY += speedY;

  context.beginPath();
  context.arc(x, y, 20, 0, Math.PI * 2);
  context.fill();

  requestAnimationFrame(render);
}
