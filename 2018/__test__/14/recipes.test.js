import Board from "../../14/recipes";

test("ten recipe score", () => {
  const b = new Board("37");

  expect(b.scoreAfter(9)).toBe("5158916779");
  expect(b.scoreAfter(5)).toBe("0124515891");
  expect(b.scoreAfter(18)).toBe("9251071085");
  expect(b.scoreAfter(2018)).toBe("5941429882");
});

test("recipes before", () => {
  const b = new Board("37");

  expect(b.recipesBefore("51589")).toBe(9);
  expect(b.recipesBefore("01245")).toBe(5);
  expect(b.recipesBefore("92510")).toBe(18);
  expect(b.recipesBefore("59414")).toBe(2018);
});
