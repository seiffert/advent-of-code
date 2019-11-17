import { checksum } from "../../2/checksum";

test("samle input", () => {
  const input = [
    "abcdef",
    "bababc",
    "abbcde",
    "abcccd",
    "aabcdd",
    "abcdee",
    "ababab"
  ];

  expect(checksum(input)).toBe(12);
});
