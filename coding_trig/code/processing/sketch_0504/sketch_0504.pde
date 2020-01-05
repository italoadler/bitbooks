float angle = 0;
float speed = 0.05;

void setup() {
  size(600, 600);
  fill(0);
  noStroke();
}

void draw() {
  background(255);
  float x = width / 2;
  float y = height / 2;
  float rotation  = 0 + sin(angle) * PI / 2;
  angle += speed;
  
  push();
  translate(width / 2, height / 2);
  rotate(rotation);
  rect(-100, -100, 200, 200);
  pop();
}
