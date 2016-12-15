module Hailstone

where
	import Data.List
	import System

-- gebruik where functie?
	--hailstone :: Int -> [Int]
	hailstone 1 = [1]
	hailstone n 
		| even n = n : hailstone (n `div` 2) -- n mod 2 == 0
		| otherwise = n : hailstone (n * 3 + 1)

	main = do
	    n <- getArgs
	    print . hailstones $ (read  (n !! 0) :: Integer)
