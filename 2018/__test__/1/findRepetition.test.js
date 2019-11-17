import { findRepetition } from "../../1/frequencies";

const cases = [
  ["+1, -1", 0],
  ["+3, +3, +4, -2, -4", 10],
  ["-6, +3, +8, +5, -6", 5],
  ["+7, +7, -2, -7, -4", 14]
];

test.each(cases)("findRepetition", (input, expectedResult) => {
  expect(findRepetition(input)).toBe(expectedResult);
});
