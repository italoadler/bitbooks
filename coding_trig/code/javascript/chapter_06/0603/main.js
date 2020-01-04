const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

const p0 = {
  x: width / 2,
  y: height / 2,
};

let p1;

canvas.addEventListener("click", onClick);
context.font = "20px Arial";

function onClick(event) {
  p1 = new Point(event.clientX, event.clientY);
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
  let rad = Math.atan2(y, x);
  let deg = Math.round(rad * 180 / Math.PI);

  context.fillText(deg + " degrees", p0.x, p0.y - 10);

}

function renderPoint(p) {
  context.beginPath();
  context.arc(p.x, p.y, 5, 0, Math.PI * 2);
  context.fill();
}
