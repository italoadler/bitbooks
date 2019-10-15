Point[] points;
int numPoints = 0;

void setup() {
  size(600, 600);
  fill(0);
  points = new Point[3];
}

void draw() {
  background(255);
  for (int i = 0; i < numPoints; i++) {
    Point p = points[i];
    circle(p.x, p.y, 5);
  }
  if (numPoints == 2) {
    line(points[0].x, points[0].y, points[1].x, points[1].y);
    float dist = points[0].distance(points[1]);
    text(dist, (points[0].x + points[1].x) / 2 - 10, (points[0].y + points[1].y) / 2 - 10);
  }
}

void mouseClicked() {
  points[numPoints] = new Point(mouseX, mouseY);
  numPoints++;
  if (numPoints > 2) {
    numPoints = 0;
  }
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
