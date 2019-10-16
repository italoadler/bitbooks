# Chapter 6 - Tangent and Arctangent

In the previous two chapters, we've discussed two of the three core trigonometric functions - sine and cosine. We've seen some useful techniques that implement these functions and how closely these two functions relate to each other. But what about the third function, tangent? Is it that different? Does it have no practical uses? Well, yes, it's quite different from both sine and cosine in how it works. But there are some extremely useful techniques that you can employ with tangent. We'll discover some of them in this chapter.

We'll also look into the arctangent function, which is the inverse of tangent. Don't worry about what that means exactly just yet. We'll get there shortly.

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

If you use a web based calculator app, however, you might get some large number like `1.6331239e+16` (that's scientific notation, equal to `1.6331239 * 10^16`. 

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

Let's plot it out and see what kind of curve this creates.
