import { Grid } from "./power";

const g = new Grid(9110);

console.log("Largest power 3x3 square: ", g.largest3x3Square());
console.log("Largest power square: ", g.largestSquare());
