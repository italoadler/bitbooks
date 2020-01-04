const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

let dist = 300;
let speed = 0.05;
let angle = 0;

render();

function render() {
  context.clearRect(0, 0, width, height);
  let x = Math.cos(angle) * dist;
  let y = Math.sin(angle) * dist;
  context.beginPath();
  context.arc(width / 2 + x, height / 2 + y, 20, 0, Math.PI * 2);
  context.fill();
  angle += speed;
  requestAnimationFrame(render);
}

