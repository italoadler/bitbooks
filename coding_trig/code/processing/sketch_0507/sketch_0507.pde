float radius = 200;

void setup() {
  size(600, 600);
  noLoop();
}

void draw() {
  background(255);
  translate(width / 2, height / 2);
  
  beginShape();
  for (float i = 0; i < PI * 2; i += 0.1) {
    float x = cos(i) * radius;
    float y = sin(i) * radius; 
    vertex(x, y);
  }
  endShape(CLOSE);
}
