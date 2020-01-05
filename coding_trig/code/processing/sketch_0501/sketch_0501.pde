float angle = 0;
float speed = 0.05;

void setup() {
  size(600, 600);
  fill(0);
}

void draw() {
  background(255);
  float x = width / 2;
  float y = height / 2 + sin(angle) * height * 0.4;
  angle += speed;
  
  circle(x, y, 40);
}
