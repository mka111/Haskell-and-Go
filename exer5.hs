
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

