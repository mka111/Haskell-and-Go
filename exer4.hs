
hailstone n
  | even n = n `div` 2
  | odd n = 3*n+1

hailLen :: Int->Int
hailLen n = hailTail 0 n
  where
    hailTail x 1 = x
    hailTail x n = hailTail (x+1) (hailstone n)

merge :: (Ord a) =>[a]->[a]->[a]
merge xs [] = xs
merge [] z =z
merge(x:xs) (z:zs)
  | x<z = x: (merge xs (z:zs))
  | z<=x = z: (merge (x:xs) zs)



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


