import { processPolymer, improvedProcessPolymer } from "../../5/polymer";

test("polymer", () => {
  expect(processPolymer("dabAcCaCBAcCcaDA").length).toBe(10);
});

test("improved polymer", () => {
  expect(improvedProcessPolymer("dabAcCaCBAcCcaDA").length).toBe(4);
});
