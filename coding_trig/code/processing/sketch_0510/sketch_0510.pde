float radius = 200;

void setup() {
  size(800, 800);
  noLoop();
  textSize(32);
  fill(0);
}

void draw() {
  background(255);
  translate(width / 2, height / 2);
  
  float num = 26;
  float dist = 200;
  float inc = PI * 2 / num;
  
  for (float i = 0; i < PI * 2; i += inc) {
    float x = cos(i) * dist;
    float y = sin(i) * dist;
    push();
    translate(x, y);
    rotate(i);
    text("Hello, world", 0, 0);
    pop();
  }
  
  //without trig
  //for (float i = 0; i < PI * 2; i += inc) {
  //  push();
  //  rotate(i);
  //  text("Hello, world", dist, 0);
  //  pop();
  //}
}
