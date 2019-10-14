const canvas = document.getElementById("canvas");
const context = canvas.getContext("2d");
const width = canvas.width = window.innerWidth;
const height = canvas.height = window.innerHeight;

context.translate(width / 2, height / 2);

let num = 20;
let dist = 200;
let inc = Math.PI * 2 / num;

for (let i = 0; i < Math.PI * 2; i += inc) {
  let x = Math.cos(i) * dist;
  let y = Math.sin(i) * dist;
  context.beginPath();
  context.arc(x, y, 20, 0, Math.PI * 2);
  context.stroke();
}

num = 30;
dist = 300;
for (let i = 0; i < Math.PI * 2; i += inc) {
  let x = Math.cos(i) * dist;
  let y = Math.sin(i) * dist;
  context.beginPath();
  context.rect(x - 20, y - 20, 40, 40);
  context.stroke();
}
