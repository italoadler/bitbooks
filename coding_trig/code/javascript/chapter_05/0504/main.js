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

  let rotation = 0 + Math.sin(angle) * Math.PI / 2;

  angle += speed;

  context.save();
  context.beginPath();
  context.translate(x, y);
  context.rotate(rotation);
  context.rect(-100, -100, 200, 200);
  context.fill();
  context.restore();

  requestAnimationFrame(render);
}
