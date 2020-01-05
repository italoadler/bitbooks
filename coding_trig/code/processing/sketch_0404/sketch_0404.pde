void setup() {
  size(628, 600);
}

void draw() {
  background(255);
  translate(0, height / 2);
  scale(1, -1);
  
  beginShape();
  for (float i = 0; i < PI * 10; i += 0.2) {
    float x = i * 20;
    float y = sin(i) * 20;
    vertex(x, y);
  }
  endShape();
  
  line(0, 0, width, 0);
}
