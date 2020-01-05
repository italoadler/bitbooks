float angleX = 0;
float angleY = 0;
float speedX = random(0.1);
float speedY = random(0.1);

void setup() {
  size(800, 600);
  noLoop();
}

void draw() {
  background(255);
  beginShape();
  for (int i = 0; i < 10000; i++) {
    float x = width / 2 + sin(angleX) * width * 0.4;
    float y = height / 2 + sin(angleY) * height * 0.4;
    angleX += speedX;
    angleY += speedY;
    
    vertex(x, y);
  }
  endShape();
}
