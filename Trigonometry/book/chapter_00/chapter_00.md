# Introduction

Well, here you are, reading a book about trigonometry. Maybe you bought it already, maybe you're just taking a quick look. But here you are. There's a good chance you're grimacing somewhat. Trigonometry can be hard. People struggle with it. You might be thinking that this is a subject you need to get a better handle on, and if this book helps you a bit, it might be worth the pain.

I can't promise you that it won't be hard or it won't be painful. That's somewhat out of my control. But I'm going to do my utmost to make it as easy as possible. If I succeed, it really might not be painful at all. You might even enjoy reading this book.

## Do you need to learn trigonometry for programming?

Not necessarily. It all depends what kind of programming you're doing. If you're doing back end database driven api coding, you might never need a line of trig. If you're doing pure front end application logic and UI code, you might need it now and then, but maybe not.

But if you're doing any kind of graphics programming, simulations, game development, low-level audio programming, or creative coding, you're probably going to use some trigonometry every day of your life.

Trigonometry is the study of the relationship of the angles and side lengths of triangles. So it may seem strange to think that is is used so often, even if you're not drawing triangles. But it turns out that trigonometry's uses far surpass just measuring triangles. Say you want to make something oscillate back and forth or up and down. A sine wave is a great solution. See figure 0-1.

![A sine wave.](images/ch00/figure_0-1.png)
*Figure 0-1. A sine wave.*

Want to move something around in a circle or oval? Or arrange objects around in a circle or oval? Trig is vital there. Want to make an object point at or move towards the mouse, or some other point? Trigonometry again. The distance from one point on the screen to another (as seen in figure 0-2)? Trigonometry is involved.

![Using trig to find distance between two points.](images/ch00/figure_0-2.png)
*Figure 0-2. Using trig to find distance between two points.*

Simulating any kind of physics with acceleration and velocity? You'll need trig. Sound waves are either made directly with sine waves, or can be synthesized by multiple sine waves at different frequencies and strengths.

The list is endless. In this book I'll cover most of the above examples, and a lot more. There's a whole lot more that I won't be able to cover, but once you know the basics, you'll be ready to tackle anything trig-related.

## Why JavaScript?

Please don't be put off by the choice to use JavaScript for this book. What I'm attempting to teach here transcends any one language. But I had do choose one language to write functioning code in. I chose JavaScript for a number of reasons:

#### 1. It's free and open source.

There's nothing to buy. No licenses to agree to. No big corporation singlehandedly deciding its fate.

#### 2. It runs on pretty much any platform in existence today.

And for the purposes of what I'll be covering in this book, there is no difference between the code needed on any particular platform.

#### 3. You don't need any special tools to code, compile and run JavaScript programs.

Just a text editor and any modern browser. Personally, I use Vim and Firefox, but use whatever editor or IDE you want, and whatever browser you prefer, as long as it has been updated in the last year or so.

#### 4. JavaScript is generally pretty easy to learn and understand.

Some may disagree. JS does have its quirks, but most of what I cover in this book is straightforward. There will be no need to worry about pointers or threads, interfaces, inheritance chains, memory management, or any other complex subject that pop up almost immediately in many other languages.

#### 5. HTML Canvas is easy to work with.

Trigonometry is very visual. If you're trying to understand it by looking at a series of numbers your program outputs, you're going to be lost. This book will have tons of visual examples - graphics, animation, interactivity. There are no special graphics libraries to download and install. It's all built into the browser. There's a bit of setup to create the canvas and get access to it, but that's going to be boilerplate that you don't have to worry about after the first couple of times. I'll cover that exact boilerplate later in this introduction.

#### 6. Most importantly, everything you'll learn in this book is easily translated to your language of choice.

I'm going to assume that you have some experience working in at least one, or maybe a few other programming languages. If so, even if you've never touched JavaScript before, the code in this book should not be hard to understand.

Beyond the fact that JavaScript uses common and recognizable syntax for things like conditionals, loops, variable and constant naming, and even classes, it has a decent built-in math library as well. We'll be relying on this library a lot, so let's see how JavaScript's version of the math functions compare to other languages that you might be using.

## Math Library Comparisons

Listing 0-1 shows a few of the more common functions of the JavaScript math library, compared to the math libraries in some other languages you might want to use instead of JavaScript.

*Listing 0-1*

JavaScript         C#                  Processing      Python              Go
-----------------  ------------------  --------------  ------------------  -------------------
Math.sin(angle)    Math.Sin(angle)     sin(angle)      math.sin(angle)     math.Sin(angle)
Math.cos(angle)    Math.Cos(angle)     cos(angle)      math.cos(angle)     math.Cos(angle)
Math.tan(angle)    Math.Tan(angle)     tan(angle)      math.tan(angle)     math.Tan(angle)
Math.atan(ratio)   Math.Atan(ratio)    atan(angle)     math.atan(ratio)    math.Atan(angle)
Math.atan2(y, x)   Math.Atan2(y, x)    atan2(angle)    math.atan2(y, x)    math.Atan2(angle)
Math.sqrt(value)   Math.Sqrt(value)    sqrt(value)     math.sqrt(value)    math.Sqrt(angle)
Math.PI            Math.PI             PI              math.pi             math.Pi

This isn't all of the functions by far, but you can see that for the most part, translating the math functions from JavaScript to some other language is, for the most part, trivial. The functions themselves are mostly standard. To a large degree, it's just a matter of getting the right combination of upper and lower cases. Also, don't worry if you don't know what some (or any) of those things are. They'll all be covered in the book.

## Boilerplate

Let's look at the code that will be used throughout the book to create and run examples.

Each project will have two base files: `index.html` and `main.js`. The HTML file will create the canvas element, apply some minimal styling, and load the `main.js` file, as seen in listing 0-2.

*Listing 0-2*

    <!DOCTYPE html>
    <html>
      <head>
        <style type="text/css">
        html, body {
          margin: 0;
          padding: 0;
        }
        canvas {
          display: block;
        }
        </style>
      </head>
      <body>
        <canvas id="canvas"></canvas>
        <script type="text/javascript" src="main.js"></script>
      </body>
    </html>

To go through it in a bit more detail, the CSS just sets the page to have no margins or padding so that the canvas can be extended to fill the whole browser window. In the body, it creates the canvas and then loads the `main.js` file. The order is important, because the canvas will need to exist before the code tries to access it.

Now the `main.js` file in listing 0-3.

*Listing 0-3*

    const canvas = document.getElementById("canvas");
    const context = canvas.getContext("2d");
    const width = canvas.width = window.innerWidth;
    const height = canvas.height = window.innerHeight;

    // sample drawing code:
    for (let i = 0; i < 100; i++) {
      context.lineTo(Math.random() * width, Math.random() * height);
    }
    context.stroke();

This gets a reference to the canvas element, and gets a reference to the 2d drawing context from the canvas. This context is what you'll use to issue all of your drawing commands.

It sets the canvas size to the size of the browser content area using `window.innerWidth` and `window.innerHeight`, and also saves those values in the `width` and `height` variables for later use.

That itself is all the boilerplate. Below that I included a few lines of sample drawing code. This just executes a `for` loop to draw 100 random lines on the canvas. Figure 0-3 shows the result.

![Random lines.](images/ch00/figure_0-3.png)
*Figure 0-3. Random lines.*

Note that the boilerplate does not include any page loading event listeners or callbacks, classes, or anything else fancy. Also notice that it does not require, include, or import any other external libraries. In more complex example, there will of course be some more complex code, but for the most part I'm going to keep things as bare-bones and simple as humanly possible. This is all in the interest of making the code as understandable and as portable as humanly possible.

- outline of book
\newpage
