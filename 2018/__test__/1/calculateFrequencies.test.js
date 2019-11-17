import { calculateFrequency } from "../../1/frequencies";

const cases = [["+1, +1, +1", 3], ["+1, +1, -2", 0], ["-1, -2, -3", -6]];

test.each(cases)("calculateFrequency", (input, expectedResult) => {
  expect(calculateFrequency(input)).toBe(expectedResult);
});
