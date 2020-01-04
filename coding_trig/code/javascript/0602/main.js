const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const p0 = new Point(width / 2, height / 2);
const p1 = new Point(0, 0);

canvas.addEventListener("click", onClick);
context.font = "20px Arial";

function onClick(event) {
  p1.x = event.clientX;
  p1.y = event.clientY;
  render();
}

function render() {
  context.clearRect(0, 0, width, height);
  renderPoint(p0);
  renderPoint(p1);

  context.beginPath();
  context.moveTo(p0.x, p0.y);
  context.lineTo(p1.x, p0.y);
  context.lineTo(p1.x, p1.y);
  context.lineTo(p0.x, p0.y);
  context.stroke();

  let x = p1.x - p0.x;
  let y = p1.y - p0.y;
  let rad = Math.atan(y / x);
  let deg = Math.round(rad * 180 / Math.PI);

  context.fillText(deg + " degrees", p0.x, p0.y - 10);

}

function renderPoint(p) {
  context.beginPath();
  context.arc(p.x, p.y, 5, 0, Math.PI * 2);
  context.fill();
}
