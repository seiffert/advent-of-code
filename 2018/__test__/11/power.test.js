import { Grid } from "../../11/power";

test("Grid 18", () => {
  const g18 = new Grid(18);

  const res3 = g18.largest3x3Square();
  expect(res3.x).toEqual(33);
  expect(res3.y).toEqual(45);

  const res = g18.largestSquare();
  expect(res.x).toEqual(90);
  expect(res.y).toEqual(269);
  expect(res.size).toEqual(16);
});

test("Grid 42", () => {
  const g42 = new Grid(42);

  const res3 = g42.largest3x3Square();
  expect(res3.x).toEqual(21);
  expect(res3.y).toEqual(61);

  const res = g42.largestSquare();
  expect(res.x).toEqual(232);
  expect(res.y).toEqual(251);
  expect(res.size).toEqual(12);
});
