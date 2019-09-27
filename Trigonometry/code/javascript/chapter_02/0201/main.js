const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const points = [];

canvas.addEventListener("click", onClick);

function onClick(event) {
  if (points.length < 2) {
    points.push({ x: event.clientX, y: event.clientY });
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
    context.lineTo(points[1].x, points[0].y);
    context.lineTo(points[0].x, points[0].y);
    context.stroke();

    const a = Math.abs(points[0].x - points[1].x);
    const b = Math.abs(points[0].y - points[1].y);
    const c = Math.sqrt(a * a + b * b);

    context.fillText(a, (points[0].x + points[1].x) / 2, points[0].y - 5);
    context.fillText(b, points[1].x + 5, (points[0].y + points[1].y) / 2);
    context.fillText(c, (points[0].x + points[1].x) / 2 - 10, (points[0].y + points[1].y) / 2 - 10);

  }
}
