module AdventOfCode2021.Day01

// Obtains the number of times a depth measurement increases
let solvePartOne =
    List.pairwise
    >> List.map (fun (x, y) -> if y > x then "I" else "D")
    >> List.fold (fun acc id -> if id = "I" then acc + 1 else acc) 0

// Obtains the number of times a sliding window of 3 increases
let solvePartTwo =
    List.windowed 3
    >> List.map (List.sum)
    >> solvePartOne

let test fn want =
    [199;200;208;210;200;207;240;269;260;263]
    |> fn
    |> fun got ->
        if got <> want
        then sprintf "want: %i, got: %i" want got
        else sprintf "OK"

test solvePartOne 7 |> printfn "Testing part one solution: %s"
test solvePartTwo 5 |> printfn "Testing part two solution: %s"

let input = [] // Puzzle input
solvePartOne input |> printfn "Part one result: %i"
solvePartTwo input |> printfn "Part two result: %i"