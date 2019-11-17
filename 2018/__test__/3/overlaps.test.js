import { overlaps, notOverlapping } from "../../3/overlaps";

const INPUT = ["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"];

test("overlaps", () => {
  expect(overlaps(INPUT)).toBe(4);
});

test("not overlapping", () => {
  expect(notOverlapping(INPUT)).toBe(3);
});
