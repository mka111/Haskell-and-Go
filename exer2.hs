
import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate
hailstone n
  | even n = n `div` 2
  | odd n = 3*n+1


divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]

primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]


join str xs = foldr func "" xs
   where func x xs = if xs == "" then x ++ xs else x ++ str ++ xs


pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(x,y,z) | x <- [1..n], y <- [1..x], z <- [1..n], (x*x)+(y*y) == (z*z)]




merge :: (Ord a) =>[a]->[a]->[a]
merge xs [] = xs
merge [] z =z
merge(x:xs) (z:zs)
  | x<z = x: (merge xs (z:zs))
  | z<=x = z: (merge (x:xs) zs)


hailLen 1 = 0
hailLen n = 1 + hailLen(hailstone n)


hailLen1 :: Int->Int
hailLen1 n = hailTail 0 n
  where
    hailTail x 1 = x
    hailTail x n = hailTail (x+1) (hailstone n)


fact :: Int->Int
fact 0 = 1
fact n = n * fact(n-1)

fact' :: Int -> Int
fact' n = foldl (*) 1 [1..n]



daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
  where jan1 = fromGregorian y 01 01
        dec31 = fromGregorian y 12 31


isFriday :: Day -> Bool
isFriday n
  |snd (mondayStartWeek n) == 5  = True
  |otherwise = False

getDay(y,m,d) = d

ifPrime :: Int -> Bool
ifPrime n 
  |divisors n == [] = True
  |otherwise = False

isPrimeDay :: Day -> Bool
isPrimeDay n = ifPrime(getDay(toGregorian n))


primeFridays :: Integer -> [Day]
primeFridays n = [i | i <- daysInYear n, (isFriday i && isPrimeDay i) ]











