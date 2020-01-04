class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  distance(p) {
    const dx = this.x - p.x;
    const dy = this.y - p.y;
    return Math.sqrt(dx * dx + dy * dy);
  }

  render(context) {
    context.save();
    context.fillStyle = "#000000";
    context.beginPath();
    context.arc(this.x, this.y, 5, 0, Math.PI * 2);
    context.fill();
    context.restore();
  }
}
