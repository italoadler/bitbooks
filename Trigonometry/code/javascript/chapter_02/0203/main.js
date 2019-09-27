const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const p0 = new Point(width / 2, height / 2);
const p1 = new Point(0, 0);
const radius = 100;
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
  if(dist < radius) {
    context.fillStyle = "red";
    canvas.style.cursor = "pointer";
  } else {
    context.fillStyle = "black";
    canvas.style.cursor = "default";
  }

  context.beginPath();
  context.arc(p0.x, p0.y, radius, 0, Math.PI * 2);
  context.fill();
}
