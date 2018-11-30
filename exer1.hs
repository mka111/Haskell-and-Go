det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a


third_a xs = xs!!2

third_b (_:_:x:xs) = x

hailstone n
  | even n = n `div` 2
  | odd n = 3*n+1





hailLen 1 = 0
hailLen n = 1 + hailLen(hailstone n)

divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]

primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]



join sep xs = foldr (\a b-> a ++ if b=="" then b else sep ++ b) "" xs


join2 sep xs = foldr fix "" xs
   where fix x xs = if xs == "" then x ++ xs else x ++ sep ++ xs


join3:: [Char] -> [[Char]] -> [Char]
join3 str [] = []
join3 str [x] = x
join3 str xs = foldr (\x acc -> x ++ str ++ acc) "" xs


terna :: Int -> [(Int,Int,Int)]
terna x = [(a,b,c)|a<-[1..x], b<-[1..x], c<-[1..x], (a^2)+(b^2) == (c^2)]


pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(x,y,z) | x <- [1..n], y <- [1..x], z <- [1..n], (x*x)+(y*y) == (z*z)]






