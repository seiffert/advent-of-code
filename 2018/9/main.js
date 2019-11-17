import { highscore } from "./marbles";

console.log(
  "Highscore:",
  highscore("403 players; last marble is worth 71920 points")
);
console.log(
  "After 100x as many rounds:",
  highscore("403 players; last marble is worth 7192000 points")
);
