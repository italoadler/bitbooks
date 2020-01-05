float dist = 300;
float speed = 0.05;
float angle = 0;

void setup() {
  size(700, 700);
  fill(0);
}

void draw() {
  background(255);

  float x = cos(angle) * dist;
  float y = sin(angle) * dist;
  circle(width / 2 + x, height / 2 + y, 40);
  angle += speed;
}
