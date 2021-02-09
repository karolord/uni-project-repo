package SE421.classes;

/**
 * main
 */
public class main {
    public static void main(String[] args) {
        var array = new Darray(5);
        array.insert(50);
        array.insert(10);
        array.insert(20);
        array.insert(30);
        array.printD();

        array.remove(1);
        array.printD();
        System.out.println(array.indexOf(20));
    }
}