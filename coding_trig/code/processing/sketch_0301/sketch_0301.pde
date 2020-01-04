int[] xvals;
int[] yvals;
int numPoints = 0;

void setup() {
  size(600, 600);
  fill(0);
  xvals = new int[3];
  yvals = new int[3];
}

void draw() {
  background(255);
  for (int i = 0; i < numPoints; i++) {
    circle(xvals[i], yvals[i], 5);
  }
  if (numPoints == 2) {
    line(xvals[0], yvals[0], xvals[1], yvals[1]);
    line(xvals[1], yvals[1], xvals[1], yvals[0]);
    line(xvals[1], yvals[0], xvals[0], yvals[0]);
  
    float a = abs(xvals[0] - xvals[1]);
    float b = abs(yvals[0] - yvals[1]);
    float c = sqrt(a * a + b * b);
    
    text(a, (xvals[0] + xvals[1]) / 2, yvals[0] - 5);
    text(b, xvals[1] + 5, (yvals[0] + yvals[1]) / 2 );
    text(c, (xvals[0] +xvals[1]) / 2 - 10, (yvals[0] + yvals[1]) / 2 - 10);
  }
}

void mouseClicked() {
  xvals[numPoints] = mouseX;
  yvals[numPoints] = mouseY;
  numPoints++;
  if (numPoints > 2) {
    numPoints = 0;
  }
}
