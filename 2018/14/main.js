import Board from "./recipes";

console.log("After 6: ", new Board("37").scoreAfter(6));
console.log("After 74501: ", new Board("37").scoreAfter(74501));

console.log("Before '074501': ", new Board("37").recipesBefore("074501"));
