float angle = 0;
float speed = 0.05;

void setup() {
  size(600, 600);
  fill(0);
}

void draw() {
  background(255);
  float x = width / 2;
  float y = height / 2;
  float radius = 100 + sin(angle) * 100;
  angle += speed;
  
  circle(x, y, radius * 2); // third arg is diameter in processing
}
