const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const p0 = new Point(width / 2, height / 2);
const p1 = new Point(0, 0);
const r0 = 100;
const r1 = 60;

render();

canvas.addEventListener("mousemove", onMouseMove);

function onMouseMove(event) {
  p1.x = event.clientX;
  p1.y = event.clientY;
  render();
}

function render() {
  context.clearRect(0, 0, width, height);
  const dist = p1.distance(p0);
  if(dist < r0 + r1) {
    context.fillStyle = "red";
  } else {
    context.fillStyle = "black";
  }

  context.beginPath();
  context.arc(p0.x, p0.y, r0, 0, Math.PI * 2);
  context.fill();

  context.beginPath();
  context.arc(p1.x, p1.y, r1, 0, Math.PI * 2);
  context.fill();
}
