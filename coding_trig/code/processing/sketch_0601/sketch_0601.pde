void setup() {
  size(1000, 600);
}

void draw() {
  background(255);
  translate(0, height / 2);
  scale(1, -1);
  
  beginShape();
  for (float i = 0; i < PI * 3; i += 0.01) {
    float x = i * 100;
    float y = tan(i) * 20;
    vertex(x, y);
  }
  endShape();
  
  line(0, 0, width, 0);
}
