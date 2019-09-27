const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const points = [];

canvas.addEventListener("click", onClick);

function onClick(event) {
  if (points.length < 2) {
    points.push(new Point(event.clientX, event.clientY));
  } else {
    points.length = 0;
  }
  render();
}

function render() {
  context.clearRect(0, 0, width, height);
  for(let i = 0; i < points.length; i++) {
    const p = points[i];
    context.beginPath();
    context.arc(p.x, p.y, 5, 0, Math.PI * 2);
    context.fill();
  }
  if (points.length === 2) {
    context.beginPath();
    context.moveTo(points[0].x, points[0].y);
    context.lineTo(points[1].x, points[1].y);
    context.stroke();

    const dist = points[0].distance(points[1]);

    context.fillText(dist, (points[0].x + points[1].x) / 2 - 10, (points[0].y + points[1].y) / 2 - 10);

  }
}
