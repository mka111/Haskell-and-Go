import Data.Ratio
import System.IO

rationalSum :: Int -> [Ratio Int]
rationalSum n = [(n-i) % i | i<- [n-1,n-2..1]]


rationalSumLowest :: Int -> [Ratio Int]
rationalSumLowest n = [(n-i) % i | i<- [n-1,n-2..1], gcd (n-i) i == 1]

rationalSumLowestt :: Int -> [Ratio Int]
rationalSumLowestt n = [(n) % (n-i) | i<- [1..n], (n-i) /= 0, gcd (n-i) i == 1]

rationals :: [Ratio Int]
rationals = [(n-i) % i | n<- [1..], i<- [n-1,n-2..1], gcd (n-i) i == 1]


sumFile :: IO ()
sumFile = do
    fh<- readFile "input.txt"
    let intList = convert (lines fh)
    let sum = sumLines(intList)
    print (sum)

convert :: [String] -> [Int]
convert xs = map (read::String->Int) xs

sumLines :: [Int]->Int
sumLines [] = 0
sumLines (x:xs) = x + sumLines xs
