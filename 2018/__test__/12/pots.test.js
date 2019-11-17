import Pots from "../../12/pots";

const INPUT = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`;

test("pots after 20th generation", () => {
  expect(new Pots(INPUT).proceed(20).sum()).toBe(325);
});
