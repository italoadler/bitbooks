float angleX = 0;
float angleY = 0;
float speedX = random(0.1);
float speedY = random(0.1);

void setup() {
  size(600, 600);
  fill(0);
  noStroke();
}

void draw() {
  background(255);
  float x = width / 2 + sin(angleX) * width * 0.4;
  float y = height / 2 + sin(angleY) * height * 0.4;
  angleX += speedX;
  angleY += speedY;
  
  push();
  circle(x, y, 40);
  pop();
}
