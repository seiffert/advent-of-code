import { license, rootValue } from "../../8/license";

const INPUT = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`;

test("license", () => expect(license(INPUT)).toBe(138));
test("value", () => expect(rootValue(INPUT)).toBe(66));
