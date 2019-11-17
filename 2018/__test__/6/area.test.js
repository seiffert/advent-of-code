import { largestArea, safeRegion } from "../../6/area";

const INPUT = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`;

test("largest area", () => {
  expect(largestArea(INPUT)).toBe(17);
});

test("safe region", () => {
  expect(safeRegion(INPUT, 32)).toBe(16);
});
