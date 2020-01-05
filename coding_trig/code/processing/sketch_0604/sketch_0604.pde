void setup() {
  size(600, 600);
  fill(0);
}

void draw() {
  background(255);
  
  float dx = mouseX - width / 2;
  float dy = mouseY - height / 2;
  float angle = atan2(dy, dx);
  
  translate(width / 2, height / 2);
  rotate(angle);
  
  beginShape();
  vertex(-30, -10);
  vertex(5, -10);
  vertex(5, -20);
  vertex(30, 0);
  vertex(5, 20);
  vertex(5, 10);
  vertex(-30, 10);
  endShape();
}
