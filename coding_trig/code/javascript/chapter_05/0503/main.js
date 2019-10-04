const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

let angle = 0;
let speed = 0.05;

render();

function render() {
  context.clearRect(0, 0, width, height);

  let x = width / 2;
  let y = height / 2;

  context.globalAlpha = 0.5 + Math.sin(angle) * 0.5;

  angle += speed;

  context.beginPath();
  context.arc(x, y, 200, 0, Math.PI * 2);
  context.fill();
  requestAnimationFrame(render);
}
