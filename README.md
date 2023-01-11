# Erect-the-Fence

You are given an array trees where trees[i] = [xi, yi] represents the location of a tree in the garden.

Fence the entire garden using the minimum length of rope, as it is expensive. The garden is well-fenced only if all the trees are enclosed.

Return the coordinates of trees that are exactly located on the fence perimeter. You may return the answer in any order.

 

Example 1:

![erect2-plane](https://user-images.githubusercontent.com/88260025/211899836-fc85e219-f183-4289-83bc-c01394476987.jpg)

Input: trees = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
Output: [[1,1],[2,0],[4,2],[3,3],[2,4]]
Explanation: All the trees will be on the perimeter of the fence except the tree at [2, 2], which will be inside the fence.

Example 2:

![erect1-plane](https://user-images.githubusercontent.com/88260025/211899970-43db9d07-a466-4b45-a136-bb8fcd6474ae.jpg)

Input: trees = [[1,2],[2,2],[4,2]]
Output: [[4,2],[2,2],[1,2]]
Explanation: The fence forms a line that passes through all the trees.
 

Constraints:

1 <= trees.length <= 3000
trees[i].length == 2
0 <= xi, yi <= 100
All the given positions are unique.
