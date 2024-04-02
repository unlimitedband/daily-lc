/**
 * @param {number[]} gifts
 * @param {number} k
 * @return {number}
 */
var pickGifts = function(gifts, k) {
    const q = new MaxPriorityQueue();
    for (const i of gifts) {
        q.enqueue(i);
    }
    
    while (k--) {
        const val = q.dequeue().element;
        const fsqrt = Math.floor(Math.sqrt(val));
        q.enqueue(fsqrt);
    }
    
    let sum = 0;
   while (!q.isEmpty()) {
       sum += q.dequeue().element;
   }
    
    return sum;
};
