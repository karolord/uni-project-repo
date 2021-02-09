package SE421.classes;

public class Darray {
    private int[] items;
    private int count = 0;

    public Darray(int length) {
        items = new int[length];
    }

    public void insert(int value) {
        if (items.length == count) {
            int[] big = new int[count * 2];
            for (int i = 0; i < count; i++) {
                big[i] = items[i];
            }
            items = big;
        }
        items[count++] = value;
    }

    public void remove(int index) {
        if (index < 0 || index >= count)
            throw new IllegalArgumentException();
        for (int i = index; i < count; i++)
            items[i] = items[i + 1];

        count--;
    }

    public void printD() {
        for (int i = 0; i < count; i++) {
            System.out.println(items[i]);
        }
    }

    public int indexOf(int value) {
        for (int i = 0; i < count; i++) {
            if (items[i] == value)
                return i;
        }
        return -1;
    }
}