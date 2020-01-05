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
  float alpha = 128 + sin(angle) * 127;
  angle += speed;
  
  fill(0, 0, 0, alpha);
  circle(x, y, 400);
}
