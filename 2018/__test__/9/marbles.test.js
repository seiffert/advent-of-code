import { highscore } from "../../9/marbles";

const cases = [
  ["9 players; last marble is worth 25 points", 32],
  ["10 players; last marble is worth 1618 points", 8317],
  ["13 players; last marble is worth 7999 points", 146373],
  ["17 players; last marble is worth 1104 points", 2764],
  ["21 players; last marble is worth 6111 points", 54718],
  ["30 players; last marble is worth 5807 points", 37305]
];

test.each(cases)("highscore of game '%s' should be %s", (game, hs) => {
  expect(highscore(game)).toBe(hs);
});
