void setup() {
  size(900, 700);
  noFill();
  noLoop();
}

void draw() {
  background(255);
  translate(width / 2, height / 2);
  
  float rx = 400;
  float ry = 200;

  beginShape();
  for (float i = 0; i < PI * 2; i += 0.1) {
    float x = cos(i) * rx;
    float y = sin(i) * ry;
    vertex(x, y);
  }
  endShape(CLOSE);
}
