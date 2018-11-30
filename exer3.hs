import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

hailstone n
  | even n = n `div` 2
  | odd n = 3*n+1


divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]

merge :: (Ord a) =>[a]->[a]->[a]
merge xs [] = xs
merge [] z =z
merge(x:xs) (z:zs)
  | x<z = x: (merge xs (z:zs))
  | z<=x = z: (merge (x:xs) zs)



hailLen :: Int->Int
hailLen n = hailTail 0 n
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



hailSeq :: Int -> [Int]
hailSeq 1 = [1]
hailSeq n = n: hailSeq(hailstone n)

hailSeq' :: Int -> [Int]
hailSeq' n = take ((hailLen n) +1) (iterate (hailstone) n)

join str xs = foldr func "" xs
   where func x xs = if xs == "" then x ++ xs else x ++ str ++ xs

halfLen :: (Ord a) =>[a] -> Int
halfLen xs = length xs `div`  2

firstList :: (Ord a) =>[a] -> [a]
firstList xs = take (halfLen xs) xs

secondList :: (Ord a) =>[a] -> [a]
secondList xs = drop (halfLen xs) xs

mergeSort :: (Ord a) =>[a]->[a]
mergeSort [] = []
mergeSort [a] = [a]
mergeSort xs = merge (mergeSort (firstList xs)) (mergeSort (secondList xs))

findElt :: (Ord a) => a ->[a] -> Maybe Int
findElt x xs = findElem 0 xs
  where 
    findElem j [] = Nothing
    findElem j (a:xs)
      |a==x  = Just j
      |otherwise = findElem(j+1) xs


myIterate :: (a -> a) ->a -> [a]
myIterate func x = x : myIterate func (func x)


myTakeWhile :: (a->Bool) -> [a] -> [a]
myTakeWhile p xs = function p xs
 where
  function p [] = []
  function p (x:xs)
     |p x = x : (myTakeWhile p xs)
     |otherwise = []

pascal :: Int -> [Int]
pascal 0 = [1]
pascal 1 = [1,1]
pascal n = [1] ++ (zipWith (+) (pascal(n-1)) (tail(pascal(n-1)))) ++ [1]




addPair :: (Int, Int) -> Int
addPair = uncurry (+)

withoutZeros :: (Eq a, Num a) => [a] -> [a]
withoutZeros = filter(/=0)


fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib 2 = 1
fib n = fib(n-1) + fib(n-2)

fibs = map fib [0..]

things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)








