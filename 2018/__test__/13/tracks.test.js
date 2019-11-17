import Tracks from "../../13/tracks";

const INPUT = `/->-\\        
|   |  /----\\
| /-+--+-\\  |
| | |  | v  |
\\-+-/  \\-+--/
  \\------/   `;

test("first collision", () => {
  const t = new Tracks(INPUT);
  expect(t.firstCollision()).toBe("7,3");
});

const INPUT2 = `/>-<\\  
|   |  
| /<+-\\
| | | v
\\>+</ |
  |   ^
  \\<->/`;

test("last cart", () => {
  const t = new Tracks(INPUT2);
  expect(t.lastCart()).toBe("6,4");
});
