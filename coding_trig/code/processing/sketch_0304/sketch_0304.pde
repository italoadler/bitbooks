Point p0;
Point p1;
int r0 = 100;
int r1 = 60;

void setup() {
  size(600, 600);
  noStroke();
  p0 = new Point(width / 2, height / 2);
  p1 = new Point(0, 0);
}

void draw() {
  background(255);
  float dist = p1.distance(p0);
  if (dist < r0 + r1) {
    fill(255, 0, 0);
  } else {
    fill(0);
  }
  circle(p0.x, p0.y, r0 * 2); // third param is diameter, not radius, in processing
  circle(p1.x, p1.y, r1 * 2); // third param is diameter, not radius, in processing
}

void mouseMoved() {
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
