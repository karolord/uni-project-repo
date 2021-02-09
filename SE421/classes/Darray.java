package SE421.classes;

public class Darray {
    private int[] items;
    private int count = 0;

    public Darray(int length) {
        items = new int[length];
    }

    public void insert(int value) {
        items[count] = value;
        count++;
    }

    public void removefirst(int value) {
        for (int i = 0; i < count; i++) {
            if()
        }
    }

    public void printD() {
        for (int i = 0; i < count; i++) {
            System.out.println(items[i]);
        }
    }

}