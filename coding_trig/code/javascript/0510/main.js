const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

context.translate(width / 2, height / 2);

let num = 26;
let dist = 200;
let inc = Math.PI * 2 / num;

context.font = "40px Arial";

for (let i = 0; i < Math.PI * 2; i += inc) {
  let x = Math.cos(i) * dist;
  let y = Math.sin(i) * dist;
  context.save();
  context.translate(x, y);
  context.rotate(i);
  context.fillText("Hello, world", 0, 0);
  context.restore();
}

// without trig:
// for (let i = 0; i < Math.PI * 2; i += inc) {
//   context.save();
//   context.rotate(i);
//   context.fillText("Hello, world", dist, 0);
//   context.restore();
// }

