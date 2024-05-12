import java.util.HashMap;
class Cashier {
    int n;
    int discount;
    int curCustomer;
    HashMap<Integer, Integer> pricesMap;
    public Cashier(int n, int discount, int[] products, int[] prices) {
        this.n = n;
        this.curCustomer = 0;
        this.discount = discount;
        this.pricesMap = new HashMap<>();
        for (int i = 0; i < products.length; i++) this.pricesMap.put(products[i], prices[i]);
    }
    
    public double getBill(int[] product, int[] amount) {
        double value = 0;
        for (int i = 0; i < product.length; i++) value += amount[i] * pricesMap.get(product[i]);
        curCustomer = (curCustomer + 1) % n;
        return (curCustomer == 0) ? (value * (1 - (double)discount/100)) : value;
    }
}

/**
 * Your Cashier object will be instantiated and called as such:
 * Cashier obj = new Cashier(n, discount, products, prices);
 * double param_1 = obj.getBill(product,amount);
 */
