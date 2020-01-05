Point p0;
Point p1;

void setup() {
  size(600, 600);
  fill(0);
  p0 = new Point(width / 2, height / 2);
  p1 = new Point(0, 0);
}

void draw() {
  background(255);
  circle(p0.x, p0.y, 10);
  circle(p1.x, p1.y, 10);
  
  noFill();
  beginShape();
  vertex(p0.x, p0.y);
  vertex(p1.x, p0.y);
  vertex(p1.x, p1.y);
  vertex(p0.x, p0.y);
  endShape();
  
  float x = p1.x - p0.x;
  float y = p1.y - p0.y;
  float rad = atan(y / x);
  float deg = round(rad * 180 / PI);
  
  fill(0);
  text(deg + " degrees", p0.x, p0.y - 10);
}

void mouseClicked() {
  p1.x = mouseX;
  p1.y = mouseY;
}



class Point {
  float x, y;
  
  Point(float x, float y) {
    this.x = x;
    this.y = y;
  }
  
  float distance(Point p) {
    float dx = this.x - p.x;
    float dy = this.y - p.y;
    return sqrt(dx * dx + dy * dy);
  }
}
