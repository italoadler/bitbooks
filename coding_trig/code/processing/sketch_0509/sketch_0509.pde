float radius = 200;

void setup() {
  size(700, 700);
  noLoop();
  noFill();
}

void draw() {
  background(255);
  translate(width / 2, height / 2);
  
  float num = 20;
  float dist = 200;
  float inc = PI * 2 / num;
  
  for (float i = 0; i < PI * 2; i += inc) {
    float x = cos(i) * dist;
    float y = sin(i) * dist;
    circle(x, y, 40);
  }
  
  num = 30;
  dist = 300;
  
  for (float i = 0; i < PI * 2; i += inc) {
    float x = cos(i) * dist;
    float y = sin(i) * dist;
    rect(x - 20, y - 20, 40, 40);
  }
}
