float radius = 200;

void setup() {
  size(600, 600);
  noLoop();
  noFill();
}

void draw() {
  background(255);
  
  for (int i = 0; i < 50; i++) {
    drawCircle(random(width), random(height), random(10, 100));
  }  
}

// note, processing already has a circle function, but if we were going to make one,
// this is what it would look like. Also, this one uses radius, rather than the 
// built in one which uses a diameter.
void drawCircle(float x, float y, float radius) {
  beginShape();
  for (float i = 0; i < PI * 2; i += 0.1) {
    vertex(x + cos(i) * radius, y + sin(i) * radius);
  }
  endShape(CLOSE);
}
