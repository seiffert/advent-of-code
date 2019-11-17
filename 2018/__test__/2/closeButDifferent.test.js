import { closeButDifferent } from "../../2/checksum";

test("sample input", () => {
  const input = ["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"];

  expect(closeButDifferent(input)).toBe("fgij");
});
