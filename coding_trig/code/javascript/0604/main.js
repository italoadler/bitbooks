const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

canvas.addEventListener("mousemove", onMouseMove);

function onMouseMove(event) {
  // determine the angle
  const mouseX = event.clientX;
  const mouseY = event.clientY;
  const dx = mouseX - width / 2;
  const dy = mouseY - height / 2;
  const angle = Math.atan2(dy, dx);

  context.clearRect(0, 0, width, height);

  // transform the context by that angle
  context.save();
  context.translate(width / 2, height / 2);
  context.rotate(angle);

  // draw an arrow
  context.beginPath();
  context.moveTo(-30, -10);
  context.lineTo(5, -10);
  context.lineTo(5, -20);
  context.lineTo(30, 0);
  context.lineTo(5, 20);
  context.lineTo(5, 10);
  context.lineTo(-30, 10);
  context.closePath();
  context.fill();

  // restore the context
  context.restore();
}
