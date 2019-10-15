void setup() {
  size(600, 600);
  noLoop();
}

void draw() {
  float x0 = 0;
  float y0 = 0;
  for (int i = 0; i < 100; i++) {
    float x1 = random(width);
    float y1 = random(height);
    line(x0, y0, x1, y1);
    x0 = x1;
    y0 = y1;
  }
}
