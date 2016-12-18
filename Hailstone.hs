-- Caitlin Schäffers
-- Programmeertalen
-- University of Amsterdam

module Hailstone

where
    import Data.List

    hailstone :: Int -> [Int]
    hailstone 1 = [1]
    hailstone n 
        | even n = n : hailstone (n `div` 2)
        | otherwise = n : hailstone (n * 3 + 1)

    hailstone' :: Int -> [Int]
    hailstone' = (!!) hailstoneMemo

    hailstoneMemo :: [[Int]]
    hailstoneMemo = (map all_hail [0 ..])
        where all_hail 1 = [1]
              all_hail n
                | even n = n : hailstone' (n `div` 2)
                | otherwise = n : hailstone' (n * 3 + 1)
