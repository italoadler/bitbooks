# Chapter 6 - Tangent and Arc Tangent

In the previous two chapters, we've discussed two of the three core trigonometric functions - sine and cosine. We've seen some useful techniques that implement these functions and how closely these two functions relate to each other. But what about the third function, tangent? Is it that different? Does it have no practical uses?

Well, yes, it's quite different from both sine and cosine in the it works and what it is used for. But tangent does make some extremely useful techniques possible. We'll discover some of them in this chapter.

We'll also look into the arc tangent function, which is the inverse of tangent. Don't worry about what that means exactly just yet. We'll get there shortly.

## The basics of tangent 

As we covered earlier, tangent is just a name for the ratio of the side opposite an angle, to the side adjacent to that angle (in a right triangle). You know well by now that sine and cosine are the ratio of the adjacent and opposite sides to the hypotenuse, which is why they wind up being so similar - they both have that hypotenuse in common. But tangent ignores that side altogether. This results in it being far more different than you might imagine. Let's walk through the tangents of an angle as that angle goes through 360 degrees, like we did for sine and cosine.

We start with an angle close to zero, as seen in figure 6-1.

![An angle close to zero.](images/figure_6-1.png)
*Figure 6-1. An angle close to zero.*

If tangent is the opposite over adjacent, you can imagine that if the angle went all the way down to zero, then the length of the opposite side would be zero. So no matter what the length of the adjacent side was, the tangent of zero will be zero.

As the angle increases, the opposite side gets longer and the adjacent side gets shorter. When we hit 45 degrees, the opposite and adjacent sides happen to be the same length, so the tangent of 45 degrees is one. See figure 6-2.

![A 45-degree angle.](images/figure_6-2.png)
*Figure 6-2. A 45-degree angle.*

At this point, you might think that we're heading into another sine-like curve. But hold that thought. As that angle continues to increase, the opposite side continues to get longer and the adjacent continues to shrink. As you can see in figure 6-3, the ratio of these two starts to become much higher than one.

![An angle close to 90 degrees.](images/figure_6-3.png)
*Figure 6-3. An angle close to 90 degrees.*

In fact, as the angle gets very close to 90 degrees, the length of the adjacent side becomes infinitesimal. This means that the tangent approaches infinity. And when the angle reaches exactly 90 degrees, the length of the adjacent side becomes exactly zero. Now, no matter what the length of the opposite side is, we wind up in a divide-by-zero situation, which is undefined. And so it is that you can't calculate the tangent of 90 degrees (or π/2 radians). Try it in a physical calculator and it should give you an error.

If you use a web based calculator app, however, you might get some large number like `1.6331239e+16` (that's scientific notation, equal to `1.6331239 * 10^16`).

That's because web based calculators are done in JavaScript. Try this in JavaScript:

    Math.tan(Math.PI / 2);

In Chrome, this gives me `16331239353195370`. You might get different results in different languages and platforms. Some might give you an error. Some might give you a large number like this. My best guess is that because π/2 is an irrational number, it cannot be accurately held in a JavaScript variable. So you are getting the tangent of a value close to, but not exactly π/2. The adjacent side will be very tiny, but not exactly zero, so you get a large number for an answer, not undefined.

Continuing past 90 degrees, the adjacent side becomes negative. When it's just barely over 90 degrees, the tangent will be approaching negative infinity. At 135 degrees, the opposite and adjacent sides will be the same length, but because the adjacent side is in the negative realm, the tangent of 135 degrees is -1.

Continuing on towards 180 degrees, the opposite side goes towards zero, as does the tangent of the angle. See figure 6-4.

![Approaching 180 degrees.](images/figure_6-4.png)
*Figure 6-4. Approaching 180 degrees.*

Beyond 180 degrees, both sides are negative, so the ratio goes positive again, to approach infinity again, and then undefined at 270 degrees. In the final quadrant, the adjacent side is positive and the opposite is negative, so we start again at near negative infinity and go back up to zero at 360 degrees. See figure 6-5.

![Angles in the third and fourth quadrants.](images/figure_6-5.png)
*Figure 6-5. Angles in the third and fourth quadrants.*

Let's plot it out and see what kind of curve this creates. See listing 6-1:

    const canvas = document.getElementById("canvas");
    const context = canvas.getContext("2d");
    const width = canvas.width = window.innerWidth;
    const height = canvas.height = window.innerHeight;

    context.translate(0, height / 2);
    context.scale(1, -1); // to flip the sine wave to Cartesian coords.

    context.beginPath();
    for (let i = 0; i <= Math.PI * 10; i += 0.01) {
      let x = i * 100;
      let y = Math.tan(i) * 20;
      context.lineTo(x, y);
    }
    context.stroke();

    context.lineWidth = 0.25;
    context.beginPath();
    context.moveTo(0, 0);
    context.lineTo(width, 0);
    context.stroke();

*Listing 6-1*

This is very similar to listing 4-4 in chapter 4, other than the fact that it uses `Math.tan` rather than `Math.sin` and a few of the values have been changed to show the curve better. Thus, I won't go through the details of the code here. What it gives you is shown in figure 6-6.

![Tangent plotted out.](images/figure_6-6.png)
*Figure 6-6. Tangent plotted out.*

As described, the curve starts at zero for zero degrees. It shoots up to infinity as it nears 90 degrees (π/2 radians) and then jump back down towards negative infinity just past 90 degrees. Because we are drawing a single continuous series of line segments to create the curve, we see a nearly vertical line connecting the high positive and high negative values. Most times you see tangent plotted out, that vertical line will either be shown as a dotted line or not shown at all. The fact that it's drawn here is just a byproduct of our drawing routine. In reality, the value of of tangent will never fall anywhere on that vertical line. It approaches infinity, becomes undefined and immediately jumps to near negative infinity.

## Uses of tangent

I'll be straight with you. For the most part, tangent is used far less than sine or cosine. I could dive into some contrived examples of where you might use tangent, but I'd rather stick to things you'll actually need to use on a regular basis. 

But I will give a very quick conceptual idea. Say you have a right triangle. You know one angle and the length of the adjacent side of that angle. And you want to know the length of the opposite side, as seen in figure 6-7. 

![Use of tangent.](images/figure_6-7.png)
*Figure 6-7. Use of tangent.*

We'll call the opposite side y and the adjacent side x. We know that the angle A is 30 degrees and x is equal to 10. Well, the formula would be:

    tan A = y / x

because the tangent is the opposite over adjacent. Substituting the values we know, we get:

    tan 30 = y / 10

Then we can put the value we don't know on one side:

    y = (tan 30) * 10

We can use a calculator to discover that the tangent of 30 degrees is about 0.577. So we get:

    y = 0.577 * 10

So the length of the opposite side is roughly 5.77.

Now, if you run into something like that, which you just might from time to time, you'll know what to do. But for now, I'm going to go into a new concept that will give us some extremely handy uses for tangent - or at least one of its close relatives.

## Arc functions (inverse trigonometric functions)

Thus far, I've introduced the three main trigonometric functions: sine, cosine and tangent. But there are several other functions that fall under the subject of trig. The first of these are known sometimes as "arc functions" or "inverse trigonometric functions. Their names are arc sine, arc cosine and arc tangent.

The inverse part makes a lot of sense because they are the mathematical inverse of the functions they are named after. For example, we know that the sine of an angle is equal to the ratio of its opposite side and the hypotenuse.  But the arc sine function takes the ratio of an opposite side and a hypotenuse and returns the angle relating to those to lengths.

Say you have the situation shown in figure 6-8:

![Arc sine.](images/figure_6-8.png)
*Figure 6-8. Arc sine.*

We have the length of the hypotenuse (10) and the opposite side (5). And we want to know what that angle is. Well, the ratio of the opposite and the hypotenuse is 0.5. And the arc sine of 0.5 is 30 degrees.

Or, to be even more concise:

The sine of 30 degrees is 0.5. 

The arc sine of 0.5 is 30 degrees.

Now you see why they are inverse functions. Sine, cosine and tangent take an angle and give you a ratio. The arc versions take a ratio and give you back the angle.

You might see inverse functions written in a couple of different ways. One is just spelling out the arc: `arc sin`. Or sometimes without the space: `arcsin`. Another way is with a superscripted -1 after the name of the function, `sin`<sup>`-1`</sup>.

Of course, this is the same with `arc cos`, `arccos`, `cos`<sup>`-1`</sup>, `arc tan`, `arctan` and `tan`<sup>`-1`</sup>

And in code, these are generally written as `asin`, `acos` and `atan`. So in JavaScript we can use the `Math` functions, `Math.asin`, `Math.acos` and `Math.atan` (and `Math.atan2` which I'll cover shortly).

## Arc tangent

When we enter the realm of arc functions, the importances reverse. Here, you'll find the occasional use for arc sine and arc cosine, but arc tangent is a real powerhouse.

A very common use case is, given two x, y positions, figure out the angle between them. Take figure 6-9 for example:

![Angle between two points](images/figure_6-9.png)
*Figure 6-9. Angle between two points.*

We have points A and B. If we subtract point A's x position from point B's x position, and the same on the y-axis, we can create a right triangle. This is very much the same thing we did in chapter 3 when we wanted to get the distance between two points. But this time, instead of using the Pythagorean theorem, we can pass the ratio of these two sides to the `Math.atan` function. This will give us the angle at point A, in radians.

Let's dive right into some code. See listing 6-2.

    const canvas = document.getElementById("canvas");
    const context = canvas.getContext("2d");
    const width = canvas.width = window.innerWidth;
    const height = canvas.height = window.innerHeight;

    const p0 = new Point(width / 2, height / 2);
    const p1 = new Point(0, 0);

    canvas.addEventListener("click", onClick);
    context.font = "20px Arial";

    function onClick(event) {
      p1.x = event.clientX;
      p1.y = event.clientY;
      render();
    }

    function render() {
      context.clearRect(0, 0, width, height);
      renderPoint(p0);
      renderPoint(p1);

      context.beginPath();
      context.moveTo(p0.x, p0.y);
      context.lineTo(p1.x, p0.y);
      context.lineTo(p1.x, p1.y);
      context.lineTo(p0.x, p0.y);
      context.stroke();

      let x = p1.x - p0.x;
      let y = p1.y - p0.y;
      let rad = Math.atan(y / x);
      let deg = Math.round(rad * 180 / Math.PI);

      context.fillText(deg + " degrees", p0.x, p0.y - 10);

    }

    function renderPoint(p) {
      context.beginPath();
      context.arc(p.x, p.y, 5, 0, Math.PI * 2);
      context.fill();
    }

*Listing 6-2*

This is pretty similar to listing 3-2 from chapter 3. In this case, `p0` is a point at the center of the canvas and `p1` is created when you click the mouse. The `render` function clears the screen and renders both points, then draws the right triangle created by those two points. It gets `x` and `y` values by subtracting the positions of the two points on the x- and y-axes. It then calls `Math.atan(y / x)` to get the angle between the points in radians, converts it to degrees, rounds it off, and writes that value on the screen.

When you run this and click above and to the right of the center of the screen, you should get something like you see in figure 6-10.

![Using arc tangent to find an angle](images/figure_6-10.png)<br/>
*Figure 6-10. Using arc tangent to find an angle.*

This shows that the angle formed between the two points is approximately -30 degrees. (-30 because it's going counter-clockwise and we're dealing with screen coordinates, not Cartesian coordinates.)

But what if we click below and to the left of center, like in figure 6-11?

![Another angle with arc tangent](images/figure_6-11.png)<br/>
*Figure 6-11. Another angle with arc tangent.*

Uh-oh, this also shows -30 degrees. Seems like a bug! But it's not. The problem is that the `Math.atan` function just takes a ratio. In the first case we have a positive x value and a negative y value, so we get a negative ratio. In the second case, we have a negative x and a positive y, but it ends up being the same ratio! There is no way that the function can know which quadrant you are referring to. So it will always give you values between -π/2 and +π/2 radians (-90 and +90 degrees). But you can see that this ambiguity is less than optimal if you are trying to find an angle that really is in one of those left-hand quadrants.

Thankfully, there's another function, `Math.atan2`. Not the most creative name in the world, but it's become the standard in most math libraries, so get used to it. Instead of taking a direct ratio, `Math.atan2` takes the two numbers that create the ratio. Its signature is: `Math.atan2(y, x)`. Note that the y value always comes first. Now, because it knows the sign of each value, it can correctly determine what quadrant you are trying to find the angle in, and give you a more useful answer.

Listing 6-3 contains the code that implements this.

    const canvas = document.getElementById("canvas");
    const context = canvas.getContext("2d");
    const width = canvas.width = window.innerWidth;
    const height = canvas.height = window.innerHeight;

    const p0 = new Point(width / 2, height / 2);
    const p1 = new Point(0, 0);

    canvas.addEventListener("click", onClick);
    context.font = "20px Arial";

    function onClick(event) {
      p1.x = event.clientX;
      p1.y = event.clientY;
      render();
    }

    function render() {
      context.clearRect(0, 0, width, height);
      renderPoint(p0);
      renderPoint(p1);

      context.beginPath();
      context.moveTo(p0.x, p0.y);
      context.lineTo(p1.x, p0.y);
      context.lineTo(p1.x, p1.y);
      context.lineTo(p0.x, p0.y);
      context.stroke();

      let x = p1.x - p0.x;
      let y = p1.y - p0.y;
      let rad = Math.atan2(y, x);
      let deg = Math.round(rad * 180 / Math.PI);

      context.fillText(deg + " degrees", p0.x, p0.y - 10);

    }

    function renderPoint(p) {
      context.beginPath();
      context.arc(p.x, p.y, 5, 0, Math.PI * 2);
      context.fill();
    }

*Listing 6-3*

This is only one line different than listing 6-2: 

    let rad = Math.atan2(y, x);

But when you run this again and click in the same area, you now get 150 degrees instead of -30 degrees, as you can see in figure 6-12.

![Using Math.atan2](images/figure_6-12.png)<br/>
*Figure 6-12. Using Math.atan2.*

Note that `Math.atan2` will always return values from -π to +π radians (-180 to +180 degrees). 

Well, let's put this to use. A common use-case for this is to make one object point at another object. For example, we can draw an arrow that always points at the mouse. The code is in listing 6-4.

    const canvas = document.getElementById("canvas");
    const context = canvas.getContext("2d");
    const width = canvas.width = window.innerWidth;
    const height = canvas.height = window.innerHeight;

    canvas.addEventListener("mousemove", onMouseMove);

    function onMouseMove(event) {
      // determine the angle
      const mouseX = event.clientX;
      const mouseY = event.clientY;
      const dx = mouseX - width / 2;
      const dy = mouseY - height / 2;
      const angle = Math.atan2(dy, dx);

      context.clearRect(0, 0, width, height);

      // transform the context by that angle
      context.save();
      context.translate(width / 2, height / 2);
      context.rotate(angle);

      // draw an arrow
      context.beginPath();
      context.moveTo(-30, -10);
      context.lineTo(5, -10);
      context.lineTo(5, -20);
      context.lineTo(30, 0);
      context.lineTo(5, 20);
      context.lineTo(5, 10);
      context.lineTo(-30, 10);
      context.closePath();
      context.fill();

      // restore the context
      context.restore();
    }

*Listing 6-4*

First we set up a listener for the `mousemove` event on the canvas. In that handler, first we get the x, y position of the mouse as `mouseX` and `mouseY`. I'm going to draw my mouse in the center of the canvas, at `width / 2`, `height / 2`, but you could put it at any position you wanted to.

We'll get the difference between where the mouse is and where the arrow will be, and store that in the variables `dx` and `dy`. Then we use `Math.atan2` to get the angle.

I wrote this all out in individual lines to make it clear, but once you know what you are doing, this can easily be compressed into a single line like so:

    const angle = Math.atan2(event.clientY - height / 2, event.clientX - width / 2);

Next we just need to draw that angle in the right position and rotated towards the mouse. First we save the state of the canvas and translate it to the point where we want to draw the arrow. Again, in this case, that's the center point of the canvas.

Then we rotate by the angle we just calculated, draw the arrow, and restore the context so we can do the same thing the next time the mouse moves. Now when you move the mouse around, the arrow should always point to it. See figure 6-13.

![Pointing at the mouse.](images/figure_6-13.png)<br/>
*Figure 6-13. Pointing at the mouse*
